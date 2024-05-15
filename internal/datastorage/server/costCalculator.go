package server

import (
	"bytes"
	"context"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
)

func (ds *ServerStorage) CalculateOrderCost(ctx context.Context, order datatransferobjects.OrderDTO, costRateName string) (datatransferobjects.OrderDTO, error) {
	ord := ds.convertOrderToModel(&order)
	q := make(map[string]string)
	q["calculatecost"] = "1"
	q["costrate"] = costRateName
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, ord, false)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}

	resp, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: orders, Queries: q}, &b)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	r := bytes.NewReader(resp)
	var gotOrder OrderFront
	err = encodedecoder.FromJSON(&gotOrder, r)
	if err != nil {
		return datatransferobjects.OrderDTO{}, err
	}
	return ds.convertOrderToDTO(&gotOrder), nil
}

func (ds *ServerStorage) CalculateBlockAdvertisementCost(ctx context.Context, block datatransferobjects.BlockAdvertisementDTO,
	costRateName string) (datatransferobjects.BlockAdvertisementDTO, error) {
	blk := ds.convertBlockAdvertisementToModel(&block)
	q := make(map[string]string)
	q["calculatecost"] = "1"
	q["costrate"] = costRateName
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, blk, false)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}

	resp, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: blockadvertisements, Queries: q}, &b)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	r := bytes.NewReader(resp)
	var gotBlock BlockAdvertisementFront
	err = encodedecoder.FromJSON(&gotBlock, r)
	if err != nil {
		return datatransferobjects.BlockAdvertisementDTO{}, err
	}
	return ds.convertBlockAdvertisementToDTO(&gotBlock), nil
}

func (ds *ServerStorage) CalculateLineAdvertisementCost(ctx context.Context, line datatransferobjects.LineAdvertisementDTO,
	costRateName string) (datatransferobjects.LineAdvertisementDTO, error) {
	blk := ds.convertLineAdvertisementToModel(&line)
	q := make(map[string]string)
	q["calculatecost"] = "1"
	q["costrate"] = costRateName
	var b bytes.Buffer
	err := encodedecoder.ToJSON(&b, blk, false)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}

	resp, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Queries: q}, &b)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	r := bytes.NewReader(resp)
	var gotLine LineAdvertisementFront
	err = encodedecoder.FromJSON(&gotLine, r)
	if err != nil {
		return datatransferobjects.LineAdvertisementDTO{}, err
	}
	return ds.convertLineAdvertisementToDTO(&gotLine), nil
}
