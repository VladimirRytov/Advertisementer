package wsreciever

import (
	"bytes"
	"context"
	"net/http"

	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"nhooyr.io/websocket"
)

func (ws *WsReciever) dialToServer(token, destAddr string) error {
	var err error
	header := http.Header{
		"Authorization": []string{"Bearer " + token},
	}
	ws.conn, _, err = websocket.Dial(context.TODO(), destAddr, &websocket.DialOptions{
		HTTPHeader: header,
	})
	if err != nil {
		return err
	}
	ws.conn.SetReadLimit(-1)
	return nil
}

func (ws *WsReciever) listen() error {
	for {
		var sw SendWrapper
		_, data, err := ws.conn.Read(context.Background())
		if err != nil {
			logging.Logger.Error("wsReciever.listen: got error while reading", "error", err)
			ws.reciever.ConnectionLost()
			ws.reconnect()
			return nil
		}
		if ws.ignore {
			continue
		}
		r := bytes.NewReader(data)
		err = encodedecoder.FromJSON(&sw, r)
		if err != nil {
			logging.Logger.Error("wsReciever.listen: cannot unmarshal data", "error", err)
			continue
		}
		ws.handleNewMessage(sw)
	}
}
