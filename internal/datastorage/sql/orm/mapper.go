package orm

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func convertClientToDTO(client *Client) datatransferobjects.ClientDTO {
	logging.Logger.Debug("orm: Converting Client databaseObject to DTO")
	return datatransferobjects.ClientDTO{
		Name:                  client.Name,
		Phones:                client.Phones,
		Email:                 client.Email,
		AdditionalInformation: client.AdditionalInformation,
	}
}

func convertClientToModel(client *datatransferobjects.ClientDTO) Client {
	logging.Logger.Debug("orm: Converting Client DTO to databaseObject")
	return Client{
		Name:                  client.Name,
		Phones:                client.Phones,
		Email:                 client.Email,
		AdditionalInformation: client.AdditionalInformation,
	}
}

func convertOrderToDTO(order *Order) datatransferobjects.OrderDTO {
	logging.Logger.Debug("orm: Converting Order databaseObject to DTO")
	var (
		blocks []datatransferobjects.BlockAdvertisementDTO
		lines  []datatransferobjects.LineAdvertisementDTO
	)
	for i := range order.AdvertisementBlocks {
		blocks = append(blocks, convertBlockAdvertisementToDTO(&order.AdvertisementBlocks[i]))
	}
	for i := range order.AdvertisementsLines {
		lines = append(lines, convertLineAdvertisementToDTO(&order.AdvertisementsLines[i]))
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

func convertOrderToModel(order *datatransferobjects.OrderDTO) Order {
	logging.Logger.Debug("orm: Converting Order DTO to databaseObject")
	var (
		blocks []AdvertisementBlock
		lines  []AdvertisementLine
	)
	for i := range order.BlockAdvertisements {
		blocks = append(blocks, convertBlockAdvertisementToModel(&order.BlockAdvertisements[i]))
	}
	for i := range order.LineAdvertisements {
		lines = append(lines, convertLineAdvertisementToModel(&order.LineAdvertisements[i]))
	}
	return Order{
		ID:                  order.ID,
		ClientName:          order.ClientName,
		Cost:                order.Cost,
		CreatedDate:         order.CreatedDate,
		PaymentType:         order.PaymentType,
		PaymentStatus:       order.PaymentStatus,
		AdvertisementBlocks: blocks,
		AdvertisementsLines: lines,
	}
}

func convertBlockAdvertisementToDTO(blockAdv *AdvertisementBlock) datatransferobjects.BlockAdvertisementDTO {
	logging.Logger.Debug("orm: Converting BlockAdvertisement databaseObject to DTO", "block", blockAdv)
	return datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           blockAdv.ID,
			OrderID:      blockAdv.OrderID,
			ReleaseCount: blockAdv.ReleaseCount,
			Cost:         blockAdv.Cost,
			Text:         blockAdv.Text,
			Tags:         fetchTagsName(blockAdv.Tags),
			ExtraCharges: fetchExtraChargesName(blockAdv.ExtraCharges),
			ReleaseDates: fetchReleaseDates(blockAdv.ReleaseDates),
		},
		Size:     blockAdv.Size,
		FileName: blockAdv.FileName,
	}
}
func convertBlockAdvertisementToModel(blockAdv *datatransferobjects.BlockAdvertisementDTO) AdvertisementBlock {
	logging.Logger.Debug("orm: Converting BlockAdvertisement DTO to databaseObject")
	return AdvertisementBlock{
		Advertisement: Advertisement{
			ID:           blockAdv.ID,
			OrderID:      blockAdv.OrderID,
			ReleaseCount: blockAdv.ReleaseCount,
			Cost:         blockAdv.Cost,
			Text:         blockAdv.Text,
		},
		Size:         blockAdv.Size,
		Tags:         tagNamesToModel(blockAdv.Tags),
		ExtraCharges: extraChargesToModel(blockAdv.ExtraCharges),
		ReleaseDates: releaseDatesToModel(blockAdv.ReleaseDates),
		FileName:     blockAdv.FileName,
	}
}

func convertLineAdvertisementToDTO(lineAdv *AdvertisementLine) datatransferobjects.LineAdvertisementDTO {
	logging.Logger.Debug("orm: Converting LineAdvertisement databaseObject to DTO", "line", lineAdv)
	return datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           lineAdv.ID,
			OrderID:      lineAdv.OrderID,
			ReleaseCount: lineAdv.ReleaseCount,
			Cost:         lineAdv.Cost,
			Text:         lineAdv.Text,
			Tags:         fetchTagsName(lineAdv.Tags),
			ExtraCharges: fetchExtraChargesName(lineAdv.ExtraCharges),
			ReleaseDates: fetchReleaseDates(lineAdv.ReleaseDates),
		},
	}
}

func convertLineAdvertisementToModel(lineAdv *datatransferobjects.LineAdvertisementDTO) AdvertisementLine {
	logging.Logger.Debug("orm: Converting LineAdvertisement DTO to databaseObject")
	return AdvertisementLine{
		Advertisement: Advertisement{
			ID:           lineAdv.ID,
			OrderID:      lineAdv.OrderID,
			ReleaseCount: lineAdv.ReleaseCount,
			Cost:         lineAdv.Cost,
			Text:         lineAdv.Text,
		},
		Tags:         tagNamesToModel(lineAdv.Tags),
		ExtraCharges: extraChargesToModel(lineAdv.ExtraCharges),
		ReleaseDates: releaseDatesToModel(lineAdv.ReleaseDates),
	}
}

func convertTagToDTO(tag *Tag) datatransferobjects.TagDTO {
	logging.Logger.Debug("orm: Converting Tag databaseObject to DTO")
	return datatransferobjects.TagDTO{
		TagName: tag.Name,
		TagCost: tag.Cost,
	}
}

func convertTagToModel(tag *datatransferobjects.TagDTO) Tag {
	logging.Logger.Debug("orm: Converting Tag DTO to databaseObject")
	return Tag{
		Name: tag.TagName,
		Cost: tag.TagCost,
	}
}

func convertExtraChargeToDTO(extraCharge *ExtraCharge) datatransferobjects.ExtraChargeDTO {
	logging.Logger.Debug("orm: Converting ExtraCharge databaseObject to DTO")
	return datatransferobjects.ExtraChargeDTO{
		ChargeName: extraCharge.Name,
		Multiplier: extraCharge.Multiplier,
	}
}

func convertExtraChargeToModel(extraCharge *datatransferobjects.ExtraChargeDTO) ExtraCharge {
	logging.Logger.Debug("orm: Converting ExtraCharge DTO to databaseObject")
	return ExtraCharge{
		Name:       extraCharge.ChargeName,
		Multiplier: extraCharge.Multiplier,
	}
}

func convertReleaseDateToModel(releaseDate time.Time) ReleaseDates {
	logging.Logger.Debug("orm: Converting ReleaseDate to databaseObject")
	return ReleaseDates{
		ReleaseDate: releaseDate,
	}
}

func tagNamesToModel(tag []string) []Tag {
	logging.Logger.Debug("orm: Fetching Tag name")
	var tagModel []Tag
	for _, v := range tag {
		tagModel = append(tagModel, Tag{Name: v})
	}
	return tagModel
}

func extraChargesToModel(extraCharge []string) []ExtraCharge {
	var extraChargeModel []ExtraCharge
	for _, v := range extraCharge {
		extraChargeModel = append(extraChargeModel, ExtraCharge{Name: v})
	}
	return extraChargeModel
}

func releaseDatesToModel(releaseDates []time.Time) []ReleaseDates {
	var releaseDatesModel []ReleaseDates
	for _, v := range releaseDates {
		releaseDatesModel = append(releaseDatesModel, ReleaseDates{
			ReleaseDate: time.Date(v.Year(), v.Month(), v.Day(), 0, 0, 0, 0, time.UTC)})
	}
	return releaseDatesModel
}

func convertCostRateToModel(costRate *datatransferobjects.CostRateDTO) CostRate {
	logging.Logger.Debug("orm: Converting ExtraCharge DTO to databaseObject")
	return CostRate{
		Name:            costRate.Name,
		OneWordOrSymbol: costRate.ForOneWordSymbol,
		OneSquare:       costRate.ForOneSquare,
		CalcForOneWord:  costRate.CalcForOneWord,
	}
}

func convertCostRateToDto(costRate *CostRate) datatransferobjects.CostRateDTO {
	logging.Logger.Debug("orm: Converting ExtraCharge DTO to databaseObject")
	return datatransferobjects.CostRateDTO{
		Name:             costRate.Name,
		ForOneWordSymbol: costRate.OneWordOrSymbol,
		ForOneSquare:     costRate.OneSquare,
		CalcForOneWord:   costRate.CalcForOneWord,
	}
}
