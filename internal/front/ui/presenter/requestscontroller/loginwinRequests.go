package requestscontroller

import (
	"encoding/json"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (r *RequestsHandler) ConnectToLocalDatabase(viewParam *presenter.LocalDSN) error {
	logging.Logger.Debug("requestsController: Got Connect to local database request")
	dto, err := r.converter.LocalDsnToDto(viewParam)
	if err != nil {
		return err
	}
	encoded, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	go r.dbRequests.ConnectToDatabase(viewParam.Type, encoded)
	return nil
}

func (r *RequestsHandler) ConnectToNetworkDatabase(viewParam *presenter.NetworkDataBaseDSN) error {
	logging.Logger.Debug("requestsController: Got Connect to network database request")

	dto, err := r.converter.NetworkDsnToDto(viewParam)
	if err != nil {
		return err
	}
	encoded, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	go r.dbRequests.ConnectToDatabase(viewParam.DatabaseName, encoded)
	return nil
}

func (r *RequestsHandler) ConnectToServer(viewParam *presenter.ServerDSN) error {
	logging.Logger.Debug("requestsController: Got Connect to network database request")

	dto, err := r.converter.ServerDsnToDto(viewParam)
	if err != nil {
		return err
	}
	encoded, err := json.Marshal(dto)
	if err != nil {
		return err
	}
	go r.dbRequests.ConnectToServer(viewParam.DatabaseName, encoded)
	return nil
}

func (r *RequestsHandler) Databases() {
	logging.Logger.Debug("requestsController: Got database list request")
	r.dbRequests.AvailableLocalDatabases()
	r.dbRequests.AvailableNetworkDatabases()
}

func (r *RequestsHandler) DefaultNetworkPort(dbName string) {
	logging.Logger.Debug("requestsController: Got default port for database database request")
	r.dbRequests.DefaultPort(dbName)
}

func (r *RequestsHandler) SaveConfig(name string, view any) error {
	raw, err := r.encodeData(&view)
	if err != nil {
		return err
	}
	err = r.configAcessor.SaveConfig(name, raw)
	if err != nil {
		return err
	}

	return nil
}

func (r *RequestsHandler) LoadConfig(name string, view any) error {
	raw, err := r.configAcessor.Load(name)
	if err != nil {
		return err
	}
	err = r.decodeData(raw, &view)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestsHandler) RemoveConfig(name string) error {
	err := r.configAcessor.Remove(name)
	if err != nil {
		return err
	}
	return nil
}

func (r *RequestsHandler) Ping() error {
	return nil
}

func (r *RequestsHandler) CreateFile(name string) error {
	f, err := r.fileStorage.New(name)
	if err != nil {
		return err
	}
	return f.Close()
}
