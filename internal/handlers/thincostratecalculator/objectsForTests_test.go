package thincostratecalculator

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

var (
	timeNow          = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	clientDtoForTest = datatransferobjects.ClientDTO{Name: "Вася", Phones: "123", Email: "wowo@asd.ru"}

	orderDtoForTest = datatransferobjects.OrderDTO{ID: 1, ClientName: "Вася", Cost: 123, CreatedDate: timeNow}

	tagsDtoForTest = []datatransferobjects.TagDTO{{TagName: "Tag A", TagCost: 100}, {TagName: "Tag B", TagCost: 200}, {TagName: "Tag C", TagCost: 300}}

	extraChargesDtoForTest = []datatransferobjects.ExtraChargeDTO{{ChargeName: "Charge A", Multiplier: 10}, {ChargeName: "Charge B", Multiplier: 20}, {ChargeName: "Charge C", Multiplier: 70}}
	costRateDto            = datatransferobjects.CostRateDTO{Name: "dsa", ForOneWordSymbol: 4, ForOneSquare: 66, CalcForOneWord: true}

	blockDtoForTest = datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"Tag A", "Tag B", "Tag C"},
			ExtraCharges: []string{"Charge A", "Charge B", "Charge C"},
			ReleaseDates: []time.Time{timeNow, time.Date(2023, 11, 12, 0, 0, 0, 0, time.UTC)},
		},
		Size: 1,
	}

	lineDtoForTest = datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         23,
			Text:         "Text",
			Tags:         []string{"Tag A", "Tag B", "Tag C"},
			ExtraCharges: []string{"Charge A", "Charge B", "Charge C"},
			ReleaseDates: []time.Time{timeNow},
		},
	}
)
