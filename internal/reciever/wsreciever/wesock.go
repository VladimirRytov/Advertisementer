package wsreciever

import (
	"bytes"
	"errors"
	"net/url"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"nhooyr.io/websocket"
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

type WsReciever struct {
	reciever RecieveHandler
	b64      Encodedecoder
	tok      string
	addr     string
	conn     *websocket.Conn
	lockRec  bool
	ignore   bool
	cache    chan SendWrapper
}

func New(app RecieveHandler, b64Dec Encodedecoder) *WsReciever {
	return &WsReciever{
		reciever: app,
		b64:      b64Dec,
		cache:    make(chan SendWrapper, 100),
	}
}

func (ws *WsReciever) parseAdress(adress string) (string, error) {
	destConn, err := url.Parse(adress)
	if err != nil {
		logging.Logger.Error("wsReciever.parseAdress: cannot parse url", "error", err)
		return "", err
	}
	addr := "ws://" + destConn.Host + destConn.Path
	return url.JoinPath(addr, "subscribers", "ws", "listener")
}

func (ws *WsReciever) Subscribe(token, adress string) error {
	newAddr, err := ws.parseAdress(adress)
	if err != nil {
		logging.Logger.Error("wsReciever.Subscribe: cannot parse adress", "error", err)
		return err
	}
	err = ws.dialToServer(token, newAddr)
	if err != nil {
		logging.Logger.Error("wsReciever: cannot subscribe", "error", err)
		return err
	}
	ws.tok = token
	ws.addr = adress
	go ws.listen()
	ws.reciever.ConnectionRestore()
	return nil
}

func (ws *WsReciever) CloseConn() error {
	return ws.conn.Close(websocket.StatusNormalClosure, "")
}

func (ws *WsReciever) reconnect() {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	for range t.C {
		if err := ws.Subscribe(ws.tok, ws.addr); err == nil {
			ws.reciever.ConnectionRestore()
			return
		}
	}
}

func (rh *WsReciever) LockRecieve(lock bool) {
	if !lock {
		close(rh.cache)
		for msg := range rh.cache {
			rh.handleNewMessage(msg)
		}
	}
	rh.lockRec = lock
	rh.cache = make(chan SendWrapper, 100)
}

func (rh *WsReciever) IgnoreMessages(lock bool) {
	rh.ignore = lock
}

func (rh *WsReciever) handleNewMessage(got SendWrapper) error {
	var err error
	if rh.lockRec {
		if len(rh.cache) == 99 {
			logging.Logger.Error("wsReciever.handleNewMessage: cache overflow. Closing connection")
			rh.CloseConn()
			return errors.New("cache overflow. Closing connection")
		}
		rh.cache <- got
		return nil
	}

	decEntry, _ := rh.b64.FromBase64(got.Entry)
	b := bytes.NewBuffer(decEntry)
	switch got.Type {
	case ClientType:
		err = rh.reciever.HandleClient(b, got.Action)
	case OrderType:
		err = rh.reciever.HandleOrder(b, got.Action)
	case BlockAdvertisementType:
		err = rh.reciever.HandleBlockAdvertisement(b, got.Action)
	case LineAdvertisementType:
		err = rh.reciever.HandleLineAdvertisement(b, got.Action)
	case TagType:
		err = rh.reciever.HandleTag(b, got.Action)
	case ExtraChargeType:
		err = rh.reciever.HandleExtraCharge(b, got.Action)
	case CostRateType:
		err = rh.reciever.HandleCostRate(b, got.Action)
	default:
		logging.Logger.Warn("reciever.handlemessage: got unexpected message")
		return errors.New("unexpected data")
	}
	return err
}
