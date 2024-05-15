package presenter

import (
	"strconv"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (dc *DataConverter) ClientToViewDTO(cl *datatransferobjects.ClientDTO) ClientDTO {
	logging.Logger.Debug("converter: Converting Client DTO to View")
	return ClientDTO{
		Name:                  cl.Name,
		Email:                 cl.Email,
		Phones:                cl.Phones,
		AdditionalInformation: cl.AdditionalInformation,
	}
}

func (dc *DataConverter) OrderToViewDTO(or *datatransferobjects.OrderDTO) OrderDTO {
	logging.Logger.Debug("converter: Converting Order DTO to View")
	return OrderDTO{
		ID:            or.ID,
		ClientName:    or.ClientName,
		Cost:          dc.CostToView(int(or.Cost)),
		PaymentType:   or.PaymentType,
		CreatedDate:   dc.DateToView(or.CreatedDate),
		PaymentStatus: or.PaymentStatus,
	}
}

func (dc *DataConverter) BlockAdvertisementToViewDTO(block *datatransferobjects.BlockAdvertisementDTO) BlockAdvertisementDTO {
	var closestStr string
	logging.Logger.Debug("converter: Converting BlockAdvertisement DTO to View")
	releaseDates := make([]string, 0, len(block.ReleaseDates))
	for _, v := range block.ReleaseDates {
		releaseDates = append(releaseDates, dc.DateToView(v))
	}
	closestTime, err := dc.ClosestRelease(block.ReleaseDates, time.Now())
	if err == nil {
		closestStr = dc.DateToView(closestTime)
	}
	return BlockAdvertisementDTO{
		Advertisement: Advertisement{
			ID:             block.ID,
			OrderID:        block.OrderID,
			ReleaseCount:   int(block.ReleaseCount),
			ClosestRelease: closestStr,
			ReleaseDates:   dc.ArrayToString(releaseDates),
			Cost:           dc.CostToView(int(block.Cost)),
			Text:           block.Text,
			Tags:           dc.ArrayToString(block.Tags),
			ExtraCharge:    dc.ArrayToString(block.ExtraCharges),
		},
		Size:     int(block.Size),
		FileName: block.FileName,
	}
}

func (dc *DataConverter) LineAdvertisementToViewDTO(line *datatransferobjects.LineAdvertisementDTO) LineAdvertisementDTO {
	var closestStr string
	logging.Logger.Debug("converter: Converting LineAdvertisement DTO to View")
	releaseDates := make([]string, 0, len(line.ReleaseDates))
	for _, v := range line.ReleaseDates {
		releaseDates = append(releaseDates, dc.DateToView(v))
	}
	closestTime, err := dc.ClosestRelease(line.ReleaseDates, time.Now())
	if err == nil {
		closestStr = dc.DateToView(closestTime)
	}
	return LineAdvertisementDTO{
		Advertisement: Advertisement{
			ID:             line.ID,
			OrderID:        line.OrderID,
			ReleaseCount:   int(line.ReleaseCount),
			ClosestRelease: closestStr,
			ReleaseDates:   dc.ArrayToString(releaseDates),
			Cost:           dc.CostToView(int(line.Cost)),
			Text:           line.Text,
			Tags:           dc.ArrayToString(line.Tags),
			ExtraCharge:    dc.ArrayToString(line.ExtraCharges),
		}}
}

func (dc *DataConverter) TagToViewDTO(tag *datatransferobjects.TagDTO) TagDTO {
	logging.Logger.Debug("converter: Converting Tag DTO to View")
	return TagDTO{
		TagName: tag.TagName,
		TagCost: dc.CostToView(int(tag.TagCost)),
	}
}

func (dc *DataConverter) ExtraChargeToViewDTO(charge *datatransferobjects.ExtraChargeDTO) ExtraChargeDTO {
	logging.Logger.Debug("converter: Converting ExtraCharge DTO to View")
	return ExtraChargeDTO{
		ChargeName: charge.ChargeName,
		Multiplier: strconv.Itoa(charge.Multiplier),
	}
}

func (dc *DataConverter) LocalDsnToView(loc *datatransferobjects.LocalDSN) LocalDSN {
	logging.Logger.Debug("converter: Converting LocalDSN DTO to View")
	return LocalDSN{
		Path: loc.Path,
		Name: loc.Name,
		Type: loc.Type,
	}
}

func (dc *DataConverter) NetworkDsnToView(net *datatransferobjects.NetworkDataBaseDSN) NetworkDataBaseDSN {
	logging.Logger.Debug("converter: Converting NetworkDSN DTO to View")
	return NetworkDataBaseDSN{
		Source:   net.Source,
		DataBase: net.DataBase,
		UserName: net.UserName,
		Password: net.Password,
		SSLMode:  net.SSLMode,
		Port:     strconv.Itoa(int(net.Port)),
	}
}

func (dc *DataConverter) CostRateToViewDTO(costRate *datatransferobjects.CostRateDTO) CostRateDTO {
	logging.Logger.Debug("converter: Converting Tag DTO to View")
	return CostRateDTO{
		Name:            costRate.Name,
		OneWordOrSymbol: dc.CostToView(int(costRate.ForOneWordSymbol)),
		Onecm2:          dc.CostToView(int(costRate.ForOneSquare)),
		CalcForOneWord:  costRate.CalcForOneWord,
	}
}

func (dc *DataConverter) FileToViewDTO(file *datatransferobjects.FileDTO) File {
	logging.Logger.Debug("converter: Converting File DTO to View")
	return File{
		Name: file.Name,
		Size: dc.calcFileSize(int(file.Size)),
		Data: file.Data,
	}
}
