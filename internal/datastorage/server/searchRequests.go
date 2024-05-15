package server

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (ds *ServerStorage) AllClients(ctx context.Context) ([]datatransferobjects.ClientDTO, error) {
	logging.Logger.Debug("orm: search request. Getting All Clients")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: clients, Name: ""}, nil)
	if err != nil {
		return nil, err
	}
	var clientF []ClientFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&clientF, r)
	if err != nil {
		return nil, err
	}
	client := make([]datatransferobjects.ClientDTO, 0, len(clientF))
	for i := range clientF {
		client = append(client, ds.convertClientToDTO(&clientF[i]))
	}
	return client, err
}

func (ds *ServerStorage) AllOrders(ctx context.Context) ([]datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("orm: search request. Getting All Orders")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: orders, Name: ""}, nil)
	if err != nil {
		return nil, err
	}
	var orderF []OrderFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&orderF, r)
	if err != nil {
		return nil, err
	}
	order := make([]datatransferobjects.OrderDTO, 0, len(orderF))
	for i := range orderF {
		order = append(order, ds.convertOrderToDTO(&orderF[i]))
	}
	return order, err
}

func (ds *ServerStorage) OrdersByClientName(ctx context.Context, name string) ([]datatransferobjects.OrderDTO, error) {
	logging.Logger.Debug("orm: search request. Searching Orders by Client name")
	q := make(map[string]string)
	q["client"] = name
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: orders, Name: "", Queries: q}, nil)
	if err != nil {
		return nil, err
	}
	var orderF []OrderFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&orderF, r)
	if err != nil {
		return nil, err
	}
	order := make([]datatransferobjects.OrderDTO, 0, len(orderF))
	for i := range orderF {
		order = append(order, ds.convertOrderToDTO(&orderF[i]))
	}
	return order, err
}

func (ds *ServerStorage) AllLineAdvertisements(ctx context.Context) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Getting all LineAdvertisements")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: ""}, nil)
	if err != nil {
		return nil, err
	}
	var lineAdvF []LineAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&lineAdvF, r)
	if err != nil {
		return nil, err
	}
	lineadv := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvF))
	for i := range lineAdvF {
		lineadv = append(lineadv, ds.convertLineAdvertisementToDTO(&lineAdvF[i]))
	}
	return lineadv, err
}

func (ds *ServerStorage) LineAdvertisementsByOrderID(ctx context.Context, id int) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching LineAdvertisements by Order id")
	q := make(map[string]string)
	q["orderid"] = strconv.Itoa(id)
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: "", Queries: q}, nil)
	if err != nil {
		return nil, err
	}
	var lineAdvF []LineAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&lineAdvF, r)
	if err != nil {
		return nil, err
	}
	lineadv := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvF))
	for i := range lineAdvF {
		lineadv = append(lineadv, ds.convertLineAdvertisementToDTO(&lineAdvF[i]))
	}
	return lineadv, err
}
func (ds *ServerStorage) AllBlockAdvertisements(ctx context.Context) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Getting all BlockAdvertisements")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: blockadvertisements, Name: ""}, nil)
	if err != nil {
		return nil, err
	}
	var blockAdvF []BlockAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&blockAdvF, r)
	if err != nil {
		return nil, err
	}
	blockAdv := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvF))
	for i := range blockAdvF {
		blockAdv = append(blockAdv, ds.convertBlockAdvertisementToDTO(&blockAdvF[i]))
	}
	return blockAdv, err
}
func (ds *ServerStorage) BlockAdvertisementsByOrderID(ctx context.Context, id int) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching BlockAdvertisements by Order id")
	q := make(map[string]string)
	q["orderid"] = strconv.Itoa(id)
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: blockadvertisements, Name: "", Queries: q}, nil)
	if err != nil {
		return nil, err
	}
	var blockAdvF []BlockAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&blockAdvF, r)
	if err != nil {
		return nil, err
	}
	blockAdv := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvF))
	for i := range blockAdvF {
		blockAdv = append(blockAdv, ds.convertBlockAdvertisementToDTO(&blockAdvF[i]))
	}
	return blockAdv, err
}

func (ds *ServerStorage) BlockAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching BlockAdvertisements Between releaseDates", "from", from, "to", to)
	q := make(map[string]string)
	q["between"] = from.Format("02.01.2006") + "-" + to.Format("02.01.2006")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: blockadvertisements, Name: "", Queries: q}, nil)
	if err != nil {
		return nil, err
	}
	var blockAdvF []BlockAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&blockAdvF, r)
	if err != nil {
		return nil, err
	}
	blockAdv := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvF))
	for i := range blockAdvF {
		blockAdv = append(blockAdv, ds.convertBlockAdvertisementToDTO(&blockAdvF[i]))
	}
	return blockAdv, err
}

func (ds *ServerStorage) BlockAdvertisementFromReleaseDates(ctx context.Context, from time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching BlockAdvertisements Between releaseDates", "from", from)
	q := make(map[string]string)
	q["fromdate"] = from.Format("02.01.2006")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: blockadvertisements, Name: "", Queries: q}, nil)
	if err != nil {
		return nil, err
	}
	var blockAdvF []BlockAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&blockAdvF, r)
	if err != nil {
		return nil, err
	}
	blockAdv := make([]datatransferobjects.BlockAdvertisementDTO, 0, len(blockAdvF))
	for i := range blockAdvF {
		blockAdv = append(blockAdv, ds.convertBlockAdvertisementToDTO(&blockAdvF[i]))
	}
	return blockAdv, err
}

func (ds *ServerStorage) LineAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching LineAdvertisements Between releaseDates", "from", from, "to", to)
	q := make(map[string]string)
	q["between"] = from.Format("02.01.2006") + "-" + to.Format("02.01.2006")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: "", Queries: q}, nil)
	if err != nil {
		return nil, err
	}
	var lineAdvF []LineAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&lineAdvF, r)
	if err != nil {
		return nil, err
	}
	lineadv := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvF))
	for i := range lineAdvF {
		lineadv = append(lineadv, ds.convertLineAdvertisementToDTO(&lineAdvF[i]))
	}
	return lineadv, err
}

func (ds *ServerStorage) LineAdvertisementFromReleaseDates(ctx context.Context, from time.Time) ([]datatransferobjects.LineAdvertisementDTO, error) {
	logging.Logger.Debug("orm: search request. Searching LineAdvertisements Between releaseDates", "from", from)
	q := make(map[string]string)
	q["fromdate"] = from.Format("02.01.2006")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: lineadvertisements, Name: "", Queries: q}, nil)
	if err != nil {
		return nil, err
	}
	var lineAdvF []LineAdvertisementFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&lineAdvF, r)
	if err != nil {
		return nil, err
	}
	lineadv := make([]datatransferobjects.LineAdvertisementDTO, 0, len(lineAdvF))
	for i := range lineAdvF {
		lineadv = append(lineadv, ds.convertLineAdvertisementToDTO(&lineAdvF[i]))
	}
	return lineadv, err
}

func (ds *ServerStorage) AllTags(ctx context.Context) ([]datatransferobjects.TagDTO, error) {
	logging.Logger.Debug("orm: search request. Getting all Tags")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: tags, Name: ""}, nil)
	if err != nil {
		return nil, err
	}
	var tagF []TagFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&tagF, r)
	if err != nil {
		return nil, err
	}
	tags := make([]datatransferobjects.TagDTO, 0, len(tagF))
	for i := range tagF {
		tags = append(tags, ds.convertTagToDTO(&tagF[i]))
	}
	return tags, err
}

func (ds *ServerStorage) AllExtraCharges(ctx context.Context) ([]datatransferobjects.ExtraChargeDTO, error) {
	logging.Logger.Debug("orm: search request. Getting all ExtraCharges")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: extraCharges, Name: ""}, nil)
	if err != nil {
		return nil, err
	}
	var extraCharfeF []ExtraChargeFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&extraCharfeF, r)
	if err != nil {
		return nil, err
	}
	extraCharfes := make([]datatransferobjects.ExtraChargeDTO, 0, len(extraCharfeF))
	for i := range extraCharfeF {
		extraCharfes = append(extraCharfes, ds.convertExtraChargeToDTO(&extraCharfeF[i]))
	}
	return extraCharfes, err
}

func (ds *ServerStorage) AllCostRates(ctx context.Context) ([]datatransferobjects.CostRateDTO, error) {
	logging.Logger.Debug("orm: search request. Getting All CostRates")
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: costRates, Name: ""}, nil)
	if err != nil {
		return nil, err
	}
	var costRateF []CostRateFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&costRateF, r)
	if err != nil {
		return nil, err
	}
	costRates := make([]datatransferobjects.CostRateDTO, 0, len(costRateF))
	for i := range costRateF {
		costRates = append(costRates, ds.convertCostRateToDto(&costRateF[i]))
	}
	return costRates, err
}

func (ds *ServerStorage) AllFiles(ctx context.Context) ([]datatransferobjects.FileDTO, error) {
	logging.Logger.Debug("orm: search request. Getting All Files")
	query := make(map[string]string)
	query["miniatures"] = "1"
	got, err := ds.s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: files, Name: "", Queries: query}, nil)
	if err != nil {
		return nil, err
	}
	var FilesFront []FileFront
	r := bytes.NewReader(got)
	err = encodedecoder.FromJSON(&FilesFront, r)
	if err != nil {
		return nil, err
	}

	files := make([]datatransferobjects.FileDTO, 0, len(FilesFront))
	for i := range FilesFront {
		decodedFile, err := ds.convertFileToDto(&FilesFront[i])
		if err != nil {
			logging.Logger.Error("server: cannot convert file to data transfer object. skipping", "err", err)
			continue
		}
		files = append(files, decodedFile)
	}
	return files, err
}
