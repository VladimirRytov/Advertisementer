package datatransferobjects

import "time"

type JsonStr struct {
	Clients      []ClientDTO      `json:"clients"`
	Tags         []TagDTO         `json:"tags"`
	ExtraCharges []ExtraChargeDTO `json:"extraCharges"`
	CostRates    []CostRateDTO    `json:"costRates"`
}

type ImportParams struct {
	FilePath      string
	ActualClients bool

	AllBlocks       bool
	AlllLines       bool
	AllTags         bool
	AllExtraCharges bool
	AllCostRates    bool
	IgnoreErrors    bool
	ThickMode       bool
}

type ReportParams struct {
	ReportType       string
	FromDate, ToDate time.Time
	BlocksFolderPath string
	DeployPath       string
}

type ClientDTO struct {
	Name                  string `json:"name"`
	Phones                string `json:"phones"`
	Email                 string `json:"email"`
	AdditionalInformation string `json:"additionalInformation"`
	Orders                []OrderDTO
}

type OrderDTO struct {
	ID                  int                     `json:"id"`
	ClientName          string                  `json:"clientName"`
	Cost                int                     `json:"cost"`
	PaymentType         string                  `json:"paymentType"`
	CreatedDate         time.Time               `json:"createdDate"`
	PaymentStatus       bool                    `json:"paymentStatus"`
	LineAdvertisements  []LineAdvertisementDTO  `json:"lineAdvertisements"`
	BlockAdvertisements []BlockAdvertisementDTO `json:"blockAdvertisements"`
}

type Advertisement struct {
	ID           int         `json:"id"`
	OrderID      int         `json:"orderID"`
	ReleaseCount int16       `json:"releaseCount"`
	Cost         int         `json:"cost"`
	Text         string      `json:"text"`
	Tags         []string    `json:"tags"`
	ExtraCharges []string    `json:"extraCharges"`
	ReleaseDates []time.Time `json:"releaseDates"`
}

type BlockAdvertisementDTO struct {
	Advertisement
	Size     int16  `json:"size"`
	FileName string `json:"fileName"`
}

type LineAdvertisementDTO struct {
	Advertisement
}

type TagDTO struct {
	TagName string `json:"name"`
	TagCost int    `json:"cost"`
}

type RequestDTO struct {
	Kind    string
	Name    string
	Queries map[string]string
}

type ExtraChargeDTO struct {
	ChargeName string `json:"name"`
	Multiplier int    `json:"multiplier"`
}

type CostRateDTO struct {
	CalcForOneWord   bool   `json:"calcForOneWord"`
	Name             string `json:"name"`
	ForOneWordSymbol int    `json:"forOneWordSymbol"`
	ForOneSquare     int    `json:"forOnecm2"`
}
