package exelexporter

import (
	"context"
	"io/fs"
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/advertisements"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/filestorage"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelDebug}, false, os.Stderr)
}

func TestNewReport(t *testing.T) {
	CreateLogger()
	blocks, lines := fillTestData()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	repGen := NewExelReporter(&filestorage.Storage{}, &filestorage.Storage{})
	par := datatransferobjects.ReportParams{
		ReportType:       "Exel",
		FromDate:         time.Now(),
		ToDate:           time.Now(),
		BlocksFolderPath: ".",
		DeployPath:       ".",
	}
	err := repGen.NewReport(ctx, blocks, lines, par)
	if err != nil {
		t.Fatal(err)
	}
	got, err := fs.Glob(os.DirFS("."), "Сводка*")
	if err != nil {
		t.Fatal(err)
	}
	for i := range got {
		os.RemoveAll(got[i])
	}
}

type Database interface {
	OrderByID(context.Context, int) (datatransferobjects.OrderDTO, error)
	NewClient(context.Context, *datatransferobjects.ClientDTO) (string, error)
	NewAdvertisementsOrder(context.Context, *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error)
	NewLineAdvertisement(context.Context, *datatransferobjects.LineAdvertisementDTO) (int, error)
	NewBlockAdvertisement(context.Context, *datatransferobjects.BlockAdvertisementDTO) (int, error)
	BlockAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.LineAdvertisementDTO, error)
}

func fillTestData() ([]datatransferobjects.BlockAdvertisementReport, []datatransferobjects.LineAdvertisementReport) {
	var (
		linesReport  []datatransferobjects.LineAdvertisementReport  = make([]datatransferobjects.LineAdvertisementReport, 0)
		blocksReport []datatransferobjects.BlockAdvertisementReport = make([]datatransferobjects.BlockAdvertisementReport, 0)
	)

	for i := 1; i < 11; i++ {
		line := advertisements.NewAdvertisementLine()
		line.SetOrderId(i)
		line.SetReleaseDates([]time.Time{time.Now().Add(24 * time.Hour)})
		linesReport = append(linesReport, datatransferobjects.LineAdvertisementReport{
			Line: mapper.LineAdvertisementToDTO(&line),
			Paid: i%2 == 0,
		})

		block := advertisements.NewAdvertisementBlock()
		block.SetOrderId(i)
		block.SetSize(10)
		block.SetFileName(strconv.Itoa(i))

		block.SetReleaseDates([]time.Time{time.Now().Add(24 * time.Hour)})
		blocksReport = append(blocksReport, datatransferobjects.BlockAdvertisementReport{
			Block: mapper.BlockAdvertisementToDTO(&block),
			Paid:  i%2 == 0,
		})
	}
	return blocksReport, linesReport
}
