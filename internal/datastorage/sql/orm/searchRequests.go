package orm

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"gorm.io/gorm/clause"
)

func (ds *DataStorageOrm) AllClients(ctx context.Context) ([]datatransferobjects.ClientDTO, error) {
	logging.Logger.Debug("orm: search request. Getting All Clients")
	var clients []Client
	err := ds.db.WithContext(ctx).Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}}).Find(&clients).Error
	if err != nil {
		return nil, handleError(err)
	}
	clientsDto := make([]datatransferobjects.ClientDTO, 0, len(clients))
	for i := range clients {
		clientsDto = append(clientsDto, convertClientToDTO(&clients[i]))
	}
	return clientsDto, nil
}

func (ds *DataStorageOrm) AllOrders(ctx context.Context) ([]datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("orm: search request. Getting All Orders")
	var orders []Order
	err := ds.db.WithContext(ctx).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&orders).Error
	if err != nil {
		return nil, handleError(err)
	}
	orderDTO := make([]datatransferobjects.OrderDTO, 0, len(orders))
	for i := range orders {
		orderDTO = append(orderDTO, convertOrderToDTO(&orders[i]))
	}
	return orderDTO, nil
}

func (ds *DataStorageOrm) OrdersByClientName(ctx context.Context, name string) ([]datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("orm: search request. Searching Orders by Client name")
	var orders []Order
	err := ds.db.WithContext(ctx).Where("client_name = ?", name).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&orders).Error
	if err != nil {
		return nil, handleError(err)
	}

	orderDTO := make([]datatransferobjects.OrderDTO, 0, len(orders))
	for i := range orders {
		orderDTO = append(orderDTO, convertOrderToDTO(&orders[i]))
	}
	return orderDTO, err
}

func (ds *DataStorageOrm) AllLineAdvertisements(ctx context.Context) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Getting all LineAdvertisements")
	var lineAdv []AdvertisementLine
	err := ds.db.WithContext(ctx).Preload(clause.Associations).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&lineAdv).Error
	if err != nil {
		return nil, handleError(err)
	}
	lineAdvDto := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdv))
	for i := range lineAdv {
		lineAdvDto = append(lineAdvDto, convertLineAdvertisementToDTO(&lineAdv[i]))
	}
	return lineAdvDto, nil
}

func (ds *DataStorageOrm) LineAdvertisementsByOrderID(ctx context.Context, id int) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching LineAdvertisements by Order id")
	var (
		lineAdvDto []datatransferobjects.LineAdvertisementDTO
		lineAdv    []AdvertisementLine
	)

	err := ds.db.WithContext(ctx).Preload(clause.Associations).Where("order_id = ?", id).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).
		Find(&lineAdv).Error
	if err != nil {
		return nil, handleError(err)
	}

	lineAdvDto = make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdv))
	for i := range lineAdv {
		lineAdvDto = append(lineAdvDto, convertLineAdvertisementToDTO(&lineAdv[i]))
	}
	return lineAdvDto, nil
}

func (ds *DataStorageOrm) AllBlockAdvertisements(ctx context.Context) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Getting all BlockAdvertisements")
	var blockAdv []AdvertisementBlock
	err := ds.db.WithContext(ctx).Preload(clause.Associations).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&blockAdv).Error
	if err != nil {
		return nil, handleError(err)
	}

	blockAdvDto := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdv))
	for i := range blockAdv {
		blockAdvDto = append(blockAdvDto, convertBlockAdvertisementToDTO(&blockAdv[i]))
	}
	return blockAdvDto, nil
}

func (ds *DataStorageOrm) BlockAdvertisementsByOrderID(ctx context.Context, id int) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching BlockAdvertisements by Order id")
	var (
		blockAdvDto []datatransferobjects.BlockAdvertisementDTO
		blockAdv    []AdvertisementBlock
	)
	err := ds.db.WithContext(ctx).Preload(clause.Associations).Where("order_id = ?", id).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&blockAdv).
		Error
	if err != nil {
		return nil, handleError(err)
	}

	blockAdvDto = make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdv))
	for i := range blockAdv {
		blockAdvDto = append(blockAdvDto, convertBlockAdvertisementToDTO(&blockAdv[i]))
	}
	return blockAdvDto, nil
}

func (ds *DataStorageOrm) BlockAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching BlockAdvertisements Between releaseDates", "from", from, "to", to)
	var (
		blockAdvDto []datatransferobjects.BlockAdvertisementDTO
		blockAdv    []AdvertisementBlock
	)
	err := ds.db.WithContext(ctx).Where("id IN (?)", ds.db.Table("advertisementsblock_releasedates").
		Select("advertisement_block_id").Where("release_dates_release_date BETWEEN ? AND ?", from, to)).
		Preload(clause.Associations).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&blockAdv).Error
	if err != nil {
		return nil, handleError(err)
	}

	blockAdvDto = make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdv))
	for i := range blockAdv {
		blockAdvDto = append(blockAdvDto, convertBlockAdvertisementToDTO(&blockAdv[i]))
	}
	return blockAdvDto, nil
}

func (ds *DataStorageOrm) BlockAdvertisementFromReleaseDates(ctx context.Context, from time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching BlockAdvertisements Between releaseDates", "from", from)
	var (
		blockAdvDto []datatransferobjects.BlockAdvertisementDTO
		blockAdv    []AdvertisementBlock
	)
	err := ds.db.WithContext(ctx).Where("id IN (?)", ds.db.Table("advertisementsblock_releasedates").
		Select("advertisement_block_id").Where("release_dates_release_date > ?", from)).
		Preload(clause.Associations).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&blockAdv).Error
	if err != nil {
		return nil, handleError(err)
	}

	blockAdvDto = make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdv))
	for i := range blockAdv {
		blockAdvDto = append(blockAdvDto, convertBlockAdvertisementToDTO(&blockAdv[i]))
	}
	return blockAdvDto, nil
}

func (ds *DataStorageOrm) LineAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching LineAdvertisements Between releaseDates", "from", from, "to", to)
	var (
		lineAdvDto []datatransferobjects.LineAdvertisementDTO
		lineAdv    []AdvertisementLine
	)
	err := ds.db.WithContext(ctx).Where("id IN (?)", ds.db.Table("advertisementsline_releasedates").
		Select("advertisement_line_id").Where("release_dates_release_date BETWEEN ? AND ?", from, to)).
		Preload(clause.Associations).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&lineAdv).Error
	if err != nil {
		return nil, handleError(err)
	}

	lineAdvDto = make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdv))
	for _, v := range lineAdv {
		lineAdvDto = append(lineAdvDto, convertLineAdvertisementToDTO(&v))
	}
	return lineAdvDto, nil
}

func (ds *DataStorageOrm) LineAdvertisementFromReleaseDates(ctx context.Context, from time.Time) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching LineAdvertisements Between releaseDates", "from", from)

	var (
		lineAdvDto []datatransferobjects.LineAdvertisementDTO
		lineAdv    []AdvertisementLine
	)
	err := ds.db.WithContext(ctx).Where("id IN (?)", ds.db.Table("advertisementsline_releasedates").
		Select("advertisement_line_id").Where("release_dates_release_date > ?", from)).
		Preload(clause.Associations).Order(clause.OrderByColumn{Column: clause.Column{Name: "id"}}).Find(&lineAdv).Error
	if err != nil {
		return nil, handleError(err)
	}

	lineAdvDto = make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdv))
	for i := range lineAdv {
		lineAdvDto = append(lineAdvDto, convertLineAdvertisementToDTO(&lineAdv[i]))
	}
	return lineAdvDto, nil
}

func (ds *DataStorageOrm) AllTags(ctx context.Context) ([]datatransferobjects.TagDTO, error) {
	logging.Logger.Debug("orm: search request. Getting all Tags")
	var (
		tags   []Tag
		tagDTO []datatransferobjects.TagDTO
	)

	err := ds.db.WithContext(ctx).Find(&tags).Error
	if err != nil {
		return nil, handleError(err)
	}
	tagDTO = make([]datatransferobjects.TagDTO, 0, len(tags))
	for i := range tags {
		tagDTO = append(tagDTO, convertTagToDTO(&tags[i]))
	}

	return tagDTO, err
}

func (ds *DataStorageOrm) AllExtraCharges(ctx context.Context) ([]datatransferobjects.ExtraChargeDTO, error) {
	logging.Logger.Debug("orm: search request. Getting all ExtraCharges")

	var (
		extraCharges   []ExtraCharge
		extraChargeDTO []datatransferobjects.ExtraChargeDTO
	)

	err := ds.db.WithContext(ctx).Find(&extraCharges).Error
	if err != nil {
		return nil, handleError(err)
	}

	extraChargeDTO = make([]datatransferobjects.ExtraChargeDTO, 0, len(extraCharges))
	for i := range extraCharges {
		extraChargeDTO = append(extraChargeDTO, convertExtraChargeToDTO(&extraCharges[i]))
	}

	return extraChargeDTO, err
}

func (ds *DataStorageOrm) AllCostRates(ctx context.Context) ([]datatransferobjects.CostRateDTO, error) {
	logging.Logger.Debug("orm: search request. Getting All CostRates")
	var costRates []CostRate
	err := ds.db.WithContext(ctx).Find(&costRates).Error
	if err != nil {
		return nil, handleError(err)
	}

	costRatesDto := make([]datatransferobjects.CostRateDTO, 0, len(costRates))
	for i := range costRates {
		costRatesDto = append(costRatesDto, convertCostRateToDto(&costRates[i]))
	}
	return costRatesDto, nil
}
