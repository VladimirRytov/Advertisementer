package handlers

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type DataBaseConnector interface {
	AvailableLocalDatabases() []string
	AvailableNetworkDatabases() []string
	ConnectToDatabase(string, []byte) (DataBase, error)
	ConnectToServer([]byte) (Server, error)
	DefaultPort(string) uint
}

type DataBase interface {
	ConnectToDatabase([]byte) error
	ConnectionInfo() map[string]string
	Creator
	Getter
	Remover
	Updater
	Searcher
	Closer
}

type Server interface {
	DataBase
	CostRequests
	FileRequests
}

type CostRequests interface {
	CalculateOrderCost(ctx context.Context, order datatransferobjects.OrderDTO, costRateName string) (datatransferobjects.OrderDTO, error)
	CalculateBlockAdvertisementCost(ctx context.Context, block datatransferobjects.BlockAdvertisementDTO,
		costRateName string) (datatransferobjects.BlockAdvertisementDTO, error)
	CalculateLineAdvertisementCost(ctx context.Context, line datatransferobjects.LineAdvertisementDTO,
		costRateName string) (datatransferobjects.LineAdvertisementDTO, error)
}

type FileRequests interface {
	AllFiles(context.Context) ([]datatransferobjects.FileDTO, error)
	RemoveFileByName(context.Context, string) error
	NewFile(ctx context.Context, file *datatransferobjects.FileDTO) (string, error)
	FileByName(ctx context.Context, name string) (datatransferobjects.FileDTO, error)
	FileMiniatureByName(ctx context.Context, fileName string, size string) (datatransferobjects.FileDTO, error)
	NewFileUpload(context.Context, datatransferobjects.FileStream) ([]string, error)
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

type Remover interface {
	RemoveClientByName(context.Context, string) error
	RemoveOrderByID(context.Context, int) error
	RemoveLineAdvertisementByID(context.Context, int) error
	RemoveBlockAdvertisementByID(context.Context, int) error
	RemoveTagByName(context.Context, string) error
	RemoveExtraChargeByName(context.Context, string) error
	RemoveCostRateByName(context.Context, string) error
}

type Updater interface {
	UpdateClient(context.Context, *datatransferobjects.ClientDTO) error
	UpdateOrder(context.Context, *datatransferobjects.OrderDTO) error
	UpdateLineAdvertisement(context.Context, *datatransferobjects.LineAdvertisementDTO) error
	UpdateBlockAdvertisement(context.Context, *datatransferobjects.BlockAdvertisementDTO) error
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

type Closer interface {
	Close() error
}
