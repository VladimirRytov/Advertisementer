package thinclienthandler

import (
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type responcerDummy struct{}

var (
	Clients             []datatransferobjects.ClientDTO
	Orders              []datatransferobjects.OrderDTO
	LineAdvertisements  []datatransferobjects.LineAdvertisementDTO
	BlockAdvertisements []datatransferobjects.BlockAdvertisementDTO
	Tags                []datatransferobjects.TagDTO
	ExtraCharges        []datatransferobjects.ExtraChargeDTO

	ClientsError             error
	OrdersError              error
	LineAdvertisementsError  error
	BlockAdvertisementsError error
	TagsError                error
	ExtraChargesError        error
	DatabaseError            error
)

func (r *responcerDummy) SendFileFirstPlace(file *datatransferobjects.FileDTO) {}

func (r *responcerDummy) UnlockFilesWindow() {}

func (r *responcerDummy) RemoveFileByName(name string) {}

func (r *responcerDummy) SendFile(file *datatransferobjects.FileDTO) {}

func (r *responcerDummy) ShowFile(file *datatransferobjects.FileDTO) {}

func (r *responcerDummy) Stop() {}

func (r *responcerDummy) SetConnectionStatus(b bool) {
}
func (r *responcerDummy) SetMode(b int8) {
}
func (r *responcerDummy) Start(msg []string) {}

func (r *responcerDummy) SendBlockAdvertisementCost(b datatransferobjects.BlockAdvertisementDTO) {}

func (r *responcerDummy) SendActiveCostRate(names string) {}

func (r *responcerDummy) SendLineAdvertisementCost(l datatransferobjects.LineAdvertisementDTO) {}

func (r *responcerDummy) SendOrderCost(o datatransferobjects.OrderDTO) {}

func (r *responcerDummy) RemoveCostRate(names string) {}

func (r *responcerDummy) SendCostRate(c *datatransferobjects.CostRateDTO) {}

func (r *responcerDummy) CancelProgressWithError(err error) {}

func (r *responcerDummy) ProgressComplete() {}

func (r *responcerDummy) SendError(err error) {}

func (r *responcerDummy) SendMessage(msg string) {}

func (r *responcerDummy) RequestComplete() {}

func (r *responcerDummy) SendLocalDatabases(cli []string) {}

func (r *responcerDummy) SendNetworkDatabases(cli []string) {}

func (r *responcerDummy) SendDefaultPort(cli uint) {}

func (r *responcerDummy) SendClient(cli *datatransferobjects.ClientDTO) {
	Clients = make([]datatransferobjects.ClientDTO, 0)
	Clients = append(Clients, *cli)
}

func (r *responcerDummy) SendAdvertisementsOrder(order *datatransferobjects.OrderDTO) {
	Orders = make([]datatransferobjects.OrderDTO, 0)
	Orders = append(Orders, *order)
}

func (r *responcerDummy) SendLineAdvertisement(line *datatransferobjects.LineAdvertisementDTO) {
	LineAdvertisements = make([]datatransferobjects.LineAdvertisementDTO, 0)
	LineAdvertisements = append(LineAdvertisements, *line)
}

func (r *responcerDummy) SendBlockAdvertisement(block *datatransferobjects.BlockAdvertisementDTO) {
	BlockAdvertisements = make([]datatransferobjects.BlockAdvertisementDTO, 0)
	BlockAdvertisements = append(BlockAdvertisements, *block)
}

func (r *responcerDummy) SendExtraCharge(charge *datatransferobjects.ExtraChargeDTO) {
	ExtraCharges = make([]datatransferobjects.ExtraChargeDTO, 0)
	ExtraCharges = append(ExtraCharges, *charge)

}

func (r *responcerDummy) SendTag(tag *datatransferobjects.TagDTO) {
	Tags = make([]datatransferobjects.TagDTO, 0)
	Tags = append(Tags, *tag)

}

func (r *responcerDummy) RemoveClientByName(names string) {}

func (r *responcerDummy) RemoveOrderByID(ids int) {}

func (r *responcerDummy) RemoveLineAdvertisementByID(ids int) {}

func (r *responcerDummy) RemoveBlockAdvertisementByID(ids int) {}

func (r *responcerDummy) RemoveTagByName(names string) {}

func (r *responcerDummy) RemoveExtraChargeByName(names string) {}

func (r *responcerDummy) SendClientsError(err error) {
	ClientsError = err
}

func (r *responcerDummy) SendAdvertisementsOrdersError(err error) {
	OrdersError = err
}

func (r *responcerDummy) SendLineAdvertisementsError(err error) {
	LineAdvertisementsError = err
}

func (r *responcerDummy) SendBlockAdvertisementsError(err error) {
	BlockAdvertisementsError = err
}

func (r *responcerDummy) SendExtraChargesError(err error) {
	ExtraChargesError = err
}

func (r *responcerDummy) SendTagsError(err error) {
	TagsError = err
}

func (r *responcerDummy) SendConnectionError(err error) {
	TagsError = err
}

func (r *responcerDummy) SendDatabaseError(err error) {
	TagsError = err
}

func (r *responcerDummy) SendSuccesConnection() {}

func (r *responcerDummy) InitializationComplete() {}
