package mapper

import (
	"github.com/VladimirRytov/advertisementer/internal/advertisements"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func DtoToClient(cl *datatransferobjects.ClientDTO) (advertisements.Client, error) {
	logging.Logger.Debug("mapper: Creating Client entity", "Client", cl)
	client, err := advertisements.NewClient(cl.Name)
	if err != nil {
		return advertisements.Client{}, err
	}
	err = client.SetContactInformation(cl.Phones, cl.Email)
	if err != nil {
		return advertisements.Client{}, err
	}
	client.SetAdditionalInformation(cl.AdditionalInformation)
	return client, nil
}

func DtoToOrder(cl *datatransferobjects.OrderDTO) (advertisements.AdvertisementOrder, error) {
	logging.Logger.Debug("mapper: Creating Order entity", "Order", cl)
	order, err := advertisements.NewAdvertisementOrder(cl.ClientName)
	if err != nil {
		return advertisements.AdvertisementOrder{}, err
	}
	order.SetOrderId(cl.ID)
	order.SetCreaatedDate(cl.CreatedDate)
	order.SetPaymentStatus(cl.PaymentStatus)
	order.SetPaymentType(cl.PaymentType)
	order.SetOrderCost(cl.Cost)
	return order, nil
}

func DtoToAdvertisementBlock(cl *datatransferobjects.BlockAdvertisementDTO) (advertisements.AdvertisementBlock, error) {
	logging.Logger.Debug("mapper: Creating BlockAdvertisement entity", "BlockAdvertisement", cl)
	blockAdv := advertisements.NewAdvertisementBlock()
	err := blockAdv.SetId(cl.ID)
	if err != nil {
		return advertisements.AdvertisementBlock{}, err
	}

	err = blockAdv.SetOrderId(cl.OrderID)
	if err != nil {
		return advertisements.AdvertisementBlock{}, err
	}

	err = blockAdv.SetSize(cl.Size)
	if err != nil {
		return advertisements.AdvertisementBlock{}, err
	}

	err = blockAdv.SetFileName(cl.FileName)
	if err != nil {
		return advertisements.AdvertisementBlock{}, err
	}

	err = blockAdv.SetCost(cl.Cost)
	if err != nil {
		return advertisements.AdvertisementBlock{}, err
	}
	err = blockAdv.SetReleaseDates(cl.ReleaseDates)
	if err != nil {
		return advertisements.AdvertisementBlock{}, err
	}
	blockAdv.SetComment(cl.Text)
	blockAdv.SetExtraCharges(cl.ExtraCharges)
	blockAdv.SetTags(cl.Tags)
	return blockAdv, nil
}

func DtoToAdvertisementLine(cl *datatransferobjects.LineAdvertisementDTO) (advertisements.AdvertisementLine, error) {
	logging.Logger.Debug("mapper: Creating LineAdvertisement entity", "LineAdvertisement", cl)
	lineAdv := advertisements.NewAdvertisementLine()
	err := lineAdv.SetId(cl.ID)
	if err != nil {
		return advertisements.AdvertisementLine{}, err
	}

	err = lineAdv.SetOrderId(cl.OrderID)
	if err != nil {
		return advertisements.AdvertisementLine{}, err
	}

	lineAdv.SetCost(cl.Cost)
	lineAdv.SetContent(cl.Text)

	lineAdv.SetExtraCharges(cl.ExtraCharges)
	lineAdv.SetTags(cl.Tags)
	lineAdv.SetReleaseDates(cl.ReleaseDates)

	return lineAdv, nil
}

func DtoToTag(cl *datatransferobjects.TagDTO) (advertisements.Tag, error) {
	logging.Logger.Debug("mapper: Creating Tag entity", "Tag", cl)
	return advertisements.NewTag(cl.TagName, cl.TagCost)
}

func DtoToExtraCharge(cl *datatransferobjects.ExtraChargeDTO) (advertisements.ExtraCharge, error) {
	logging.Logger.Debug("mapper: Creating ExtraCharge entity", "ExtraCharge", cl)
	return advertisements.NewExtraCharge(cl.ChargeName, cl.Multiplier)
}

func DtoToCostRate(cr *datatransferobjects.CostRateDTO) (advertisements.CostRate, error) {
	logging.Logger.Debug("mapper: Creating CostRate entity", "CostRate", cr)
	return advertisements.NewCostRate(cr.Name, cr.ForOneWordSymbol, cr.ForOneSquare, cr.CalcForOneWord)
}
