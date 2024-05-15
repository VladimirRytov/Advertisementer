package requestscontroller

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
)

type ToDtoConverter interface {
	TagToDto(*presenter.TagDTO) (datatransferobjects.TagDTO, error)
	ExtraChargeToDto(*presenter.ExtraChargeDTO) (datatransferobjects.ExtraChargeDTO, error)
	ClientToDto(*presenter.ClientDTO) (datatransferobjects.ClientDTO, error)
	OrderToDto(*presenter.OrderDTO) (datatransferobjects.OrderDTO, error)
	BlockAdvertisementToDto(*presenter.BlockAdvertisementDTO) (datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementToDto(*presenter.LineAdvertisementDTO) (datatransferobjects.LineAdvertisementDTO, error)
	CostRateToDTO(*presenter.CostRateDTO) (datatransferobjects.CostRateDTO, error)
	DateToDto(string) (time.Time, error)
	CostToDto(string) (int, error)
	NetworkDsnToDto(*presenter.NetworkDataBaseDSN) (datatransferobjects.NetworkDataBaseDSN, error)
	LocalDsnToDto(*presenter.LocalDSN) (datatransferobjects.LocalDSN, error)
	ServerDsnToDto(*presenter.ServerDSN) (datatransferobjects.ServerDSN, error)
	FileToDto(*presenter.File) datatransferobjects.FileDTO
	AdvReportToDTO(report *presenter.ReportParams) (datatransferobjects.ReportParams, error)
}

type DataComparer interface {
	ClosestRelease([]time.Time, time.Time) (time.Time, error)
	ParsePath(string) string
}

type DataHandler interface {
	ArrayToString(arr []string) string
	YearMonthDayToString(year, month, day uint) string
	ToDtoConverter
	DataComparer
}
