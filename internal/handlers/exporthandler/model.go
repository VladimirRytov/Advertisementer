package exporthandler

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
	SendError(error)
	SendMessage(string)
	ProgressComplete()
}

type DataBase interface {
	Getter
	Searcher
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

type Searcher interface {
	OrdersByClientName(context.Context, string) ([]datatransferobjects.OrderDTO, error)
	BlockAdvertisementsByOrderID(context.Context, int) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementsByOrderID(context.Context, int) ([]datatransferobjects.LineAdvertisementDTO, error)
	BlockAdvertisementBetweenReleaseDates(context.Context, time.Time, time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementBetweenReleaseDates(context.Context, time.Time, time.Time) ([]datatransferobjects.LineAdvertisementDTO, error)

	BlockAdvertisementFromReleaseDates(context.Context, time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementFromReleaseDates(context.Context, time.Time) ([]datatransferobjects.LineAdvertisementDTO, error)

	AllTags(context.Context) ([]datatransferobjects.TagDTO, error)
	AllExtraCharges(context.Context) ([]datatransferobjects.ExtraChargeDTO, error)

	AllClients(context.Context) ([]datatransferobjects.ClientDTO, error)
	AllOrders(context.Context) ([]datatransferobjects.OrderDTO, error)

	AllLineAdvertisements(context.Context) ([]datatransferobjects.LineAdvertisementDTO, error)
	AllBlockAdvertisements(context.Context) ([]datatransferobjects.BlockAdvertisementDTO, error)
	AllCostRates(context.Context) ([]datatransferobjects.CostRateDTO, error)
}
