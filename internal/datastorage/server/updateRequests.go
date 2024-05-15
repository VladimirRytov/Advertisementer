package server

import (
	"bytes"
	"context"
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (ds *ServerStorage) UpdateClient(ctx context.Context, client *datatransferobjects.ClientDTO) error {
	logging.Logger.Debug("orm: update request. Updating Client")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertClientToModel(client), false)
	if err != nil {
		return err
	}
	_, err = ds.s.UpdateRequest(ctx, &datatransferobjects.RequestDTO{Kind: clients, Name: client.Name}, &b)
	return err
}

func (ds *ServerStorage) UpdateOrder(ctx context.Context, order *datatransferobjects.OrderDTO) error {
	logging.Logger.Debug("orm: update request. Updating Order")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertOrderToModel(order), false)
	if err != nil {
		return err
	}
	_, err = ds.s.UpdateRequest(ctx, &datatransferobjects.RequestDTO{Kind: orders, Name: strconv.Itoa(order.ID)}, &b)
	return err
}

func (ds *ServerStorage) UpdateLineAdvertisement(ctx context.Context, lineadv *datatransferobjects.LineAdvertisementDTO) error {
	logging.Logger.Debug("orm: update request. Updating LineAdvertisement")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertLineAdvertisementToModel(lineadv), false)
	if err != nil {
		return err
	}
	_, err = ds.s.UpdateRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: strconv.Itoa(lineadv.ID)}, &b)
	return err
}

func (ds *ServerStorage) UpdateBlockAdvertisement(ctx context.Context, blockadv *datatransferobjects.BlockAdvertisementDTO) error {
	logging.Logger.Debug("orm: update request. Updating BlockAdvertisement")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertBlockAdvertisementToModel(blockadv), false)
	if err != nil {
		return err
	}
	_, err = ds.s.UpdateRequest(ctx, &datatransferobjects.RequestDTO{Kind: blockadvertisements, Name: strconv.Itoa(blockadv.ID)}, &b)
	return err
}

func (ds *ServerStorage) UpdateExtraCharge(ctx context.Context, extraCharge *datatransferobjects.ExtraChargeDTO) error {
	logging.Logger.Debug("orm: update request. Updating ExtraCharge")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertExtraChargeToModel(extraCharge), false)
	if err != nil {
		return err
	}
	_, err = ds.s.UpdateRequest(ctx, &datatransferobjects.RequestDTO{Kind: extraCharges, Name: extraCharge.ChargeName}, &b)
	return err
}

func (ds *ServerStorage) UpdateTag(ctx context.Context, tag *datatransferobjects.TagDTO) error {
	logging.Logger.Debug("orm: update request. Updating Tag")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertTagToModel(tag), false)
	if err != nil {
		return err
	}
	_, err = ds.s.UpdateRequest(ctx, &datatransferobjects.RequestDTO{Kind: tags, Name: tag.TagName}, &b)
	return err
}

func (ds *ServerStorage) UpdateCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) error {
	logging.Logger.Debug("orm: update request. Updating CostRate")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertCostRateToModel(costRate), false)
	if err != nil {
		return err
	}
	_, err = ds.s.UpdateRequest(ctx, &datatransferobjects.RequestDTO{Kind: costRates, Name: costRate.Name}, &b)
	return err
}
