package mapper

import (
	"github.com/VladimirRytov/advertisementer/internal/advertisements"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

type ConvertError struct {
	EntityID string
}

func (e *ConvertError) Error() string {
	return e.EntityID
}

func ClientToDTO(cl *advertisements.Client) datatransferobjects.ClientDTO {
	logging.Logger.Debug("mapper: Creating Client DTO")
	return datatransferobjects.ClientDTO{
		Name:                  cl.Name(),
		Phones:                cl.ContactNumber(),
		Email:                 cl.Email(),
		AdditionalInformation: cl.AdditionalInformation(),
	}
}

func OrderToDTO(cl *advertisements.AdvertisementOrder) datatransferobjects.OrderDTO {
	logging.Logger.Debug("mapper: Creating Order DTO")
	return datatransferobjects.OrderDTO{
		ID:            cl.OrderId(),
		ClientName:    cl.ClientId(),
		Cost:          cl.Cost(),
		PaymentType:   cl.PaymentType(),
		CreatedDate:   cl.CreaatedDate(),
		PaymentStatus: cl.PaymentStatus(),
	}
}

func BlockAdvertisementToDTO(blk *advertisements.AdvertisementBlock) datatransferobjects.BlockAdvertisementDTO {
	logging.Logger.Debug("mapper: Creating BlockAdvertisement DTO")
	return datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           blk.Id(),
			OrderID:      blk.OrderId(),
			ReleaseCount: blk.ReseaseCount(),
			Cost:         blk.Cost(),
			Text:         blk.Comment(),
			Tags:         blk.Tags(),
			ExtraCharges: blk.ExtraCharge(),
			ReleaseDates: blk.ReleaseDates(),
		},
		Size:     blk.Size(),
		FileName: blk.FileName(),
	}
}

func LineAdvertisementToDTO(cl *advertisements.AdvertisementLine) datatransferobjects.LineAdvertisementDTO {
	logging.Logger.Debug("mapper: Creating LineAdvertisement DTO")
	return datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           cl.Id(),
			OrderID:      cl.OrderId(),
			ReleaseCount: cl.ReseaseCount(),
			Cost:         cl.Cost(),
			Text:         cl.Content(),
			Tags:         cl.Tags(),
			ExtraCharges: cl.ExtraCharge(),
			ReleaseDates: cl.ReleaseDates(),
		},
	}
}

func TagToDTO(cl *advertisements.Tag) datatransferobjects.TagDTO {
	logging.Logger.Debug("mapper: Creating Tag DTO")
	return datatransferobjects.TagDTO{
		TagName: cl.Name(),
		TagCost: cl.Cost(),
	}
}

func ExtraChargeToDTO(cl *advertisements.ExtraCharge) datatransferobjects.ExtraChargeDTO {
	logging.Logger.Debug("mapper: Creating ExtraCharge DTO")
	return datatransferobjects.ExtraChargeDTO{
		ChargeName: cl.Name(),
		Multiplier: cl.Multiplier(),
	}
}

func CostRateToDTO(cr *advertisements.CostRate) datatransferobjects.CostRateDTO {
	logging.Logger.Debug("mapper: Creating CostRate DTO")
	return datatransferobjects.CostRateDTO{
		CalcForOneWord:   cr.CalsForOneWord(),
		Name:             cr.Name(),
		ForOneWordSymbol: cr.CostForWordOrSymbol(),
		ForOneSquare:     cr.CostForOnecm2(),
	}
}
