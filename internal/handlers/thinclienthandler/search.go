package thinclienthandler

import (
	"context"
	"errors"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (a *AdvertisementController) AllClients() {
	logging.Logger.Debug("advertisementhandler: search request. Search All Clients")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	recievedClientsDto, err := a.req.AllClients(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllClients: cannot get clients", "error", err)
		a.handleError(err)
		return
	}
	for i := range recievedClientsDto {
		a.app.SendClient(&recievedClientsDto[i])
	}
}

func (a *AdvertisementController) AllOrders() {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllOrders")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	orders, err := a.req.AllOrders(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllOrders: cannot get orders", "error", err)
		a.handleError(err)
		return
	}
	for i := range orders {
		a.app.SendAdvertisementsOrder(&orders[i])
	}
}

func (a *AdvertisementController) AllBlockAdvertisements() {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllBlockAdvertisements")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.AllBlockAdvertisements(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllBlockAdvertisements: cannot get all BlockAdvertisements", "error", err)
		a.handleError(err)
		return
	}
	for i := range blockAdvs {
		a.app.SendBlockAdvertisement(&blockAdvs[i])
	}
}

func (a *AdvertisementController) AllLineAdvertisements() {
	logging.Logger.Debug("advertisementhandler: search request. Search All AllLineAdvertisements")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.AllLineAdvertisements(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllLineAdvertisements: cannot get all LineAdvertisements", "error", err)
		a.handleError(err)
	}

	for i := range lineAdvs {
		a.app.SendLineAdvertisement(&lineAdvs[i])
	}
}

func (a *AdvertisementController) OrdersByClientName(name string) {
	logging.Logger.Debug("advertisementhandler: search request. Search Order by Client name")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	orders, err := a.req.OrdersByClientName(ctx, name)
	if err != nil {
		logging.Logger.Error("advertisementhandler.OrdersByClientName: cannot get OrdersByClientName", "error", err)
		a.handleError(err)
	}
	for i := range orders {
		a.app.SendAdvertisementsOrder(&orders[i])
	}
}

func (a *AdvertisementController) BlockAdvertisementsByOrderID(id int) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements by Order id", "id", id)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.BlockAdvertisementsByOrderID(ctx, id)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsByOrderID: cannot get BlockAdvertisementsByOrderID", "error", err)
		a.handleError(err)
	}

	for i := range blockAdvs {
		a.app.SendBlockAdvertisement(&blockAdvs[i])
	}
}

func (a *AdvertisementController) LineAdvertisementsByOrderID(id int) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements by Order id", "id", id)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.LineAdvertisementsByOrderID(ctx, id)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsByOrderID: cannot get LineAdvertisementsByOrderID", "error", err)
		a.handleError(err)
	}

	for i := range lineAdvs {
		a.app.SendLineAdvertisement(&lineAdvs[i])
	}
}

func (a *AdvertisementController) BlockAdvertisementsBetweenReleaseDates(from, to time.Time) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements Between release dates", "from", from, "to", to)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.BlockAdvertisementBetweenReleaseDates(ctx, from, to)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsBetweenReleaseDates: cannot get BlockAdvertisementsBetweenReleaseDates", "error", err)
		a.handleError(err)
	}

	for i := range blockAdvs {
		a.app.SendBlockAdvertisement(&blockAdvs[i])
	}
}

func (a *AdvertisementController) LineAdvertisementsBetweenReleaseDates(from, to time.Time) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements Between release dates", "from", from, "to", to)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.LineAdvertisementBetweenReleaseDates(ctx, from, to)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsBetweenReleaseDates: cannot get LineAdvertisementsBetweenReleaseDates", "error", err)
		a.handleError(err)
	}

	for i := range lineAdvs {
		a.app.SendLineAdvertisement(&lineAdvs[i])
	}
}

func (a *AdvertisementController) BlockAdvertisementsActualReleaseDate() {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements with actual release dates", "from", time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.BlockAdvertisementFromReleaseDates(ctx, time.Now())
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsActualReleaseDate: cannot get BlockAdvertisementsActualReleaseDate", "error", err)
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			a.app.SendError(errors.New(CancelStr))
		case errors.Is(err, context.Canceled):
			a.app.SendError(errors.New(TimeoutStr))
		default:
			a.app.SendError(err)
		}
	}

	for i := range blockAdvs {
		a.app.SendBlockAdvertisement(&blockAdvs[i])
	}
}

func (a *AdvertisementController) LineAdvertisementsActualReleaseDate() {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements with actual release dates", "from", time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.LineAdvertisementFromReleaseDates(ctx, time.Now())
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsActualReleaseDate: cannot get LineAdvertisementsActualReleaseDate", "error", err)
		a.handleError(err)
	}

	for i := range lineAdvs {
		a.app.SendLineAdvertisement(&lineAdvs[i])
	}
}

func (a *AdvertisementController) BlockAdvertisementsFromReleaseDate(from time.Time) {
	logging.Logger.Debug("advertisementhandler: search request. Search BlockAdvertisements from release dates", "from", from)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	blockAdvs, err := a.req.BlockAdvertisementFromReleaseDates(ctx, from)
	if err != nil {
		logging.Logger.Error("advertisementhandler.BlockAdvertisementsFromReleaseDate: cannot get BlockAdvertisementsFromReleaseDate", "error", err)
		a.handleError(err)
	}

	for i := range blockAdvs {
		a.app.SendBlockAdvertisement(&blockAdvs[i])
	}
}

func (a *AdvertisementController) LineAdvertisementsFromReleaseDate(from time.Time) {
	logging.Logger.Debug("advertisementhandler: search request. Search LineAdvertisements from release dates", "from", from)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	lineAdvs, err := a.req.LineAdvertisementFromReleaseDates(ctx, from)
	if err != nil {
		logging.Logger.Error("advertisementhandler.LineAdvertisementsFromReleaseDate: cannot get LineAdvertisementsFromReleaseDate", "error", err)
		a.handleError(err)
	}

	for i := range lineAdvs {
		a.app.SendLineAdvertisement(&lineAdvs[i])
	}
}

func (a *AdvertisementController) AllTags() {
	logging.Logger.Debug("advertisementhandler: search request. Search all Tags")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	tags, err := a.req.AllTags(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllTags: cannot get AllTags", "error", err)
		a.handleError(err)
	}

	for i := range tags {
		a.app.SendTag(&tags[i])
	}
}

func (a *AdvertisementController) AllExtraCharges() {
	logging.Logger.Debug("advertisementhandler: search request. Search all ExtraCharges")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	extraCharges, err := a.req.AllExtraCharges(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllExtraCharges: cannot get AllExtraCharges", "error", err)
		a.handleError(err)
		return
	}

	for i := range extraCharges {
		a.app.SendExtraCharge(&extraCharges[i])
	}
}

func (a *AdvertisementController) AllCostRates() {
	logging.Logger.Debug("advertisementhandler: search request. Search all CostRates")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	costRates, err := a.req.AllCostRates(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllCostRates: cannot get AllCostRates", "error", err)
		a.handleError(err)
		return
	}

	for i := range costRates {
		a.app.SendCostRate(&costRates[i])
	}
}

func (a *AdvertisementController) AllFiles(ctx context.Context) {
	logging.Logger.Debug("advertisementhandler.AllFiles: search request. Search all Files")
	files, err := a.req.AllFiles(ctx)
	if err != nil {
		logging.Logger.Error("advertisementhandler.AllFiles: cannot get Files", "error", err)
		switch {
		case errors.Is(err, context.Canceled):
			a.app.UnlockFilesWindow()
			return
		case errors.Is(err, context.DeadlineExceeded):
			a.app.SendError(errors.New(TimeoutStr))
		default:
			a.app.SendError(err)
		}
		return
	}

	for i := range files {
		a.app.SendFile(&files[i])
	}
	a.app.UnlockFilesWindow()
}
