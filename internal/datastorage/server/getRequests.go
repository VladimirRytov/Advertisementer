package server

import (
	"bytes"
	"context"
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (ds *ServerStorage) ClientByName(ctx context.Context, name string) (datatransferobjects.ClientDTO, error) {
	logging.Logger.Debug("server: get request. Getting Client by name")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: clients, Name: name}, nil)
	if err != nil {
		return datatransferobjects.ClientDTO{}, err
	}
	var client ClientFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&client, r)
	if err != nil {
		return datatransferobjects.ClientDTO{}, err
	}
	return ds.convertClientToDTO(&client), err
}

func (ds *ServerStorage) OrderByID(ctx context.Context, id int) (datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("server: get request. Getting Order by id")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: orders, Name: strconv.Itoa(id)}, nil)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	var order OrderFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&order, r)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	return ds.convertOrderToDTO(&order), err
}

func (ds *ServerStorage) LineAdvertisementByID(ctx context.Context, id int) (datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("server: get request. Getting LineAdvertisement by id")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: strconv.Itoa(id)}, nil)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	var line LineAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&line, r)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	return ds.convertLineAdvertisementToDTO(&line), err
}

func (ds *ServerStorage) BlockAdvertisementByID(ctx context.Context, id int) (datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("server: get request. Getting BlockAdvertisement by id")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: strconv.Itoa(id)}, nil)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	var block BlockAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&block, r)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	return ds.convertBlockAdvertisementToDTO(&block), err
}

func (ds *ServerStorage) TagByName(ctx context.Context, tagName string) (datatransferobjects.TagDTO, error) {
	logging.Logger.Debug("server: get request. Getting Tag by name")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: tags, Name: tagName}, nil)
	if err != nil {
		return datatransferobjects.TagDTO{}, err
	}
	var tag TagFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&tag, r)
	if err != nil {
		return datatransferobjects.TagDTO{}, err
	}
	return ds.convertTagToDTO(&tag), err
}

func (ds *ServerStorage) ExtraChargeByName(ctx context.Context, chargeName string) (datatransferobjects.ExtraChargeDTO, error) {
	logging.Logger.Debug("server: get request. Getting ExtraCharge by name")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: extraCharges, Name: chargeName}, nil)
	if err != nil {
		return datatransferobjects.ExtraChargeDTO{}, err
	}
	var charge ExtraChargeFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&charge, r)
	if err != nil {
		return datatransferobjects.ExtraChargeDTO{}, err
	}
	return ds.convertExtraChargeToDTO(&charge), err
}

func (ds *ServerStorage) CostRateByName(ctx context.Context, name string) (datatransferobjects.CostRateDTO, error) {
	logging.Logger.Debug("server: get request. Getting CostRate by name")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: costRates, Name: name}, nil)
	if err != nil {
		return datatransferobjects.CostRateDTO{}, err
	}
	var costRate CostRateFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&costRate, r)
	if err != nil {
		return datatransferobjects.CostRateDTO{}, err
	}
	return ds.convertCostRateToDto(&costRate), err
}

func (ds *ServerStorage) FileByName(ctx context.Context, fileName string) (datatransferobjects.FileDTO, error) {
	logging.Logger.Debug("server: get request. Getting File by name")
	q := make(map[string]string)
	q["format"] = "json"
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: files, Name: fileName, Queries: q}, nil)
	if err != nil {
		return datatransferobjects.FileDTO{}, nil
	}

	var file FileFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&file, r)
	if err != nil {
		return datatransferobjects.FileDTO{}, err
	}
	return ds.convertFileToDto(&file)
}

func (ds *ServerStorage) FileMiniatureByName(ctx context.Context, fileName string, size string) (datatransferobjects.FileDTO, error) {
	logging.Logger.Debug("server: get request. Getting File by name")
	q := make(map[string]string)
	q["format"] = "json"
	q["size"] = size
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: files, Name: fileName, Queries: q}, nil)
	if err != nil {
		return datatransferobjects.FileDTO{}, nil
	}

	var file FileFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&file, r)
	if err != nil {
		return datatransferobjects.FileDTO{}, err
	}
	return ds.convertFileToDto(&file)
}
