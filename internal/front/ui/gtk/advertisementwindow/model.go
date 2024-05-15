package advertisementwindow

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
)

type RequestGate interface {
	LoadConfig(string, any) error
	SaveConfig(string, any) error
	StartExportingJson(context.Context, string)
	AllTags()
	AllExtraCharges()
	AllClients()
	AllOrders()
	AllCostRates()
	BlockAdvertisementsActual()
	LineAdvertisementsActual()
	AllBlockAdvertisements()
	AllLineAdvertisements()
}

type Converter interface {
	YearMonthDayToString(uint, uint, uint) string
}

type Application interface {
	Mode() int8
	ConnectionStatus() bool
	CreateProgressWindow() application.ProgressWindow
	CreateImportDataWindow() application.ImportDataWindow
	CreateCostRatesWindow() application.CostRateWindow
	NewAdvertisementReportWindow() application.AdvertisementReportWindow
}
