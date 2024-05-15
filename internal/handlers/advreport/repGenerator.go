package advreport

import (
	"context"
	"errors"
	"os"
	"slices"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/handlers"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

const (
	CancelStr  = "операция была отменена пользователем"
	TimeoutStr = "превышено время ожидания ответа базы данных. Операция отменена"
)

type ReportGenerators map[string]handlers.ReportGenerator

type Rep struct {
	reqGate          handlers.ReportRequestGate
	progresHandle    handlers.Progresser
	reportGenerators ReportGenerators
	params           datatransferobjects.ReportParams
	orderPayStatus   map[int]bool
}

func NewReportGenerator() *Rep {
	return &Rep{}
}

func (r *Rep) InitReportGenerator(progr handlers.Progresser, reqGate handlers.ReportRequestGate) {
	r.progresHandle = progr
	r.reqGate = reqGate
	r.reportGenerators = make(ReportGenerators)
	r.orderPayStatus = make(map[int]bool)
}

func (r *Rep) RegisterGenerator(name string, gen handlers.ReportGenerator) {
	r.reportGenerators[name] = gen
}

func (r *Rep) GenerateReport(ctx context.Context, params datatransferobjects.ReportParams) {
	r.params = params

	r.progresHandle.SendMessage("обработка строчных объявлений...")
	lines, err := r.collectLines(ctx)
	if err != nil {
		r.handleError(err)
		return
	}

	r.progresHandle.SendMessage("обработка блочных объявлений...")
	blocks, err := r.collectBlocks(ctx)
	if err != nil {
		r.handleError(err)
		return
	}

	r.progresHandle.SendMessage("генерация сводки...")
	reportGenerator, ok := r.reportGenerators[params.ReportType]
	if !ok {
		r.progresHandle.CancelProgressWithError(errors.New("генератор сводки не найден"))
		return
	}
	err = reportGenerator.NewReport(ctx, blocks, lines, r.params)
	if err != nil {
		r.handleError(err)
		return
	}
	for k := range r.orderPayStatus {
		delete(r.orderPayStatus, k)
	}
	r.progresHandle.ProgressComplete()
}

func (r *Rep) collectBlocks(ctx context.Context) ([]datatransferobjects.BlockAdvertisementReport, error) {
	logging.Logger.Debug("collectBlocks: start collecting block advertisements")
	var (
		blocks []datatransferobjects.BlockAdvertisementDTO
		err    error
	)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		timectx, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()
		if r.params.FromDate.Equal(r.params.ToDate) {
			blocks, err = r.reqGate.BlockAdvertisementBetweenReleaseDates(timectx, r.params.FromDate, r.params.ToDate.Add(24*time.Hour))
		} else {
			blocks, err = r.reqGate.BlockAdvertisementBetweenReleaseDates(timectx, r.params.FromDate, r.params.ToDate)
		}
		logging.Logger.Debug("Rep.collectBlocks: search request", "got blockAdvertisements", len(blocks))

		if err != nil {
			return nil, err
		}
		blockAdvs, err := r.convertAndSortBlocks(ctx, blocks)
		if err != nil {
			return nil, err
		}
		return blockAdvs, nil
	}
}

func (r *Rep) convertAndSortBlocks(ctx context.Context, blocks []datatransferobjects.BlockAdvertisementDTO) ([]datatransferobjects.BlockAdvertisementReport, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		blockAdvs := make([]datatransferobjects.BlockAdvertisementReport, 0, len(blocks))
		for i := range blocks {
			block, err := mapper.DtoToAdvertisementBlock(&blocks[i])
			if err != nil {
				logging.Logger.Error("collectLines: an error occured while converting line to entiry", "error", err)
				r.progresHandle.SendMessage(err.Error())
				continue
			}

			paymentStatus, ok := r.orderPayStatus[block.OrderId()]
			if !ok {
				paymentStatus, err = r.checkOrderPaymentStatus(ctx, block.OrderId())
				if err != nil {
					logging.Logger.Error("collectLines: an error occured while converting line to entiry", "error", err)
					r.progresHandle.SendMessage(err.Error())
					continue
				}
			}
			r.orderPayStatus[block.OrderId()] = paymentStatus
			blockAdvs = append(blockAdvs, datatransferobjects.BlockAdvertisementReport{
				Block: mapper.BlockAdvertisementToDTO(&block),
				Paid:  paymentStatus,
			})
		}
		slices.SortFunc(blockAdvs, func(a, b datatransferobjects.BlockAdvertisementReport) int { return a.Block.OrderID - b.Block.OrderID })
		return blockAdvs, nil
	}
}

func (r *Rep) checkOrderPaymentStatus(ctx context.Context, orderID int) (bool, error) {
	order, err := r.reqGate.OrderByID(ctx, orderID)
	if err != nil {
		return false, err
	}
	orderAdv, err := mapper.DtoToOrder(&order)
	if err != nil {
		return false, err
	}
	return orderAdv.PaymentStatus(), nil
}

func (r *Rep) collectLines(ctx context.Context) ([]datatransferobjects.LineAdvertisementReport, error) {
	logging.Logger.Debug("collectBlocks: start collecting line advertisements")
	var (
		lines []datatransferobjects.LineAdvertisementDTO
		err   error
	)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		timectx, cancel := context.WithTimeout(ctx, 1*time.Minute)
		defer cancel()
		if r.params.FromDate.Equal(r.params.ToDate) {
			lines, err = r.reqGate.LineAdvertisementBetweenReleaseDates(timectx, r.params.FromDate, r.params.ToDate.Add(24*time.Hour))
		} else {
			lines, err = r.reqGate.LineAdvertisementBetweenReleaseDates(timectx, r.params.FromDate, r.params.ToDate)
		}
		if err != nil {
			return nil, err
		}

		lineAdvs, err := r.convertAndSortLines(ctx, lines)
		if err != nil {
			return nil, err
		}
		return lineAdvs, nil
	}
}

func (r *Rep) convertAndSortLines(ctx context.Context, lines []datatransferobjects.LineAdvertisementDTO) ([]datatransferobjects.LineAdvertisementReport, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		lineAdvs := make([]datatransferobjects.LineAdvertisementReport, 0, len(lines))
		for i := range lines {
			line, err := mapper.DtoToAdvertisementLine(&lines[i])
			if err != nil {
				logging.Logger.Error("collectLines: an error occured while converting line to entiry", "error", err)
				r.progresHandle.SendMessage(err.Error())
				continue
			}
			paymentStatus, ok := r.orderPayStatus[line.OrderId()]
			if !ok {
				paymentStatus, err = r.checkOrderPaymentStatus(ctx, line.OrderId())
				if err != nil {
					logging.Logger.Error("collectLines: an error occured while converting line to entiry", "error", err)
					r.progresHandle.SendMessage(err.Error())
					continue
				}
			}

			lineAdvs = append(lineAdvs, datatransferobjects.LineAdvertisementReport{
				Line: mapper.LineAdvertisementToDTO(&line),
				Paid: paymentStatus,
			})
		}
		slices.SortFunc(lineAdvs, func(a, b datatransferobjects.LineAdvertisementReport) int { return a.Line.OrderID - b.Line.OrderID })
		return lineAdvs, nil
	}
}

func (r *Rep) handleError(err error) {
	logging.Logger.Error("generateReport: got error ", "error", err)
	switch {
	case errors.Is(err, context.Canceled):
		r.progresHandle.CancelProgressWithError(errors.New(CancelStr))
	case errors.Is(err, context.DeadlineExceeded):
		r.progresHandle.CancelProgressWithError(errors.New(TimeoutStr))
	case errors.Is(err, os.ErrPermission):
		r.progresHandle.CancelProgressWithError(errors.New("недостаточно прав для создания папки"))
	case errors.Is(err, os.ErrExist):
		r.progresHandle.CancelProgressWithError(errors.New("каталог уже существует"))
	default:
		r.progresHandle.CancelProgressWithError(err)
	}
}
