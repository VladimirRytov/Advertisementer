package thincostratecalculator

import (
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type Responcer interface {
	SendError(error)
	SendActiveCostRate(string)
	SendBlockAdvertisementCost(datatransferobjects.BlockAdvertisementDTO)
	SendLineAdvertisementCost(datatransferobjects.LineAdvertisementDTO)
	SendOrderCost(datatransferobjects.OrderDTO)
}
