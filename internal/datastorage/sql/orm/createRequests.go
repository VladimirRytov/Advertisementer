package orm

import (
	"context"
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (ds *DataStorageOrm) NewClient(ctx context.Context, client *datatransferobjects.ClientDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving Client to database")
	newClient := convertClientToModel(client)
	result := ds.db.WithContext(ctx).Table("clients").Create(&newClient)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return newClient.Name, errors.New("контрагент уже существует")
	}
	return newClient.Name, result.Error
}

func (ds *DataStorageOrm) NewAdvertisementsOrder(ctx context.Context, order *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("orm: create request. Saving Order to database")
	orderModel := convertOrderToModel(order)
	logging.Logger.Debug("orm: create request. Saving Order to database", "orderModel", orderModel)

	result := ds.db.WithContext(ctx).Table("orders").Create(&orderModel)
	if result.Error == gorm.ErrDuplicatedKey {
		result.Error = errors.New("заказ уже существует")
	}
	newOrder := convertOrderToDTO(&orderModel)
	return newOrder, result.Error
}

func (ds *DataStorageOrm) NewLineAdvertisement(ctx context.Context, lineadv *datatransferobjects.LineAdvertisementDTO) (int, error) {
	logging.Logger.Debug("orm: create request. Saving LineAdvertisement to database")
	newLineAdv := convertLineAdvertisementToModel(lineadv)
	result := ds.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Table("advertisement_lines").Create(&newLineAdv)
	if result.Error == gorm.ErrDuplicatedKey {
		result.Error = errors.New("строковое объявление уже существует")
	}

	return newLineAdv.ID, result.Error
}

func (ds *DataStorageOrm) NewBlockAdvertisement(ctx context.Context, blockadv *datatransferobjects.BlockAdvertisementDTO) (int, error) {
	logging.Logger.Debug("orm: create request. Saving BlockAdvertisement to database")
	newBlockAdv := convertBlockAdvertisementToModel(blockadv)
	result := ds.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&newBlockAdv)
	if result.Error == gorm.ErrDuplicatedKey {
		result.Error = errors.New("блочное объявление уже существует")
	}

	return newBlockAdv.ID, result.Error
}

func (ds *DataStorageOrm) NewTag(ctx context.Context, tag *datatransferobjects.TagDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving Tag to database")
	newTag := convertTagToModel(tag)
	result := ds.db.WithContext(ctx).Create(&newTag)
	if result.Error == gorm.ErrDuplicatedKey {
		result.Error = errors.New("тэг уже существует")
	}
	return newTag.Name, result.Error
}

func (ds *DataStorageOrm) NewExtraCharge(ctx context.Context, ExtraCharges *datatransferobjects.ExtraChargeDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving ExtraCharges to database")
	newExtraCharges := convertExtraChargeToModel(ExtraCharges)
	result := ds.db.WithContext(ctx).Create(&newExtraCharges)
	if result.Error == gorm.ErrDuplicatedKey {
		result.Error = errors.New("наценка уже существует")
	}
	return newExtraCharges.Name, result.Error
}

func (ds *DataStorageOrm) NewCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) (string, error) {
	logging.Logger.Debug("orm: create request. Saving CostRate to database")
	newCostRate := convertCostRateToModel(costRate)
	result := ds.db.WithContext(ctx).Create(&newCostRate)
	if result.Error == gorm.ErrDuplicatedKey {
		result.Error = errors.New("тариф уже существует")
	}
	return newCostRate.Name, result.Error
}
