package handlers

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type ThinCostRateCalculator interface {
	InitCostRateCalculator(Responcer, CostRequests)
	CostRateCalculator
}

type ThickCostRateCalculator interface {
	InitCostRateCalculator(CostRateResponcer, ThinCostRateRequests)
	CostRateCalculator
}

type CostRateCalculator interface {
	SetActiveCostRate(string) error
	ActiveCostRate()
	CalculateBlockAdvertisementCost(datatransferobjects.BlockAdvertisementDTO)
	CalculateLineAdvertisementCost(datatransferobjects.LineAdvertisementDTO)
	CalculateOrderCost(datatransferobjects.OrderDTO)
	SelectedCostRate()
}

type CostRateResponcer interface {
	SendError(error)
	SendActiveCostRate(string)
	SendBlockAdvertisementCost(datatransferobjects.BlockAdvertisementDTO)
	SendLineAdvertisementCost(datatransferobjects.LineAdvertisementDTO)
	SendOrderCost(datatransferobjects.OrderDTO)
}

type ThinCostRateRequests interface {
	CostRateGetter
	CostRateSearcher
	Closer
}

type CostRateGetter interface {
	ClientByName(context.Context, string) (datatransferobjects.ClientDTO, error)
	OrderByID(context.Context, int) (datatransferobjects.OrderDTO, error)
	LineAdvertisementByID(context.Context, int) (datatransferobjects.LineAdvertisementDTO, error)
	BlockAdvertisementByID(context.Context, int) (datatransferobjects.BlockAdvertisementDTO, error)
	TagByName(context.Context, string) (datatransferobjects.TagDTO, error)
	ExtraChargeByName(context.Context, string) (datatransferobjects.ExtraChargeDTO, error)
	CostRateByName(context.Context, string) (datatransferobjects.CostRateDTO, error)
}

type CostRateSearcher interface {
	OrdersByClientName(context.Context, string) ([]datatransferobjects.OrderDTO, error)
	BlockAdvertisementsByOrderID(context.Context, int) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementsByOrderID(context.Context, int) ([]datatransferobjects.LineAdvertisementDTO, error)
}
