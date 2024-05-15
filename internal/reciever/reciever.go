package reciever

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net/url"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

const (
	ClientType             = "Client"
	OrderType              = "Order"
	BlockAdvertisementType = "BlockAdv"
	LineAdvertisementType  = "LineAdv"
	TagType                = "Tag"
	ExtraChargeType        = "ExtraCharge"
	CostRateType           = "CostRate"
	Pong                   = "Check"

	DeleteAction = 1
	CreateAction = 2
)

type SendWrapper struct {
	Type   string `json:"type"`
	Action int    `json:"action"`
	Entry  []byte `json:"entry"`
}

var ErrPingTimeout = errPingTimeout()

func errPingTimeout() error {
	return errors.New("send ping signal but got nothing")
}

type RecieverHandler struct {
	recievers map[string]Reciever
	resp      Responcer
}

func NewRecieveHandler(app Responcer) *RecieverHandler {
	return &RecieverHandler{resp: app, recievers: make(map[string]Reciever)}
}

func (rh *RecieverHandler) RegisterReciever(name string, rec Reciever) {
	rh.recievers[name] = rec
}

func (rh *RecieverHandler) Subscribe(token, addr, path string) error {
	url, err := url.JoinPath(addr, path)
	if err != nil {
		return err
	}
	err = rh.recievers["webSock"].Subscribe(token, url)
	if err != nil {
		return err
	}
	return nil
}

func (rh *RecieverHandler) ConnectionLost() {
	rh.resp.SetConnectionStatus(false)
}

func (rh *RecieverHandler) ConnectionRestore() {
	rh.resp.SetConnectionStatus(true)
}

func (rh *RecieverHandler) HandleClient(data *bytes.Buffer, action int) error {
	logging.Logger.Debug("reciever.handleClient: handling Client")
	switch action {
	case CreateAction:
		var cli datatransferobjects.ClientDTO
		err := encodedecoder.FromJSON(&cli, data)
		if err != nil {
			logging.Logger.Error("reciever.handleClient: decode error", "error", err)
			return err
		}
		rh.resp.SendClient(&cli)
	case DeleteAction:
		rh.resp.RemoveClientByName(data.String())
	default:
		return errors.New("unexpected action")
	}
	return nil
}

func (rh *RecieverHandler) HandleOrder(data *bytes.Buffer, action int) error {
	logging.Logger.Debug("reciever.handleOrder: handling Order")
	switch action {
	case CreateAction:
		var order datatransferobjects.OrderDTO
		err := encodedecoder.FromJSON(&order, data)
		if err != nil {
			logging.Logger.Error("reciever.handleOrder: decode error", "error", err)
			return err
		}
		rh.resp.SendAdvertisementsOrder(&order)
	case DeleteAction:
		if data.Len() > 1 {
			return errors.New("wrong remove param")
		}
		id, err := binary.ReadVarint(data)
		if err != nil {
			return err
		}
		rh.resp.RemoveOrderByID(int(id))
	}
	return nil
}

func (rh *RecieverHandler) HandleBlockAdvertisement(data *bytes.Buffer, action int) error {
	logging.Logger.Debug("reciever.handleBlockAdvertisement: handling BlockAdvertisement")
	switch action {
	case CreateAction:
		var blockAdv datatransferobjects.BlockAdvertisementDTO
		err := encodedecoder.FromJSON(&blockAdv, data)
		if err != nil {
			logging.Logger.Error("reciever.handleBlockAdvertisement: decode error", "error", err)
			return err
		}
		rh.resp.SendBlockAdvertisement(&blockAdv)
	case DeleteAction:
		if data.Len() > 1 {
			return errors.New("wrong remove param")
		}
		id, err := binary.ReadVarint(data)
		if err != nil {
			return err
		}
		rh.resp.RemoveBlockAdvertisementByID(int(id))
	}
	return nil
}

func (rh *RecieverHandler) HandleLineAdvertisement(data *bytes.Buffer, action int) error {
	logging.Logger.Debug("reciever.handleLineAdvertisement: handling LineAdvertisement")
	switch action {
	case CreateAction:
		var lineAdv datatransferobjects.LineAdvertisementDTO
		err := encodedecoder.FromJSON(&lineAdv, data)
		if err != nil {
			logging.Logger.Error("reciever.handleLineAdvertisement: decode error", "error", err)
			return err
		}
		rh.resp.SendLineAdvertisement(&lineAdv)
	case DeleteAction:
		if data.Len() > 1 {
			return errors.New("wrong remove param")
		}
		id, err := binary.ReadVarint(data)
		if err != nil {
			return err
		}
		rh.resp.RemoveLineAdvertisementByID(int(id))
	}
	return nil
}

func (rh *RecieverHandler) HandleTag(data *bytes.Buffer, action int) error {
	logging.Logger.Debug("reciever.handleTag: handling Tag")
	switch action {
	case CreateAction:
		var tag datatransferobjects.TagDTO
		err := encodedecoder.FromJSON(&tag, data)
		if err != nil {
			logging.Logger.Error("reciever.handleTag: decode error", "error", err)
			return err
		}
		rh.resp.SendTag(&tag)
	case DeleteAction:
		rh.resp.RemoveTagByName(data.String())
	default:
		return errors.New("unexpected action")
	}
	return nil
}

func (rh *RecieverHandler) HandleExtraCharge(data *bytes.Buffer, action int) error {
	logging.Logger.Debug("reciever.handleExtraCharge: handling ExtraCharge")
	switch action {
	case CreateAction:
		var charge datatransferobjects.ExtraChargeDTO
		err := encodedecoder.FromJSON(&charge, data)
		if err != nil {
			logging.Logger.Error("reciever.handleExtraCharge: decode error", "error", err)
			return err
		}
		rh.resp.SendExtraCharge(&charge)
	case DeleteAction:
		rh.resp.RemoveExtraChargeByName(data.String())
	default:
		return errors.New("unexpected action")
	}
	return nil
}

func (rh *RecieverHandler) HandleCostRate(data *bytes.Buffer, action int) error {
	logging.Logger.Debug("reciever.handleCostRate: handling CostRate")
	switch action {
	case CreateAction:
		var costRate datatransferobjects.CostRateDTO
		err := encodedecoder.FromJSON(&costRate, data)
		if err != nil {
			logging.Logger.Error("reciever.handleCostRate: decode error", "error", err)
			return err
		}
		rh.resp.SendCostRate(&costRate)
	case DeleteAction:
		rh.resp.RemoveCostRate(data.String())
	default:
		return errors.New("unexpected action")
	}
	return nil
}
func (rh *RecieverHandler) LockRecieving(lock bool) {
	for k := range rh.recievers {
		rh.recievers[k].LockRecieve(lock)
	}
}

func (rh *RecieverHandler) IgnoreMessages(lock bool) {
	for k := range rh.recievers {
		rh.recievers[k].IgnoreMessages(lock)
	}
}
