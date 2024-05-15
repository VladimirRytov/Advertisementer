package server

import "time"

type ClientFront struct {
	Name                  string       `json:"name"`
	Phones                string       `json:"phones"`
	Email                 string       `json:"email"`
	AdditionalInformation string       `json:"additionalInformation,omitempty"`
	Orders                []OrderFront `json:"orders,omitempty"`
}

type OrderFront struct {
	ID                  int                       `json:"id"`
	ClientName          string                    `json:"clientName"`
	Cost                int                       `json:"cost"`
	PaymentType         string                    `json:"paymentType,omitempty"`
	CreatedDate         time.Time                 `json:"createdDate"`
	PaymentStatus       bool                      `json:"paymentStatus"`
	LineAdvertisements  []LineAdvertisementFront  `json:"lineAdvertisements,omitempty"`
	BlockAdvertisements []BlockAdvertisementFront `json:"blockAdvertisements,omitempty"`
}

type Advertisement struct {
	ID           int         `json:"id"`
	OrderID      int         `json:"orderID"`
	ReleaseCount int16       `json:"releaseCount"`
	Cost         int         `json:"cost"`
	Text         string      `json:"text,omitempty"`
	Tags         []string    `json:"tags,omitempty"`
	ExtraCharges []string    `json:"extraCharges,omitempty"`
	ReleaseDates []time.Time `json:"releaseDates,omitempty"`
}

type BlockAdvertisementFront struct {
	Advertisement
	Size     int16  `json:"size"`
	FileName string `json:"fileName,omitempty"`
}

type LineAdvertisementFront struct {
	Advertisement
}

type TagFront struct {
	TagName string `json:"name"`
	TagCost int    `json:"cost"`
}

type ExtraChargeFront struct {
	ChargeName string `json:"name"`
	Multiplier int    `json:"multiplier"`
}

type CostRateFront struct {
	CalcForOneWord   bool   `json:"calcForOneWord"`
	Name             string `json:"name"`
	ForOneWordSymbol int    `json:"forOneWordSymbol"`
	ForOnecm2        int    `json:"forOnecm2"`
}

type UserFront struct {
	Name        string `json:"name"`
	Password    []byte `json:"password,omitempty"`
	Permissions byte   `json:"permissions,omitempty"`
}

type FileFront struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
	Data string `json:"data"`
}
