package server

import (
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (ds *ServerStorage) convertClientToDTO(client *ClientFront) datatransferobjects.ClientDTO {
	logging.Logger.Debug("orm: Converting Client databaseObject to DTO")

	return datatransferobjects.ClientDTO{
		Name:                  client.Name,
		Phones:                client.Phones,
		Email:                 client.Email,
		AdditionalInformation: client.AdditionalInformation,
	}
}

func (ds *ServerStorage) convertClientToModel(client *datatransferobjects.ClientDTO) ClientFront {
	logging.Logger.Debug("orm: Converting Client DTO to databaseObject")
	return ClientFront{
		Name:                  client.Name,
		Phones:                client.Phones,
		Email:                 client.Email,
		AdditionalInformation: client.AdditionalInformation,
	}
}

func (ds *ServerStorage) convertOrderToDTO(order *OrderFront) datatransferobjects.OrderDTO {
	logging.Logger.Debug("orm: Converting Order databaseObject to DTO")
	var (
		blocks []datatransferobjects.BlockAdvertisementDTO
		lines  []datatransferobjects.LineAdvertisementDTO
	)
	for i := range order.BlockAdvertisements {
		blocks = append(blocks, ds.convertBlockAdvertisementToDTO(&order.BlockAdvertisements[i]))
	}
	for i := range order.LineAdvertisements {
		lines = append(lines, ds.convertLineAdvertisementToDTO(&order.LineAdvertisements[i]))
	}

	return datatransferobjects.OrderDTO{
		ID:                  order.ID,
		ClientName:          order.ClientName,
		Cost:                order.Cost,
		CreatedDate:         order.CreatedDate,
		PaymentType:         order.PaymentType,
		PaymentStatus:       order.PaymentStatus,
		BlockAdvertisements: blocks,
		LineAdvertisements:  lines,
	}
}

func (ds *ServerStorage) convertOrderToModel(order *datatransferobjects.OrderDTO) OrderFront {
	logging.Logger.Debug("orm: Converting Order DTO to databaseObject")
	var (
		blocks []BlockAdvertisementFront
		lines  []LineAdvertisementFront
	)
	for i := range order.BlockAdvertisements {
		blocks = append(blocks, ds.convertBlockAdvertisementToModel(&order.BlockAdvertisements[i]))
	}
	for i := range order.LineAdvertisements {
		lines = append(lines, ds.convertLineAdvertisementToModel(&order.LineAdvertisements[i]))
	}
	return OrderFront{
		ID:                  order.ID,
		ClientName:          order.ClientName,
		Cost:                order.Cost,
		CreatedDate:         order.CreatedDate,
		PaymentType:         order.PaymentType,
		PaymentStatus:       order.PaymentStatus,
		BlockAdvertisements: blocks,
		LineAdvertisements:  lines,
	}
}

func (ds *ServerStorage) convertBlockAdvertisementToDTO(blockAdv *BlockAdvertisementFront) datatransferobjects.BlockAdvertisementDTO {
	logging.Logger.Debug("orm: Converting BlockAdvertisement databaseObject to DTO")
	return datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           blockAdv.ID,
			OrderID:      blockAdv.OrderID,
			ReleaseCount: blockAdv.ReleaseCount,
			Cost:         blockAdv.Cost,
			Text:         blockAdv.Text,
			Tags:         blockAdv.Tags,
			ExtraCharges: blockAdv.ExtraCharges,
			ReleaseDates: blockAdv.ReleaseDates,
		},
		Size:     blockAdv.Size,
		FileName: blockAdv.FileName,
	}
}
func (ds *ServerStorage) convertBlockAdvertisementToModel(blockAdv *datatransferobjects.BlockAdvertisementDTO) BlockAdvertisementFront {
	logging.Logger.Debug("orm: Converting BlockAdvertisement DTO to databaseObject")
	return BlockAdvertisementFront{
		Advertisement: Advertisement{
			ID:           blockAdv.ID,
			OrderID:      blockAdv.OrderID,
			ReleaseCount: blockAdv.ReleaseCount,
			Cost:         blockAdv.Cost,
			Text:         blockAdv.Text,
			Tags:         blockAdv.Tags,
			ExtraCharges: blockAdv.ExtraCharges,
			ReleaseDates: blockAdv.ReleaseDates,
		},
		Size:     blockAdv.Size,
		FileName: blockAdv.FileName,
	}
}

func (ds *ServerStorage) convertLineAdvertisementToDTO(lineAdv *LineAdvertisementFront) datatransferobjects.LineAdvertisementDTO {
	logging.Logger.Debug("orm: Converting LineAdvertisement databaseObject to DTO")
	return datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           lineAdv.ID,
			OrderID:      lineAdv.OrderID,
			ReleaseCount: lineAdv.ReleaseCount,
			Cost:         lineAdv.Cost,
			Text:         lineAdv.Text,
			Tags:         lineAdv.Tags,
			ExtraCharges: lineAdv.ExtraCharges,
			ReleaseDates: lineAdv.ReleaseDates,
		},
	}
}

func (ds *ServerStorage) convertLineAdvertisementToModel(lineAdv *datatransferobjects.LineAdvertisementDTO) LineAdvertisementFront {
	logging.Logger.Debug("orm: Converting LineAdvertisement DTO to databaseObject")
	return LineAdvertisementFront{
		Advertisement: Advertisement{
			ID:           lineAdv.ID,
			OrderID:      lineAdv.OrderID,
			ReleaseCount: lineAdv.ReleaseCount,
			Cost:         lineAdv.Cost,
			Text:         lineAdv.Text,
			Tags:         lineAdv.Tags,
			ExtraCharges: lineAdv.ExtraCharges,
			ReleaseDates: lineAdv.ReleaseDates,
		},
	}
}

func (ds *ServerStorage) convertTagToDTO(tag *TagFront) datatransferobjects.TagDTO {
	logging.Logger.Debug("orm: Converting Tag databaseObject to DTO")
	return datatransferobjects.TagDTO{
		TagName: tag.TagName,
		TagCost: tag.TagCost,
	}
}

func (ds *ServerStorage) convertTagToModel(tag *datatransferobjects.TagDTO) TagFront {
	logging.Logger.Debug("orm: Converting Tag DTO to databaseObject")
	return TagFront{
		TagName: tag.TagName,
		TagCost: tag.TagCost,
	}
}

func (ds *ServerStorage) convertExtraChargeToDTO(extraCharge *ExtraChargeFront) datatransferobjects.ExtraChargeDTO {
	logging.Logger.Debug("orm: Converting ExtraCharge databaseObject to DTO")
	return datatransferobjects.ExtraChargeDTO{
		ChargeName: extraCharge.ChargeName,
		Multiplier: extraCharge.Multiplier,
	}
}

func (ds *ServerStorage) convertExtraChargeToModel(extraCharge *datatransferobjects.ExtraChargeDTO) ExtraChargeFront {
	logging.Logger.Debug("orm: Converting ExtraCharge DTO to databaseObject")
	return ExtraChargeFront{
		ChargeName: extraCharge.ChargeName,
		Multiplier: extraCharge.Multiplier,
	}
}

func (ds *ServerStorage) convertCostRateToModel(costRate *datatransferobjects.CostRateDTO) CostRateFront {
	logging.Logger.Debug("orm: Converting ExtraCharge DTO to front object")
	return CostRateFront{
		Name:             costRate.Name,
		ForOneWordSymbol: costRate.ForOneWordSymbol,
		ForOnecm2:        costRate.ForOneSquare,
		CalcForOneWord:   costRate.CalcForOneWord,
	}
}

func (ds *ServerStorage) convertCostRateToDto(costRate *CostRateFront) datatransferobjects.CostRateDTO {
	logging.Logger.Debug("orm: Converting CostRate DTO to databaseObject")
	return datatransferobjects.CostRateDTO{
		Name:             costRate.Name,
		ForOneWordSymbol: costRate.ForOneWordSymbol,
		ForOneSquare:     costRate.ForOnecm2,
		CalcForOneWord:   costRate.CalcForOneWord,
	}
}

func (ds *ServerStorage) convertFileToDto(file *FileFront) (datatransferobjects.FileDTO, error) {
	logging.Logger.Debug("orm: Converting File DTO to databaseObject")
	decodedData, err := ds.b64.FromBase64String(file.Data)
	if err != nil {
		return datatransferobjects.FileDTO{}, err
	}
	return datatransferobjects.FileDTO{
		Name: file.Name,
		Size: file.Size,
		Data: decodedData,
	}, nil
}

func (ds *ServerStorage) convertFileToModel(file *datatransferobjects.FileDTO) FileFront {
	logging.Logger.Debug("orm: Converting File DTO to front object")
	return FileFront{
		Name: file.Name,
		Size: file.Size,
		Data: ds.b64.ToBase64String(file.Data),
	}
}

func (ds *ServerStorage) filesFrontToStringArray(files []FileFront) []string {
	filesName := make([]string, 0, len(files))
	for i := range files {
		filesName = append(filesName, files[i].Name)
	}
	return filesName
}

func (ds *ServerStorage) fileFrontToString(files []FileFront) string {
	if len(files) == 0 {
		return ""
	}
	return files[0].Name
}
