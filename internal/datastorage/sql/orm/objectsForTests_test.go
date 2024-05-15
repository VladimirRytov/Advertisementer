package orm

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

var (
	databases = []string{"Sql Server", "Postgres", "Sqlite", "Mysql"}

	timeNow          = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	clientModel      = Client{Name: "Вася", Phones: "123", Email: "wowo", AdditionalInformation: "bnmbnmbnm"}
	clientDtoForTest = datatransferobjects.ClientDTO{Name: "Вася", Phones: "123", Email: "wowo", AdditionalInformation: "bnmbnmbnm"}

	orderModel      = Order{ID: 1, ClientName: clientModel.Name, Cost: 123, CreatedDate: timeNow, PaymentStatus: false}
	orderDtoForTest = datatransferobjects.OrderDTO{ID: 1, ClientName: clientModel.Name, Cost: 123, CreatedDate: timeNow, PaymentStatus: false}

	tagsDtoForTest = []datatransferobjects.TagDTO{{TagName: "Tag A", TagCost: 123}, {TagName: "Tag B", TagCost: 456}, {TagName: "Tag C", TagCost: 789}}
	tagsModel      = []Tag{{Name: "Tag A", Cost: 123}, {Name: "Tag B", Cost: 456}, {Name: "Tag C", Cost: 789}}

	extraChargesModel      = []ExtraCharge{{Name: "Charge A", Multiplier: 1}, {Name: "Charge B", Multiplier: 2}, {Name: "Charge C", Multiplier: 3}}
	extraChargesDtoForTest = []datatransferobjects.ExtraChargeDTO{{ChargeName: "Charge A", Multiplier: 1}, {ChargeName: "Charge B", Multiplier: 2}, {ChargeName: "Charge C", Multiplier: 3}}
	costRateModel          = CostRate{Name: "asd", OneWordOrSymbol: 1, OneSquare: 2, CalcForOneWord: true}
	costRateDto            = datatransferobjects.CostRateDTO{Name: "dsa", ForOneWordSymbol: 4, ForOneSquare: 66, CalcForOneWord: true}
	releaseDate            = time.Now()
	releaseDateModel       = &ReleaseDates{ReleaseDate: releaseDate}

	blockModel = AdvertisementBlock{
		Advertisement: Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         23,
			Text:         "asd",
		},
		Size:         10,
		Tags:         tagsModel,
		ExtraCharges: extraChargesModel,
		ReleaseDates: []ReleaseDates{{timeNow}},
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

	lineModel = AdvertisementLine{
		Advertisement: Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         23,
			Text:         "asd",
		},
		Tags:         tagsModel,
		ExtraCharges: extraChargesModel,
		ReleaseDates: []ReleaseDates{{timeNow}},
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
