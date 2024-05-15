package importhandler

import (
	"context"
	"io"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type FileStorage interface {
	OpenForWrite(string) (io.WriteCloser, error)
	OpenForRead(string) (io.ReadSeekCloser, error)
}

type Responcer interface {
	CancelProgressWithError(error)
	ProgressComplete()
	SendMessage(string)
}

type DateChecker interface {
	ClosestRelease([]time.Time, time.Time) (time.Time, error)
}

type DataBase interface {
	Creator
	Getter
	Updater
}

type Creator interface {
	NewClient(context.Context, *datatransferobjects.ClientDTO) (string, error)
	NewAdvertisementsOrder(context.Context, *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error)
	NewLineAdvertisement(context.Context, *datatransferobjects.LineAdvertisementDTO) (int, error)
	NewBlockAdvertisement(context.Context, *datatransferobjects.BlockAdvertisementDTO) (int, error)
	NewExtraCharge(context.Context, *datatransferobjects.ExtraChargeDTO) (string, error)
	NewTag(context.Context, *datatransferobjects.TagDTO) (string, error)
	NewCostRate(context.Context, *datatransferobjects.CostRateDTO) (string, error)
}

type Updater interface {
	UpdateExtraCharge(context.Context, *datatransferobjects.ExtraChargeDTO) error
	UpdateTag(context.Context, *datatransferobjects.TagDTO) error
	UpdateCostRate(context.Context, *datatransferobjects.CostRateDTO) error
}

type Getter interface {
	ClientByName(context.Context, string) (datatransferobjects.ClientDTO, error)
	OrderByID(context.Context, int) (datatransferobjects.OrderDTO, error)
	LineAdvertisementByID(context.Context, int) (datatransferobjects.LineAdvertisementDTO, error)
	BlockAdvertisementByID(context.Context, int) (datatransferobjects.BlockAdvertisementDTO, error)
	TagByName(context.Context, string) (datatransferobjects.TagDTO, error)
	ExtraChargeByName(context.Context, string) (datatransferobjects.ExtraChargeDTO, error)
	CostRateByName(context.Context, string) (datatransferobjects.CostRateDTO, error)
}
