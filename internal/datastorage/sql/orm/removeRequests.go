package orm

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (ds *DataStorageOrm) RemoveClientByName(ctx context.Context, name string) error {
	logging.Logger.Debug("orm: remove request. Removing Client by name")
	err := ds.db.WithContext(ctx).Delete(&Client{}, "name = ?", name).Error
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStorageOrm) RemoveOrderByID(ctx context.Context, id int) error {
	logging.Logger.Debug("orm: remove request. Removing Order by id")
	err := ds.db.WithContext(ctx).Delete(&Order{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStorageOrm) RemoveLineAdvertisementByID(ctx context.Context, id int) error {
	logging.Logger.Debug("orm: remove request. Removing LineAdvertisement by id")
	err := ds.db.WithContext(ctx).Model(&AdvertisementLine{Advertisement: Advertisement{ID: id}}).Association("Tags").Unscoped().Clear()
	if err != nil {
		return err
	}
	err = ds.db.WithContext(ctx).Model(&AdvertisementLine{Advertisement: Advertisement{ID: id}}).Association("ExtraCharges").Unscoped().Clear()
	if err != nil {
		return err
	}
	err = ds.db.WithContext(ctx).Preload("ReleaseDates").Model(&AdvertisementLine{Advertisement: Advertisement{ID: id}}).Association("ReleaseDates").Unscoped().Clear()
	if err != nil {
		return err
	}
	err = ds.db.WithContext(ctx).Delete(&AdvertisementLine{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStorageOrm) RemoveBlockAdvertisementByID(ctx context.Context, id int) error {
	logging.Logger.Debug("orm: remove request. Removing BlockAdvertisement by id")
	err := ds.db.WithContext(ctx).Model(&AdvertisementBlock{Advertisement: Advertisement{ID: id}}).Association("Tags").Unscoped().Clear()
	if err != nil {
		return err
	}
	err = ds.db.WithContext(ctx).Model(&AdvertisementBlock{Advertisement: Advertisement{ID: id}}).Association("ExtraCharges").Unscoped().Clear()
	if err != nil {
		return err
	}
	err = ds.db.WithContext(ctx).Model(&AdvertisementBlock{Advertisement: Advertisement{ID: id}}).Association("ReleaseDates").Unscoped().Clear()
	if err != nil {
		return err
	}
	err = ds.db.WithContext(ctx).Delete(&AdvertisementBlock{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (ds *DataStorageOrm) RemoveTagByName(ctx context.Context, name string) error {
	logging.Logger.Debug("orm: remove request. Removing Tag by Name")
	err := ds.db.WithContext(ctx).Delete(&Tag{}, "name = ?", name).Error
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStorageOrm) RemoveExtraChargeByName(ctx context.Context, name string) error {
	logging.Logger.Debug("orm: remove request. Removing ExtraCharge by Name")
	err := ds.db.WithContext(ctx).Delete(&ExtraCharge{}, "name = ?", name).Error
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStorageOrm) RemoveCostRateByName(ctx context.Context, name string) error {
	logging.Logger.Debug("orm: remove request. Removing CostRate by Name")
	err := ds.db.WithContext(ctx).Delete(&CostRate{}, "name = ?", name).Error
	if err != nil {
		return err
	}
	return nil
}
