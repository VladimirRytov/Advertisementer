package importhandler

import (
	"context"
	"errors"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

type Importer struct {
	ctx         context.Context
	fileManager FileStorage
	resp        Responcer
	closest     DateChecker
	reqGate     DataBase
	params      datatransferobjects.ImportParams
	js          datatransferobjects.JsonStr
}

func CreateImporter(app Responcer, checker DateChecker, fileManager FileStorage, reqGate DataBase) *Importer {
	return &Importer{resp: app, fileManager: fileManager, closest: checker, reqGate: reqGate}
}

func (im *Importer) ImportJson(ctx context.Context, importPath string, params datatransferobjects.ImportParams) error {
	im.params = params
	im.ctx = ctx
	f, err := im.fileManager.OpenForRead(importPath)
	if err != nil {
		im.resp.CancelProgressWithError(err)
		return err
	}
	err = encodedecoder.FromJSON(&im.js, f)
	if err != nil {
		logging.Logger.Error("importer.ImportJson: an error occured while parging json", "error", err)

		im.resp.CancelProgressWithError(errors.New("неверный тип файла"))
		return err
	}
	im.resp.SendMessage("Выполняется импорт данных...")

	if im.params.AllTags {
		im.resp.SendMessage("Выполняется импорт тэгов...")
		err = im.createTags()
		if err != nil {
			im.resp.CancelProgressWithError(err)
			return err
		}
	}

	if im.params.AllExtraCharges {
		im.resp.SendMessage("Выполняется импорт наценок...")
		err = im.createExtraCharges()
		if err != nil {
			im.resp.CancelProgressWithError(err)
			return err
		}
	}

	im.resp.SendMessage("Выполняется импорт данных контрагентов...")
	err = im.createClients()
	if err != nil {
		im.resp.CancelProgressWithError(err)
		return err
	}

	err = im.updateTags()
	if err != nil {
		im.resp.CancelProgressWithError(err)
		return err
	}
	err = im.updateExtraCharges()
	if err != nil {
		im.resp.CancelProgressWithError(err)
		return err
	}
	im.resp.SendMessage("Выполняется импорт тарифных планов...")

	err = im.createCostRates()
	if err != nil {
		im.resp.CancelProgressWithError(err)
		return err
	}
	im.resp.ProgressComplete()
	return nil
}

func (im *Importer) createClients() error {
	logging.Logger.Debug("Importer.createClients: creating client")

	for i := range im.js.Clients {
		select {
		case <-im.ctx.Done():
			return im.ctx.Err()
		default:
			if im.params.ActualClients {
				if !im.actualClients(&im.js.Clients[i]) {
					continue
				}
			}
			if im.params.ThickMode {
				orders := im.js.Clients[i].Orders
				client, err := mapper.DtoToClient(&im.js.Clients[i])
				if err != nil {
					logging.Logger.Error("Importer.createClients: cant convert client to model", "error", err)
					im.resp.SendMessage(err.Error())
					if !im.params.IgnoreErrors {
						return err
					}
				}
				im.js.Clients[i] = mapper.ClientToDTO(&client)
				im.js.Clients[i].Orders = orders
			}
			_, err := im.reqGate.NewClient(im.ctx, &im.js.Clients[i])
			if err != nil {
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
			err = im.createOrders(im.js.Clients[i].Orders)
			if err != nil {
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
		}
	}
	return nil
}

func (im *Importer) actualClients(client *datatransferobjects.ClientDTO) bool {
	for i := range client.Orders {
		if im.actualOrders(&client.Orders[i]) {
			return true
		}
	}
	return false
}

func (im *Importer) actualOrders(order *datatransferobjects.OrderDTO) bool {
	order.BlockAdvertisements = im.filterBlockAdvertisement(order.BlockAdvertisements)
	order.LineAdvertisements = im.filterLineAdvertisements(order.LineAdvertisements)
	return len(order.BlockAdvertisements) != 0 || len(order.LineAdvertisements) != 0
}

func (im *Importer) createOrders(orders []datatransferobjects.OrderDTO) error {
	for i := range orders {
		select {
		case <-im.ctx.Done():
			return im.ctx.Err()

		default:
			orders[i].ID = 0
			if !im.actualOrders(&orders[i]) {
				continue
			}

			if im.params.ThickMode {
				order, err := mapper.DtoToOrder(&orders[i])
				if err != nil {
					im.resp.SendMessage(err.Error())
					if !im.params.IgnoreErrors {
						return err
					}
				}
				block, err := im.convertBlockAdvertisements(orders[i].BlockAdvertisements)
				if err != nil {
					im.resp.SendMessage(err.Error())
					if !im.params.IgnoreErrors {
						return err
					}
				}
				line, err := im.convertLineAdvertisements(orders[i].LineAdvertisements)
				if err != nil {
					im.resp.SendMessage(err.Error())
					if !im.params.IgnoreErrors {
						return err
					}
				}
				orders[i] = mapper.OrderToDTO(&order)
				orders[i].BlockAdvertisements = block
				orders[i].LineAdvertisements = line
			}

			_, err := im.reqGate.NewAdvertisementsOrder(im.ctx, &orders[i])
			if err != nil {
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
		}
	}
	return nil
}

func (im *Importer) convertBlockAdvertisements(blocks []datatransferobjects.BlockAdvertisementDTO) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	blocksDto := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blocks))
	for i := range blocks {
		blockAdv, err := mapper.DtoToAdvertisementBlock(&blocks[i])
		if err != nil {
			im.resp.SendMessage(err.Error())
			if !im.params.IgnoreErrors {
				return nil, err
			}
		}
		blocksDto = append(blocksDto, mapper.BlockAdvertisementToDTO(&blockAdv))
	}
	return blocksDto, nil
}

func (im *Importer) filterBlockAdvertisement(blockAdv []datatransferobjects.BlockAdvertisementDTO) []datatransferobjects.BlockAdvertisementDTO {
	blocks := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdv))
	for i := range blockAdv {
		select {
		case <-im.ctx.Done():
			return nil
		default:
			blockAdv[i].ID = 0
			blockAdv[i].OrderID = 0
			if im.params.AllBlocks {
				blocks = append(blocks, blockAdv[i])
				continue
			}
			_, err := im.closest.ClosestRelease(blockAdv[i].ReleaseDates, time.Now())
			if err == nil {
				blocks = append(blocks, blockAdv[i])
			}
		}
	}
	return blocks
}

func (im *Importer) convertLineAdvertisements(lines []datatransferobjects.LineAdvertisementDTO) ([]datatransferobjects.LineAdvertisementDTO, error) {
	lineDto := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lines))
	for i := range lines {
		lineAdv, err := mapper.DtoToAdvertisementLine(&lines[i])
		if err != nil {
			im.resp.SendMessage(err.Error())
			if !im.params.IgnoreErrors {
				return nil, err
			}
		}
		lineDto = append(lineDto, mapper.LineAdvertisementToDTO(&lineAdv))
	}
	return lineDto, nil
}

func (im *Importer) filterLineAdvertisements(lineAdv []datatransferobjects.LineAdvertisementDTO) []datatransferobjects.LineAdvertisementDTO {
	actualLine := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdv))
	for i := range lineAdv {
		select {
		case <-im.ctx.Done():
			return nil
		default:
			lineAdv[i].ID = 0
			lineAdv[i].OrderID = 0
			if im.params.AlllLines {
				actualLine = append(actualLine, lineAdv[i])
				continue
			}
			_, err := im.closest.ClosestRelease(lineAdv[i].ReleaseDates, time.Now())
			if err == nil {
				actualLine = append(actualLine, lineAdv[i])
			}
		}
	}
	return actualLine
}

func (im *Importer) createTags() error {
	for i := range im.js.Tags {
		select {
		case <-im.ctx.Done():
			return im.ctx.Err()
		default:
			tag, err := mapper.DtoToTag(&im.js.Tags[i])
			if err != nil {
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
			tagDto := mapper.TagToDTO(&tag)
			_, err = im.reqGate.NewTag(im.ctx, &tagDto)
			if err != nil {
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
		}
	}
	return nil
}

func (im *Importer) updateTags() error {
	for i := range im.js.Tags {
		select {
		case <-im.ctx.Done():
			return im.ctx.Err()
		default:
			tag, err := mapper.DtoToTag(&im.js.Tags[i])
			if err != nil {
				logging.Logger.Error("Importer.updateTags: an error occured while converting tag", "error", err)
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
			tagDto := mapper.TagToDTO(&tag)
			err = im.reqGate.UpdateTag(im.ctx, &tagDto)
			if err != nil {
				logging.Logger.Error("Importer.updateTags: an error occured while update tag", "error", err)
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
		}
	}
	return nil
}

func (im *Importer) createExtraCharges() error {
	for i := range im.js.ExtraCharges {
		select {
		case <-im.ctx.Done():
			return im.ctx.Err()
		default:
			chargeAdv, err := mapper.DtoToExtraCharge(&im.js.ExtraCharges[i])
			if err != nil {
				logging.Logger.Error("Importer.sendExtraCharges: an error occured while converting extraCharge", "error", err)
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
			chargeDto := mapper.ExtraChargeToDTO(&chargeAdv)
			_, err = im.reqGate.NewExtraCharge(im.ctx, &chargeDto)
			if err != nil {
				logging.Logger.Error("Importer.sendExtraCharges: an error occured while sending extraCharge to database", "error", err)
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
		}
	}
	return nil
}

func (im *Importer) updateExtraCharges() error {
	for i := range im.js.ExtraCharges {
		select {
		case <-im.ctx.Done():
			return im.ctx.Err()
		default:
			chargeAdv, err := mapper.DtoToExtraCharge(&im.js.ExtraCharges[i])
			if err != nil {
				logging.Logger.Error("Importer.sendExtraCharges: an error occured while converting extraCharge", "error", err)
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
			chargeDto := mapper.ExtraChargeToDTO(&chargeAdv)
			err = im.reqGate.UpdateExtraCharge(im.ctx, &chargeDto)
			if err != nil {
				logging.Logger.Error("Importer.sendExtraCharges: an error occured while sending extraCharge to database", "error", err)
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
		}
	}
	return nil
}

func (im *Importer) createCostRates() error {
	for i := range im.js.CostRates {
		select {
		case <-im.ctx.Done():
			return im.ctx.Err()
		default:
			costRateAdv, err := mapper.DtoToCostRate(&im.js.CostRates[i])
			if err != nil {
				logging.Logger.Error("Importer.createCostRates: an error occured while converting costRate", "error", err)
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
			costRateDto := mapper.CostRateToDTO(&costRateAdv)
			_, err = im.reqGate.NewCostRate(im.ctx, &costRateDto)
			if err != nil {
				logging.Logger.Error("Importer.createCostRates: an error occured while sending costRate to database", "error", err)
				im.resp.SendMessage(err.Error())
				if !im.params.IgnoreErrors {
					return err
				}
			}
		}
	}
	return nil
}
