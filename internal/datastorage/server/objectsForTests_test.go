package server

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

var (
	databases = []string{"Server"}

	timeNow          = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	clientModel      = ClientFront{Name: "Вася", Phones: "88005553535", Email: "wowo@asdads.com", AdditionalInformation: "bnmbnmbnm"}
	clientDtoForTest = datatransferobjects.ClientDTO{Name: "Вася", Phones: "88005553535", Email: "wowo@asdads.com", AdditionalInformation: "bnmbnmbnm"}

	orderModel      = OrderFront{ID: 1, ClientName: clientModel.Name, Cost: 123, CreatedDate: timeNow, PaymentStatus: false}
	orderDtoForTest = datatransferobjects.OrderDTO{ID: 1, ClientName: clientModel.Name, Cost: 123, CreatedDate: timeNow, PaymentStatus: false}

	tagsDtoForTest = []datatransferobjects.TagDTO{{TagName: "Tag A", TagCost: 123}, {TagName: "Tag B", TagCost: 456}, {TagName: "Tag C", TagCost: 789}}
	tagsModel      = []TagFront{{TagName: "Tag A", TagCost: 123}, {TagName: "Tag B", TagCost: 456}, {TagName: "Tag C", TagCost: 789}}

	extraChargesModel      = []ExtraChargeFront{{ChargeName: "Charge A", Multiplier: 1}, {ChargeName: "Charge B", Multiplier: 2}, {ChargeName: "Charge C", Multiplier: 3}}
	extraChargesDtoForTest = []datatransferobjects.ExtraChargeDTO{{ChargeName: "Charge A", Multiplier: 1}, {ChargeName: "Charge B", Multiplier: 2}, {ChargeName: "Charge C", Multiplier: 3}}
	costRateModel          = CostRateFront{Name: "asd", ForOneWordSymbol: 1, ForOnecm2: 2, CalcForOneWord: true}
	costRateDto            = datatransferobjects.CostRateDTO{Name: "dsa", ForOneWordSymbol: 4, ForOneSquare: 66, CalcForOneWord: true}
	releaseDate            = time.Now()

	blockModel = BlockAdvertisementFront{
		Advertisement: Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"Tag A", "Tag B", "Tag C"},
			ExtraCharges: []string{"Charge A", "Charge B", "Charge C"},
			ReleaseDates: []time.Time{timeNow},
		},
		Size: 10,
	}
	blockDtoForTest = datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           blockModel.ID,
			OrderID:      blockModel.OrderID,
			ReleaseCount: blockModel.ReleaseCount,
			Cost:         blockModel.Cost,
			Text:         blockModel.Text,
			Tags:         []string{"Tag A", "Tag B", "Tag C"},
			ExtraCharges: []string{"Charge A", "Charge B", "Charge C"},
			ReleaseDates: []time.Time{timeNow},
		},
		Size: blockModel.Size,
	}

	lineModel = LineAdvertisementFront{
		Advertisement: Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"Tag A", "Tag B", "Tag C"},
			ExtraCharges: []string{"Charge A", "Charge B", "Charge C"},
			ReleaseDates: []time.Time{timeNow},
		},
	}
	lineDtoForTest = datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           blockModel.ID,
			OrderID:      blockModel.OrderID,
			ReleaseCount: blockModel.ReleaseCount,
			Cost:         blockModel.Cost,
			Text:         blockModel.Text,
			Tags:         []string{"Tag A", "Tag B", "Tag C"},
			ExtraCharges: []string{"Charge A", "Charge B", "Charge C"},
			ReleaseDates: []time.Time{timeNow},
		},
	}
)
