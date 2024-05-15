package requestscontroller

import (
	"context"
	"regexp"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/handlers"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type b64URL interface {
	ToBase64URLString(in []byte) string
}

type RequestsHandler struct {
	costAdvRegex *regexp.Regexp
	costRegex    *regexp.Regexp
	dateRegex    *regexp.Regexp

	converter       DataHandler
	fileStorage     handlers.FileStorage
	fileManager     handlers.ServerRequests
	configAcessor   handlers.ConfigsAccesor
	costRateManager handlers.CostRateCalculator
	req             handlers.HandlerRequests
	dbRequests      handlers.DBManager
	jsonExporter    handlers.JSONExporter
	jsonImporter    handlers.JSONImporter
	reportGenerator handlers.Reports
	collator        collate.Collator
	rec             handlers.Reciever
	b64             b64URL
}

func NewRequestsHandler(configSaver handlers.ConfigsAccesor, converter DataHandler, b64 b64URL) *RequestsHandler {
	return &RequestsHandler{
		costAdvRegex:  regexp.MustCompile(`^([0-9]+)[\.|\,]?([0-9]{1,2})?$`),
		costRegex:     regexp.MustCompile(`^([+-])?([0-9]+)[\.|\,]?([0-9]{1,2})?$`),
		dateRegex:     regexp.MustCompile(`^([0-9]{2})\.([0-9]{2})\.([0-9]{4,})$`),
		configAcessor: configSaver,
		converter:     converter,
		collator:      *collate.New(language.Russian, collate.IgnoreCase),
		b64:           b64,
	}
}

func (r *RequestsHandler) SetReciever(rec handlers.Reciever) {
	r.rec = rec
}

func (r *RequestsHandler) SetRequestsGateway(req handlers.HandlerRequests) {
	r.req = req
}

func (r *RequestsHandler) SetFileManager(filesReq handlers.ServerRequests) {
	r.fileManager = filesReq
}

func (r *RequestsHandler) SetDatabaseGateway(req handlers.DBManager) {
	r.dbRequests = req
}

func (r *RequestsHandler) SetFileStorage(storage handlers.FileStorage) {
	r.fileStorage = storage
}

func (r *RequestsHandler) SetSecondaryCompopnents(jsExp handlers.JSONExporter, jsImp handlers.JSONImporter,
	repGen handlers.Reports, costCalc handlers.CostRateCalculator) {
	r.jsonExporter = jsExp
	r.jsonImporter = jsImp
	r.reportGenerator = repGen
	r.costRateManager = costCalc
}

func (r *RequestsHandler) AllTags() {
	r.req.AllTags()
}

func (r *RequestsHandler) CloseDatabaseConnection() error {
	if r.dbRequests.DatabaseConnected() {
		return r.req.Close()
	}
	return nil
}

func (r *RequestsHandler) AllExtraCharges() {
	r.req.AllExtraCharges()
}

func (r *RequestsHandler) AllCostRates() {
	r.req.AllCostRates()
}

func (r *RequestsHandler) AllClients() {
	r.req.AllClients()
}

func (r *RequestsHandler) AllOrders() {
	r.req.AllOrders()
}

func (r *RequestsHandler) AllBlockAdvertisements() {
	r.req.AllBlockAdvertisements()
}

func (r *RequestsHandler) BlockAdvertisementsActual() {
	r.req.BlockAdvertisementsActualReleaseDate()
}

func (r *RequestsHandler) AllLineAdvertisements() {
	r.req.AllLineAdvertisements()
}

func (r *RequestsHandler) LineAdvertisementsActual() {
	go r.req.LineAdvertisementsActualReleaseDate()
}

func (r *RequestsHandler) RemoveTag(tag string) {
	go r.req.RemoveTagByName(tag)
}

func (r *RequestsHandler) UpdateTag(tag *presenter.TagDTO) error {
	tagdto, err := r.converter.TagToDto(tag)
	if err != nil {
		return err
	}

	go r.req.UpdateTag(tagdto)
	return nil
}

func (r *RequestsHandler) CreateTag(tag *presenter.TagDTO) error {
	tagdto, err := r.converter.TagToDto(tag)
	if err != nil {
		return err
	}
	go r.req.NewTag(tagdto)
	return nil
}

func (r *RequestsHandler) RemoveExtraCharge(charge string) {
	go r.req.RemoveExtraChargeByName(charge)
}

func (r *RequestsHandler) UpdateExtraCharge(charge *presenter.ExtraChargeDTO) error {
	extraChargedto, err := r.converter.ExtraChargeToDto(charge)
	if err != nil {
		return err
	}
	go r.req.UpdateExtraCharge(extraChargedto)
	return nil
}

func (r *RequestsHandler) CreateExtraCharge(charge *presenter.ExtraChargeDTO) error {
	extraChargedto, err := r.converter.ExtraChargeToDto(charge)
	if err != nil {
		return err
	}
	go r.req.NewExtraCharge(extraChargedto)
	return nil
}

func (r *RequestsHandler) CreateClient(client *presenter.ClientDTO) error {
	clientdto, err := r.converter.ClientToDto(client)
	if err != nil {
		return err
	}
	go r.req.NewClient(clientdto)
	return nil
}

func (r *RequestsHandler) RemoveClient(client string) {
	go r.req.RemoveClientByName(client)
}

func (r *RequestsHandler) UpdateClient(client *presenter.ClientDTO) error {
	clientdto, err := r.converter.ClientToDto(client)
	if err != nil {
		return err
	}
	go r.req.UpdateClient(clientdto)
	return nil
}

func (r *RequestsHandler) RemoveOrder(order *presenter.OrderDTO) {
	orderdto, err := r.converter.OrderToDto(order)
	if err != nil {
		return
	}
	go r.req.RemoveOrderByID(orderdto.ID)
}

func (r *RequestsHandler) CreateOrder(order *presenter.OrderDTO, blkAdv []presenter.BlockAdvertisementDTO, lineAdv []presenter.LineAdvertisementDTO) error {
	orderdto, err := r.converter.OrderToDto(order)
	if err != nil {
		return err
	}
	for i := range blkAdv {
		block, err := r.converter.BlockAdvertisementToDto(&blkAdv[i])
		if err != nil {
			return err
		}
		orderdto.BlockAdvertisements = append(orderdto.BlockAdvertisements, block)
	}
	for i := range lineAdv {
		line, err := r.converter.LineAdvertisementToDto(&lineAdv[i])
		if err != nil {
			return err
		}
		orderdto.LineAdvertisements = append(orderdto.LineAdvertisements, line)
	}

	go r.req.NewAdvertisementsOrder(orderdto)
	return nil
}

func (r *RequestsHandler) CalculateOrderCost(order *presenter.OrderDTO, blkAdv []presenter.BlockAdvertisementDTO, lineAdv []presenter.LineAdvertisementDTO) error {
	orderdto, err := r.converter.OrderToDto(order)
	if err != nil {
		return err
	}
	if blkAdv != nil {
		orderdto.BlockAdvertisements = make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blkAdv))
		for i := range blkAdv {
			block, err := r.converter.BlockAdvertisementToDto(&blkAdv[i])
			if err != nil {
				return err
			}
			orderdto.BlockAdvertisements = append(orderdto.BlockAdvertisements, block)
		}
	}
	if lineAdv != nil {
		orderdto.LineAdvertisements = make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdv))
		for i := range lineAdv {
			line, err := r.converter.LineAdvertisementToDto(&lineAdv[i])
			if err != nil {
				return err
			}
			orderdto.LineAdvertisements = append(orderdto.LineAdvertisements, line)
		}
	}
	go r.costRateManager.CalculateOrderCost(orderdto)
	return nil
}

func (r *RequestsHandler) UpdateOrder(order *presenter.OrderDTO) error {
	orderdto, err := r.converter.OrderToDto(order)
	if err != nil {
		return err
	}
	go r.req.UpdateOrder(orderdto)
	return nil
}

func (r *RequestsHandler) ArrayToString(arr []string) string {
	return r.converter.ArrayToString(arr)
}

func (r *RequestsHandler) YearMonthDayToString(year, month, day uint) string {
	return r.converter.YearMonthDayToString(year, month, day)
}

func (r *RequestsHandler) CreateLineAdvertisement(line *presenter.LineAdvertisementDTO) error {
	linedto, err := r.converter.LineAdvertisementToDto(line)
	if err != nil {
		logging.Logger.Error("presenter.CreateLineAdvertisement: an error occuren while converting", "error", err)
		return err
	}
	go r.req.NewLineAdvertisement(linedto)
	return nil
}

func (r *RequestsHandler) CalculateLineAdvertisementCost(line *presenter.LineAdvertisementDTO) error {
	linedto, err := r.converter.LineAdvertisementToDto(line)
	if err != nil {
		logging.Logger.Error("presenter.LineAdvertisementCost: an error occuren while converting", "error", err)
		return err
	}

	go r.costRateManager.CalculateLineAdvertisementCost(linedto)
	return nil
}

func (r *RequestsHandler) RemoveLineAdvertisement(line *presenter.LineAdvertisementDTO) {
	linedto, err := r.converter.LineAdvertisementToDto(line)
	if err != nil {
		return
	}
	go r.req.RemoveLineAdvertisementByID(linedto.ID)
}

func (r *RequestsHandler) UpdateLineAdvertisement(line *presenter.LineAdvertisementDTO) error {
	linedto, err := r.converter.LineAdvertisementToDto(line)
	if err != nil {
		return err
	}
	go r.req.UpdateLineAdvertisement(linedto)
	return nil
}

func (r *RequestsHandler) CreateBlockAdvertisement(block *presenter.BlockAdvertisementDTO) error {
	blockdto, err := r.converter.BlockAdvertisementToDto(block)
	if err != nil {
		logging.Logger.Error("presenter.CreateBlockAdvertisement: an error occuren while converting", "error", err)
		return err
	}
	go r.req.NewBlockAdvertisement(blockdto)
	return nil
}

func (r *RequestsHandler) CalculateBlockAdvertisementCost(block *presenter.BlockAdvertisementDTO) error {
	blockdto, err := r.converter.BlockAdvertisementToDto(block)
	if err != nil {
		logging.Logger.Error("presenter.BlockAdvertisementCost: an error occuren while converting", "error", err)
		return err
	}
	go r.costRateManager.CalculateBlockAdvertisementCost(blockdto)
	return nil
}

func (r *RequestsHandler) RemoveBlockAdvertisement(block *presenter.BlockAdvertisementDTO) {
	blockdto, err := r.converter.BlockAdvertisementToDto(block)
	if err != nil {
		return
	}
	go r.req.RemoveBlockAdvertisementByID(blockdto.ID)
}

func (r *RequestsHandler) UpdateBlockAdvertisement(block *presenter.BlockAdvertisementDTO) error {
	blockdto, err := r.converter.BlockAdvertisementToDto(block)
	if err != nil {
		return err
	}
	go r.req.UpdateBlockAdvertisement(blockdto)
	return nil
}

func (r *RequestsHandler) StartExportingJson(ctx context.Context, path string) {
	go r.jsonExporter.ExportJsonToFile(ctx, path)
}

func (r *RequestsHandler) StartImportingJson(ctx context.Context, path string, params *presenter.ImportParams) {
	importParams := datatransferobjects.ImportParams{
		AllBlocks:       params.AllBlocks,
		AlllLines:       params.AlllLines,
		ActualClients:   params.ActualClients,
		AllTags:         params.AllTags,
		AllExtraCharges: params.AllExtraCharges,
		AllCostRates:    params.AllCostRates,
		IgnoreErrors:    params.IgnoreErrors,
		ThickMode:       params.ThickMode,
	}
	r.IgnoreMessages(true)
	go r.jsonImporter.ImportJson(ctx, path, importParams)
}

func (r *RequestsHandler) CreateCostRate(costRate *presenter.CostRateDTO) error {
	costRateDro, err := r.converter.CostRateToDTO(costRate)
	if err != nil {
		return err
	}
	go r.req.NewCostRate(costRateDro)
	return nil
}

func (r *RequestsHandler) UpdateCostRate(costRate *presenter.CostRateDTO) error {
	costRateDro, err := r.converter.CostRateToDTO(costRate)
	if err != nil {
		return err
	}
	go r.req.UpdateCostRate(costRateDro)
	return nil
}

func (r *RequestsHandler) RemoveCostRate(costRate string) {
	go r.req.RemoveCostRateByName(costRate)
}

func (r *RequestsHandler) SetActiveCostRate(costRateName string) error {
	go r.costRateManager.SetActiveCostRate(costRateName)
	return nil
}

func (r *RequestsHandler) ActiveCostRate() {
	go r.costRateManager.ActiveCostRate()
}

func (r *RequestsHandler) CheckCostRate() {
	go r.costRateManager.SelectedCostRate()
}

func (r *RequestsHandler) CreateAfvertisementReport(ctx context.Context, repParam *presenter.ReportParams) error {
	rep, err := r.converter.AdvReportToDTO(repParam)
	if err != nil {
		return err
	}

	go r.reportGenerator.GenerateReport(ctx, rep)
	return nil
}

func (r *RequestsHandler) ParsePath(path string) string {
	return r.converter.ParsePath(path)
}

func (r *RequestsHandler) ConnectionInfo() presenter.ConnectionInfo {
	connectionInfo := r.req.ConnectionInfo()
	return presenter.ConnectionInfo{
		Addres: connectionInfo["adress"],
		Path:   connectionInfo["apiPath"],
		Token:  connectionInfo["token"],
	}
}

func (r *RequestsHandler) LockReciever(lock bool) {
	if r.rec != nil {
		r.rec.LockRecieving(lock)
	}
}

func (r *RequestsHandler) IgnoreMessages(lock bool) {
	if r.rec != nil {
		r.rec.IgnoreMessages(lock)
	}
}
