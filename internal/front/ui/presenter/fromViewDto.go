package presenter

import (
	"errors"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (dc *DataConverter) ClientToDto(cl *ClientDTO) (datatransferobjects.ClientDTO, error) {
	logging.Logger.Debug("presenter: Converting Client View to DTO")
	return datatransferobjects.ClientDTO{
		Name:                  cl.Name,
		Email:                 cl.Email,
		Phones:                cl.Phones,
		AdditionalInformation: cl.AdditionalInformation,
	}, nil
}

func (dc *DataConverter) OrderToDto(or *OrderDTO) (datatransferobjects.OrderDTO, error) {
	var date time.Time
	logging.Logger.Debug("presenter: Converting Order View to DTO")
	cost, err := dc.CostToDto(or.Cost)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	if len(or.CreatedDate) > 0 {
		date, err = dc.DateToDto(or.CreatedDate)
		if err != nil {
			return datatransferobjects.OrderDTO{}, err
		}
	} else {
		date = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	}

	return datatransferobjects.OrderDTO{
		ID:            or.ID,
		ClientName:    or.ClientName,
		Cost:          cost,
		CreatedDate:   date,
		PaymentType:   or.PaymentType,
		PaymentStatus: or.PaymentStatus,
	}, nil

}

func (dc *DataConverter) BlockAdvertisementToDto(block *BlockAdvertisementDTO) (datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("presenter: Converting BlockAdvertisement View to DTO")

	cost, err := dc.CostToDto(block.Cost)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}

	stringReleaseDates := strings.Split(block.ReleaseDates, ", ")
	releaseDates := make([]time.Time, 0, len(stringReleaseDates))
	for _, v := range stringReleaseDates {
		releaseDate, err := dc.DateToDto(v)
		if err != nil {
			continue
		}
		releaseDates = append(releaseDates, releaseDate)
	}

	return datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           block.ID,
			OrderID:      block.OrderID,
			ReleaseCount: int16(block.ReleaseCount),
			Cost:         cost,
			Text:         block.Text,
			ReleaseDates: releaseDates,
			Tags:         dc.splitStr(block.Tags),
			ExtraCharges: dc.splitStr(block.ExtraCharge),
		},
		Size:     int16(block.Size),
		FileName: block.FileName,
	}, nil

}

func (dc *DataConverter) LineAdvertisementToDto(line *LineAdvertisementDTO) (datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("presenter: Converting LineAdvertisement View to DTO")

	cost, err := dc.CostToDto(line.Cost)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, dc.handleError(err)
	}

	stringReleaseDates := strings.Split(line.ReleaseDates, ", ")
	releaseDates := make([]time.Time, 0, len(stringReleaseDates))
	for _, v := range stringReleaseDates {
		releaseDate, err := dc.DateToDto(v)
		if err != nil {
			continue
		}
		releaseDates = append(releaseDates, releaseDate)
	}
	return datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           line.ID,
			OrderID:      line.OrderID,
			ReleaseCount: int16(line.ReleaseCount),
			Cost:         cost,
			Text:         line.Text,
			Tags:         dc.splitStr(line.Tags),
			ExtraCharges: dc.splitStr(line.ExtraCharge),
			ReleaseDates: releaseDates,
		}}, nil

}

func (dc *DataConverter) TagToDto(tag *TagDTO) (datatransferobjects.TagDTO, error) {
	logging.Logger.Debug("converter: Converting Tag View to DTO")
	cost, err := dc.CostToDto(tag.TagCost)
	if err != nil {
		logging.Logger.Debug("converter: Cannot convert TagView to DTO", "error", err)
		return datatransferobjects.TagDTO{}, err
	}
	return datatransferobjects.TagDTO{TagName: tag.TagName, TagCost: cost}, nil
}

func (dc *DataConverter) ExtraChargeToDto(charge *ExtraChargeDTO) (datatransferobjects.ExtraChargeDTO, error) {
	logging.Logger.Debug("presenter: Converting ExtraCharge View to DTO")
	multiplier, err := strconv.Atoi(charge.Multiplier)
	if err != nil {
		return datatransferobjects.ExtraChargeDTO{}, err
	}
	return datatransferobjects.ExtraChargeDTO{ChargeName: charge.ChargeName, Multiplier: multiplier}, nil
}

func (dc *DataConverter) LocalDsnToDto(loc *LocalDSN) (datatransferobjects.LocalDSN, error) {
	logging.Logger.Debug("presenter: Converting LocalDSN View to DTO")
	return datatransferobjects.LocalDSN{
		Path: loc.Path,
		Name: loc.Name,
		Type: loc.Type,
	}, nil
}

func (dc *DataConverter) ServerDsnToDto(server *ServerDSN) (datatransferobjects.ServerDSN, error) {
	port, err := strconv.Atoi(server.Port)
	if err != nil {
		return datatransferobjects.ServerDSN{}, err
	}
	logging.Logger.Debug("presenter: Converting LocalDSN View to DTO")
	return datatransferobjects.ServerDSN{
		Source:   server.Source,
		UserName: server.UserName,
		Password: server.Password,
		Port:     uint(port),
	}, nil
}

func (dc *DataConverter) NetworkDsnToDto(net *NetworkDataBaseDSN) (datatransferobjects.NetworkDataBaseDSN, error) {
	logging.Logger.Debug("presenter: Converting NetworkDSN View to DTO")
	port, err := strconv.Atoi(net.Port)
	if err != nil {
		return datatransferobjects.NetworkDataBaseDSN{}, err
	}

	return datatransferobjects.NetworkDataBaseDSN{
		Source:   net.Source,
		DataBase: net.DataBase,
		UserName: net.UserName,
		Password: net.Password,
		SSLMode:  net.SSLMode,
		Port:     uint(port),
	}, nil
}

func (dc *DataConverter) AdvReportToDTO(report *ReportParams) (datatransferobjects.ReportParams, error) {
	logging.Logger.Debug("converter: Converting AdvertisementReportView to DTO")
	var startDate, endTime time.Time
	if !filepath.IsAbs(report.BlocksFolderPath) {
		report.BlocksFolderPath = "."
	}

	if !filepath.IsAbs(report.DeployPath) {
		return datatransferobjects.ReportParams{}, errors.New("каталог для выгрузки данных должен быть выбран")
	}

	startDate, err := dc.DateToDto(report.FromDate)
	if err != nil {
		startDate = time.UnixMicro(0)
	}
	endTime, err = dc.DateToDto(report.ToDate)
	if err != nil {
		endTime = time.Date(9999, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	return datatransferobjects.ReportParams{
		ReportType:       report.ReportType,
		FromDate:         startDate,
		ToDate:           endTime,
		BlocksFolderPath: report.BlocksFolderPath,
		DeployPath:       report.DeployPath,
	}, nil
}

func (dc *DataConverter) CostRateToDTO(costRate *CostRateDTO) (datatransferobjects.CostRateDTO, error) {
	logging.Logger.Debug("converter: Converting CostRateView to DTO")
	forOne, err := dc.CostToDto(costRate.OneWordOrSymbol)
	if err != nil {
		return datatransferobjects.CostRateDTO{}, err
	}
	oneCM, err := dc.CostToDto(costRate.Onecm2)
	if err != nil {
		return datatransferobjects.CostRateDTO{}, err
	}

	return datatransferobjects.CostRateDTO{
		Name:             costRate.Name,
		ForOneWordSymbol: forOne,
		ForOneSquare:     oneCM,
		CalcForOneWord:   costRate.CalcForOneWord,
	}, nil
}

func (dc *DataConverter) FileToDto(file *File) datatransferobjects.FileDTO {
	logging.Logger.Debug("converter: Converting Tag DTO to View")
	return datatransferobjects.FileDTO{
		Name: file.Name,
		Data: file.Data,
	}
}
