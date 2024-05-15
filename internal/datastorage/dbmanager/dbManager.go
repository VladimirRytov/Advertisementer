package dbmanager

import (
	"errors"
	"slices"

	"github.com/VladimirRytov/advertisementer/internal/handlers"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

const (
	_ = iota
	LocalDB
	NetworkDB
	Server
)

type AvailableDataBases struct {
	dbType             int
	name               string
	defaultNetworkPort uint
	connection         handlers.DataBase
}

type AvailableServer struct {
	defaultNetworkPort uint
	connection         handlers.Server
}

type DataBaseManager struct {
	availableDatabases map[string]AvailableDataBases
	server             AvailableServer
}

func NewDatabaseManager() *DataBaseManager {
	logging.Logger.Info("db manager: Initializing database manager")
	dbm := &DataBaseManager{
		availableDatabases: make(map[string]AvailableDataBases),
	}
	return dbm
}

func (dbm *DataBaseManager) Register(name string, port uint, dbType int, connector handlers.DataBase) {
	dbm.availableDatabases[name] = AvailableDataBases{
		dbType:             dbType,
		name:               name,
		defaultNetworkPort: port,
		connection:         connector,
	}
}

func (dbm *DataBaseManager) RegisterServerGateway(port uint, dbType int, connector handlers.Server) {
	dbm.server = AvailableServer{
		defaultNetworkPort: port,
		connection:         connector,
	}
}

func (dbm *DataBaseManager) AvailableLocalDatabases() []string {
	dbList := make([]string, 0)
	for k := range dbm.availableDatabases {
		if dbm.availableDatabases[k].dbType == LocalDB {
			dbList = append(dbList, dbm.availableDatabases[k].name)
		}
	}
	slices.Sort(dbList)
	return dbList
}

func (dbm *DataBaseManager) AvailableNetworkDatabases() []string {
	dbList := make([]string, 0)
	for k := range dbm.availableDatabases {
		if dbm.availableDatabases[k].dbType == NetworkDB {
			dbList = append(dbList, dbm.availableDatabases[k].name)
		}
	}
	slices.Sort(dbList)
	return dbList
}

func (dbm *DataBaseManager) DefaultPort(dnname string) uint {
	return dbm.availableDatabases[dnname].defaultNetworkPort
}

func (dbm *DataBaseManager) ConnectToDatabase(dbName string, params []byte) (handlers.DataBase, error) {
	logging.Logger.Info("db manager: Search database for connecting")
	var err error
	for i := range dbm.availableDatabases {
		if dbm.availableDatabases[i].name == dbName {
			err = dbm.availableDatabases[i].connection.ConnectToDatabase(params)
			if err != nil {
				return nil, err
			}
			logging.Logger.Info("db manager: Connection to the database was successful")
			return dbm.availableDatabases[i].connection, nil
		}
	}
	logging.Logger.Error("db manager: Database not found", "dbName", dbName)
	return nil, errors.New("error: this database not supported")
}

func (dbm *DataBaseManager) ConnectToServer(params []byte) (handlers.Server, error) {
	logging.Logger.Info("db manager: Search database for connecting")
	err := dbm.server.connection.ConnectToDatabase(params)
	if err != nil {
		return nil, err
	}
	logging.Logger.Info("db manager: Connection to the database was successful")
	return dbm.server.connection, nil
}
