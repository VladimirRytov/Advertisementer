package thickclienthandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

func (a *AdvertisementController) AllClients() {
	logging.Logger.Debug("advertisementhandler: search request. Search All Clients")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	recievedClientsDto, err := a.req.AllClients(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllClients: cannot get clients", "error", err)
		a.handleError(err)
		return
	}

	for i := range recievedClientsDto {
		clientModel, err := mapper.DtoToClient(&recievedClientsDto[i])
		if err != nil {
			logging.Logger.Error("client dont pass the filter", "clientName", recievedClientsDto[i].Name, "error", err)
			continue
		}
		clent := mapper.ClientToDTO(&clientModel)
		a.responcer.SendClient(&clent)
	}
}

func (a *AdvertisementController) AllOrders() {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllOrders")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	orders, err := a.req.AllOrders(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllOrders: cannot get orders", "error", err)
		a.handleError(err)
		return
	}

	for i := range orders {
		orderModel, err := mapper.DtoToOrder(&orders[i])
		if err != nil {
			logging.Logger.Error("order dont pass the filter", "orderID", orders[i].ID, "error", err)
			continue
		}
		order := mapper.OrderToDTO(&orderModel)
		a.responcer.SendAdvertisementsOrder(&order)
	}
}

func (a *AdvertisementController) AllBlockAdvertisements() {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllBlockAdvertisements")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.AllBlockAdvertisements(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllBlockAdvertisements: cannot get all BlockAdvertisements", "error", err)
		a.handleError(err)
		return
	}

	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		blockAdv := mapper.BlockAdvertisementToDTO(&blockModel)
		a.responcer.SendBlockAdvertisement(&blockAdv)
	}
}

func (a *AdvertisementController) AllLineAdvertisements() {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllLineAdvertisements")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.AllLineAdvertisements(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllLineAdvertisements: cannot get all LineAdvertisements", "error", err)
		a.handleError(err)
		return
	}

	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		lineAdv := mapper.LineAdvertisementToDTO(&lineModel)
		a.responcer.SendLineAdvertisement(&lineAdv)
	}
}

func (a *AdvertisementController) OrdersByClientName(name string) {
	logging.Logger.Debug("advertisementhandler: search request. Search Order by Client name")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	orders, err := a.req.OrdersByClientName(ctx, name)
	if err != nil {
		logging.Logger.Error("advertisementhandler.OrdersByClientName: cannot get OrdersByClientName", "error", err)
		a.handleError(err)
		return
	}

	for i := range orders {
		orderModel, err := mapper.DtoToOrder(&orders[i])
		if err != nil {
			logging.Logger.Error("order dont pass the filter", "orderID", orders[i].ID, "error", err)
			continue
		}
		order := mapper.OrderToDTO(&orderModel)
		a.responcer.SendAdvertisementsOrder(&order)
	}
}

func (a *AdvertisementController) BlockAdvertisementsByOrderID(id int) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements by Order id", "id", id)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.BlockAdvertisementsByOrderID(ctx, id)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsByOrderID: cannot get BlockAdvertisementsByOrderID", "error", err)
		a.handleError(err)
		return
	}

	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		blockAdv := mapper.BlockAdvertisementToDTO(&blockModel)
		a.responcer.SendBlockAdvertisement(&blockAdv)
	}
}

func (a *AdvertisementController) LineAdvertisementsByOrderID(id int) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements by Order id", "id", id)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.LineAdvertisementsByOrderID(ctx, id)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsByOrderID: cannot get LineAdvertisementsByOrderID", "error", err)
		a.handleError(err)
		return
	}

	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		lineAdv := mapper.LineAdvertisementToDTO(&lineModel)
		a.responcer.SendLineAdvertisement(&lineAdv)
	}
}

func (a *AdvertisementController) BlockAdvertisementsBetweenReleaseDates(from, to time.Time) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements Between release dates", "from", from, "to", to)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.BlockAdvertisementBetweenReleaseDates(ctx, from, to)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsBetweenReleaseDates: cannot get BlockAdvertisementsBetweenReleaseDates", "error", err)
		a.handleError(err)
		return
	}

	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		blockAdv := mapper.BlockAdvertisementToDTO(&blockModel)
		a.responcer.SendBlockAdvertisement(&blockAdv)
	}
}

func (a *AdvertisementController) LineAdvertisementsBetweenReleaseDates(from, to time.Time) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements Between release dates", "from", from, "to", to)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.LineAdvertisementBetweenReleaseDates(ctx, from, to)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsBetweenReleaseDates: cannot get LineAdvertisementsBetweenReleaseDates", "error", err)
		a.handleError(err)
		return
	}

	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		lineAdv := mapper.LineAdvertisementToDTO(&lineModel)
		a.responcer.SendLineAdvertisement(&lineAdv)
	}
}

func (a *AdvertisementController) BlockAdvertisementsActualReleaseDate() {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements with actual release dates", "from", time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.BlockAdvertisementFromReleaseDates(ctx, time.Now())
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsActualReleaseDate: cannot get BlockAdvertisementsActualReleaseDate", "error", err)

		a.handleError(err)
		return
	}

	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		blockAdv := mapper.BlockAdvertisementToDTO(&blockModel)
		a.responcer.SendBlockAdvertisement(&blockAdv)
	}
}

func (a *AdvertisementController) LineAdvertisementsActualReleaseDate() {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements with actual release dates", "from", time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.LineAdvertisementFromReleaseDates(ctx, time.Now())
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsActualReleaseDate: cannot get LineAdvertisementsActualReleaseDate", "error", err)
		a.handleError(err)
		return
	}

	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		lineAdv := mapper.LineAdvertisementToDTO(&lineModel)
		a.responcer.SendLineAdvertisement(&lineAdv)
	}
}

func (a *AdvertisementController) BlockAdvertisementsFromReleaseDate(from time.Time) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements from release dates", "from", from)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.BlockAdvertisementFromReleaseDates(ctx, from)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsFromReleaseDate: cannot get BlockAdvertisementsFromReleaseDate", "error", err)
		a.handleError(err)
		return
	}

	for i := range blockAdvs {
		blockModel, err := mapper.DtoToAdvertisementBlock(&blockAdvs[i])
		if err != nil {
			logging.Logger.Error("blockAdvertisement dont pass the filter", "id", blockAdvs[i].ID, "error", err)
			continue
		}
		blockAdv := mapper.BlockAdvertisementToDTO(&blockModel)
		a.responcer.SendBlockAdvertisement(&blockAdv)
	}
}

func (a *AdvertisementController) LineAdvertisementsFromReleaseDate(from time.Time) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements from release dates", "from", from)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.LineAdvertisementFromReleaseDates(ctx, from)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsFromReleaseDate: cannot get LineAdvertisementsFromReleaseDate", "error", err)
		a.handleError(err)
		return
	}

	for i := range lineAdvs {
		lineModel, err := mapper.DtoToAdvertisementLine(&lineAdvs[i])
		if err != nil {
			logging.Logger.Error("lineAdvertisement dont pass the filter", "id", lineAdvs[i].ID, "error", err)
			continue
		}
		lineAdv := mapper.LineAdvertisementToDTO(&lineModel)
		a.responcer.SendLineAdvertisement(&lineAdv)
	}
}

func (a *AdvertisementController) AllTags() {
	logging.Logger.Debug("advertisementhandler: search request. Search all Tags")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	tags, err := a.req.AllTags(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllTags: cannot get AllTags", "error", err)
		a.handleError(err)
		return
	}

	for i := range tags {
		tagModel, err := mapper.DtoToTag(&tags[i])
		if err != nil {
			logging.Logger.Error("tag dont pass the filter", "name", tags[i].TagName, "error", err)
			continue
		}
		tag := mapper.TagToDTO(&tagModel)
		a.responcer.SendTag(&tag)
	}
}

func (a *AdvertisementController) AllExtraCharges() {
	logging.Logger.Debug("advertisementhandler: search request. Search all ExtraCharges")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	extraCharges, err := a.req.AllExtraCharges(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllExtraCharges: cannot get AllExtraCharges", "error", err)
		a.handleError(err)
		return
	}

	for i := range extraCharges {
		extraChargeModel, err := mapper.DtoToExtraCharge(&extraCharges[i])
		if err != nil {
			logging.Logger.Error("extraCharge dont pass the filter", "name", extraCharges[i].ChargeName, "error", err)
			continue
		}
		extraCharge := mapper.ExtraChargeToDTO(&extraChargeModel)
		a.responcer.SendExtraCharge(&extraCharge)
	}
}

func (a *AdvertisementController) AllCostRates() {
	logging.Logger.Debug("advertisementhandler: search request. Search all CostRates")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	costRates, err := a.req.AllCostRates(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllCostRates: cannot get AllCostRates", "error", err)
		a.handleError(err)
		return
	}

	for i := range costRates {
		costRateModel, err := mapper.DtoToCostRate(&costRates[i])
		if err != nil {
			logging.Logger.Error("costRate dont pass the filter", "name", costRates[i].Name, "error", err)
			continue
		}
		costRate := mapper.CostRateToDTO(&costRateModel)
		a.responcer.SendCostRate(&costRate)
	}
}
