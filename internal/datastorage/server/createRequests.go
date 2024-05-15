package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (ds *ServerStorage) NewClient(ctx context.Context, client *datatransferobjects.ClientDTO) (string, error) {
	logging.Logger.Debug("server: create request. Saving Client to server")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertClientToModel(client), false)
	if err != nil {
		return "", err
	}
	_, err = ds.s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: clients, Name: "", Queries: nil}, &b)
	return "", err
}

func (ds *ServerStorage) NewAdvertisementsOrder(ctx context.Context, order *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("server: create request. Saving Order to server")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertOrderToModel(order), false)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	q := make(map[string]string)
	q["nested"] = "1"
	_, err = ds.s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: orders, Name: "", Queries: q}, &b)
	return datatransferobjects.OrderDTO{}, err
}

func (ds *ServerStorage) NewLineAdvertisement(ctx context.Context, lineadv *datatransferobjects.LineAdvertisementDTO) (int, error) {
	logging.Logger.Debug("server: create request. Saving LineAdvertisement to server")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertLineAdvertisementToModel(lineadv), false)
	if err != nil {
		return 0, err
	}
	_, err = ds.s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: "", Queries: nil}, &b)
	return 0, err
}

func (ds *ServerStorage) NewBlockAdvertisement(ctx context.Context, blockadv *datatransferobjects.BlockAdvertisementDTO) (int, error) {
	logging.Logger.Debug("server: create request. Saving BlockAdvertisement to server")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertBlockAdvertisementToModel(blockadv), false)
	if err != nil {
		return 0, err
	}
	_, err = ds.s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: blockadvertisements, Name: "", Queries: nil}, &b)
	return 0, err
}

func (ds *ServerStorage) NewTag(ctx context.Context, tag *datatransferobjects.TagDTO) (string, error) {
	logging.Logger.Debug("server: create request. Saving Tag to server")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertTagToModel(tag), false)
	if err != nil {
		return "", err
	}
	_, err = ds.s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: tags, Name: "", Queries: nil}, &b)
	return "", err
}

func (ds *ServerStorage) NewExtraCharge(ctx context.Context, ExtraCharges *datatransferobjects.ExtraChargeDTO) (string, error) {
	logging.Logger.Debug("server: create request. Saving ExtraCharges to server")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertExtraChargeToModel(ExtraCharges), false)
	if err != nil {
		return "", err
	}
	_, err = ds.s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: extraCharges, Name: "", Queries: nil}, &b)
	return "", err
}

func (ds *ServerStorage) NewCostRate(ctx context.Context, costRate *datatransferobjects.CostRateDTO) (string, error) {
	logging.Logger.Debug("server: create request. Saving CostRate to server")
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ds.convertCostRateToModel(costRate), false)
	if err != nil {
		return "", err
	}
	_, err = ds.s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: costRates, Name: "", Queries: nil}, &b)
	return "", err
}

func (ds *ServerStorage) NewFile(ctx context.Context, file *datatransferobjects.FileDTO) (string, error) {
	logging.Logger.Debug("server: create request. Saving CostRate to server")
	var b bytes.Buffer

	err := encodedecoder.ToJSON(&b, ds.convertFileToModel(file), false)
	if err != nil {
		return "", err
	}
	data, err := ds.s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: files, Name: "", Queries: nil}, &b)
	if err != nil {
		return "", err
	}
	var files []FileFront
	err = json.Unmarshal(data, &files)
	if err != nil {
		return "", err
	}
	return ds.fileFrontToString(files), err
}

func (ds *ServerStorage) NewFileUpload(ctx context.Context, file datatransferobjects.FileStream) ([]string, error) {
	var b bytes.Buffer
	multipartWriter := multipart.NewWriter(&b)

	f1, err := multipartWriter.CreateFormFile("files", file.Name)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(f1, file.Data)
	if err != nil {
		return nil, err
	}
	defer file.Data.Close()
	multipartWriter.Close()

	data, err := ds.s.UploadMultipart(ctx, &datatransferobjects.RequestDTO{Kind: files, Name: "", Queries: nil}, &b, multipartWriter.FormDataContentType())
	if err != nil {
		return nil, err
	}
	var files []FileFront
	err = json.Unmarshal(data, &files)
	if err != nil {
		return nil, err
	}
	return ds.filesFrontToStringArray(files), err
}
