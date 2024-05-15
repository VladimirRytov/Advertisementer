package orm

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"gorm.io/gorm/clause"
)

func (ds *DataStorageOrm) ClientByName(ctx context.Context, name string) (datatransferobjects.ClientDTO, error) {
	logging.Logger.Debug("orm: get request. Getting Client by name")
	var client Client
	result := ds.db.WithContext(ctx).First(&client, "name = ?", name)
	if result.Error != nil {
		return datatransferobjects.ClientDTO{}, result.Error
	}

	newClient := convertClientToDTO(&client)
	return newClient, nil
}

func (ds *DataStorageOrm) OrderByID(ctx context.Context, id int) (datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("orm: get request. Getting Order by id")
	var order Order
	result := ds.db.WithContext(ctx).First(&order, id)
	if result.Error != nil {
		return datatransferobjects.OrderDTO{}, result.Error
	}

	newOrder := convertOrderToDTO(&order)
	return newOrder, nil
}

func (ds *DataStorageOrm) LineAdvertisementByID(ctx context.Context, id int) (datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: get request. Getting LineAdvertisement by id")
	var lineAdv AdvertisementLine
	result := ds.db.WithContext(ctx).Preload(clause.Associations).First(&lineAdv, id)
	if result.Error != nil {
		return datatransferobjects.LineAdvertisementDTO{}, result.Error
	}

	lineAdvDto := convertLineAdvertisementToDTO(&lineAdv)
	return lineAdvDto, nil
}

func (ds *DataStorageOrm) BlockAdvertisementByID(ctx context.Context, id int) (datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: get request. Getting BlockAdvertisement by id")
	var blockAdv AdvertisementBlock
	result := ds.db.WithContext(ctx).Preload(clause.Associations).First(&blockAdv, id)
	if result.Error != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, result.Error
	}

	blockAdvDTO := convertBlockAdvertisementToDTO(&blockAdv)
	return blockAdvDTO, nil
}

func (ds *DataStorageOrm) TagByName(ctx context.Context, tagName string) (datatransferobjects.TagDTO, error) {
	logging.Logger.Debug("orm: get request. Getting Tag by name")
	var tag Tag
	result := ds.db.WithContext(ctx).Where("name = ?", tagName).First(&tag)
	if result.Error != nil {
		return datatransferobjects.TagDTO{}, result.Error
	}

	tagDTO := convertTagToDTO(&tag)
	return tagDTO, result.Error
}

func (ds *DataStorageOrm) ExtraChargeByName(ctx context.Context, chargeName string) (datatransferobjects.ExtraChargeDTO, error) {
	logging.Logger.Debug("orm: get request. Getting ExtraCharge by name")
	var extraCharge ExtraCharge
	result := ds.db.WithContext(ctx).Where("name = ?", chargeName).First(&extraCharge)
	if result.Error != nil {
		return datatransferobjects.ExtraChargeDTO{}, result.Error
	}

	extraChargeDTO := convertExtraChargeToDTO(&extraCharge)
	return extraChargeDTO, nil
}

func (ds *DataStorageOrm) CostRateByName(ctx context.Context, name string) (datatransferobjects.CostRateDTO, error) {
	logging.Logger.Debug("orm: get request. Getting CostRate by name")
	var costRate CostRate
	result := ds.db.WithContext(ctx).Where("name = ?", name).First(&costRate)
	if result.Error != nil {
		return datatransferobjects.CostRateDTO{}, result.Error
	}

	costRateDTO := convertCostRateToDto(&costRate)
	return costRateDTO, nil
}
