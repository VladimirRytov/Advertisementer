package exporthandler

import (
	"context"
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

type ExportHandler struct {
	fileStorage FileStorage
	reqGate     DataBase
	app         Responcer
}

func NewExportHandler(app Responcer, reqGate DataBase, files FileStorage) *ExportHandler {
	return &ExportHandler{app: app, reqGate: reqGate, fileStorage: files}
}

func (ex *ExportHandler) ExportJsonToFile(ctx context.Context, path string) error {
	ex.app.SendMessage("Экспорт клиентов...")
	clients, err := ex.reqGate.AllClients(ctx)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return errors.New("операция отменена")
		}
		ex.app.SendError(err)
		return err
	}
	ex.app.SendMessage("Экспорт заказов...")

	for i := range clients {
		ex.app.SendMessage("Экспорт заказов клиента " + clients[i].Name + "...")
		order, err := ex.reqGate.OrdersByClientName(ctx, clients[i].Name)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				ex.app.SendError(errors.New("операция отменена"))
				return errors.New("операция отменена")
			}
			ex.app.SendError(err)
			return err
		}

		clients[i].Orders = append(clients[i].Orders, order...)
		for j := range clients[i].Orders {
			blocks, err := ex.reqGate.BlockAdvertisementsByOrderID(ctx, clients[i].Orders[j].ID)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					ex.app.SendError(errors.New("операция отменена"))
					return err
				}
			}

			clients[i].Orders[j].BlockAdvertisements = append(clients[i].Orders[j].BlockAdvertisements, blocks...)

			lines, err := ex.reqGate.LineAdvertisementsByOrderID(ctx, clients[i].Orders[j].ID)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					ex.app.SendError(errors.New("операция отменена"))
					return err
				}
			}
			clients[i].Orders[j].LineAdvertisements = append(clients[i].Orders[j].LineAdvertisements, lines...)
		}
	}

	ex.app.SendMessage("Экспорт тэгов...")
	tags, err := ex.reqGate.AllTags(ctx)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			ex.app.SendError(errors.New("операция отменена"))
			return err
		}
		ex.app.SendError(err)
		return err
	}

	ex.app.SendMessage("Экспорт наценок...")
	extraCharges, err := ex.reqGate.AllExtraCharges(ctx)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			ex.app.SendError(errors.New("операция отменена"))
			return err
		}
	}

	costRates, err := ex.reqGate.AllCostRates(ctx)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			ex.app.SendError(errors.New("операция отменена"))
			return err
		}
	}
	select {
	case <-ctx.Done():
		logging.Logger.Info("exportHandler.ExportJsonToFile: canceled")
		return errors.New("exportHandler.ExportJsonToFile: canceled")
	default:
		ex.app.SendMessage("Запись в файл...")
		js := &datatransferobjects.JsonStr{
			Clients:      clients,
			Tags:         tags,
			ExtraCharges: extraCharges,
			CostRates:    costRates,
		}

		err := ex.toJSON(js, path)
		if err != nil {
			ex.app.SendError(err)
			return err
		}
	}
	ex.app.ProgressComplete()
	return nil
}

func (ex *ExportHandler) toJSON(packedEntities *datatransferobjects.JsonStr, path string) error {
	ex.app.SendMessage("Запись в файл...")

	f, err := ex.fileStorage.OpenForWrite(path)
	if err != nil {
		ex.app.SendError(err)
		return err
	}
	defer f.Close()
	err = encodedecoder.ToJSON(f, packedEntities, false)
	if err != nil {
		ex.app.SendError(err)
		return err
	}

	ex.app.ProgressComplete()
	return nil
}
