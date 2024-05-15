package server

import (
	"context"
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (ds *ServerStorage) RemoveClientByName(ctx context.Context, name string) error {
	logging.Logger.Debug("server: remove request. Removing Client by name")
	_, err := ds.s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: clients, Name: name})
	return err
}

func (ds *ServerStorage) RemoveOrderByID(ctx context.Context, id int) error {
	logging.Logger.Debug("server: remove request. Removing Order by id")
	_, err := ds.s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: orders, Name: strconv.Itoa(id)})
	return err
}

func (ds *ServerStorage) RemoveLineAdvertisementByID(ctx context.Context, id int) error {
	logging.Logger.Debug("server: remove request. Removing LineAdvertisement by id")
	_, err := ds.s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: strconv.Itoa(id)})
	return err
}

func (ds *ServerStorage) RemoveBlockAdvertisementByID(ctx context.Context, id int) error {
	logging.Logger.Debug("server: remove request. Removing BlockAdvertisement by id")
	_, err := ds.s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: blockadvertisements, Name: strconv.Itoa(id)})
	return err
}
func (ds *ServerStorage) RemoveTagByName(ctx context.Context, name string) error {
	logging.Logger.Debug("server: remove request. Removing Tag by Name")
	_, err := ds.s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: tags, Name: name})
	return err
}

func (ds *ServerStorage) RemoveExtraChargeByName(ctx context.Context, name string) error {
	logging.Logger.Debug("server: remove request. Removing ExtraCharge by Name")
	_, err := ds.s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: extraCharges, Name: name})
	return err
}

func (ds *ServerStorage) RemoveCostRateByName(ctx context.Context, name string) error {
	logging.Logger.Debug("server: remove request. Removing CostRate by Name")
	_, err := ds.s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: costRates, Name: name})
	return err
}

func (ds *ServerStorage) RemoveFileByName(ctx context.Context, name string) error {
	logging.Logger.Debug("server: remove request. Removing File by Name")
	_, err := ds.s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: files, Name: name})
	return err
}
