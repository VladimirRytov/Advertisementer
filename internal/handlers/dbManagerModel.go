package handlers

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type DBManager interface {
	AvailableLocalDatabases()
	AvailableNetworkDatabases()
	ConnectToDatabase(string, []byte)
	ConnectToServer(string, []byte)
	DefaultPort(string)
	DatabaseConnected() bool
}

type RequestHandler interface {
	DBManager
	HandlerRequests
	ConfigsAccesor
	JSONExporter
	JSONImporter
	CostCalculator
	Reports
}
type ConfigsAccesor interface {
	Load(string) ([]byte, error)
	SaveConfig(string, []byte) error
	Remove(string) error
}

type CostCalculator interface {
	InitCalculator()
	SetActiveCostRate(string) error
	ActiveCostRate()
	CalculateBlockAdvertisementCost(datatransferobjects.BlockAdvertisementDTO)
	CalculateLineAdvertisementCost(datatransferobjects.LineAdvertisementDTO)
	CalculateOrderCost(datatransferobjects.OrderDTO)
	SelectedCostRate()
}

type JSONExporter interface {
	ExportJsonToFile(context.Context, string) error
}

type JSONImporter interface {
	ImportJson(context.Context, string, datatransferobjects.ImportParams) error
}

type Reports interface {
	GenerateReport(context.Context, datatransferobjects.ReportParams)
}
