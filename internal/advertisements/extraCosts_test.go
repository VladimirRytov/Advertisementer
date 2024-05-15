package advertisements

import (
	"testing"
	"time"
)

func TestNewCostRate(t *testing.T) {
	want := CostRate{forOneWordSymbol: 1, forOneSquare: 2}
	got, err := NewCostRate("asd", 1, 2, false)
	if err != nil {
		t.Fatal(err)
	}
	if want.forOneWordSymbol != got.forOneWordSymbol && want.forOneSquare != got.forOneSquare {
		t.Errorf("want %v,got %v", want, got)
	}
}

func TestSetCostForOneWord(t *testing.T) {
	want := CostRate{forOneWordSymbol: 3, forOneSquare: 2}
	got, err := NewCostRate("asd", 1, 2, true)
	if err != nil {
		t.Fatal(err)
	}
	got.SetCostForWordOrSymbol(3)
	if want.forOneWordSymbol != got.forOneWordSymbol {
		t.Errorf("want %v,got %v", want, got)
	}
}

func TestSetCostForOnecm2(t *testing.T) {
	want := &CostRate{forOneWordSymbol: 1, forOneSquare: 5}
	got, err := NewCostRate("asd", 1, 2, true)
	if err != nil {
		t.Fatal(err)
	}
	got.SetCostForOnecm2(5)
	if want.forOneSquare != got.forOneSquare {
		t.Errorf("want %v,got %v", want, got)
	}
}
func TestCostForOneWord(t *testing.T) {
	want := 50
	cost, err := NewCostRate("asd", want, 2, true)
	if err != nil {
		t.Fatal(err)
	}
	got := cost.CostForWordOrSymbol()

	if want != got {
		t.Errorf("want %d,got %d", want, got)
	}
}

func TestCostForOnecm2(t *testing.T) {
	want := 50
	cost, err := NewCostRate("asd", 1, want, true)
	if err != nil {
		t.Fatal(err)
	}
	got := cost.CostForOnecm2()

	if want != got {
		t.Errorf("want %d,got %d", want, got)
	}
}
func TestSetGetCostForSymbol(t *testing.T) {
	want := 3
	Cost, err := NewCostRate("asd", want, 2, false)
	if err != nil {
		t.Fatal(err)
	}
	Cost.SetCostForWordOrSymbol(want)
	if got := Cost.CostForWordOrSymbol(); got != want {
		t.Errorf("want %v,get %v", want, got)
	}
}
func TestNewTag(t *testing.T) {
	want := Tag{name: "a", cost: 4}
	got, _ := NewTag("a", 4)
	if want != got {
		t.Errorf("want %v,get %v", want, got)
	}
}
func TestGetTagValues(t *testing.T) {
	wantName := "a"
	wantCost := 3
	got, _ := NewTag(wantName, wantCost)
	gotName := got.Name()
	gotCost := got.Cost()
	if wantName != gotName {
		t.Errorf("want %v,get %v", wantName, gotName)
	}
	if wantCost != gotCost {
		t.Errorf("want %v,get %v", wantCost, gotCost)
	}
}
func TestSetCostTag(t *testing.T) {
	wantCost := 3
	got, _ := NewTag("a", 1)
	got.SetCost(wantCost)
	if cost := got.Cost(); cost != wantCost {
		t.Errorf("want %v,get %v", wantCost, cost)
	}
}

func TestNewExtraCharge(t *testing.T) {
	want := ExtraCharge{name: "a", multiplier: 4}
	got, _ := NewExtraCharge("a", 4)
	if want != got {
		t.Errorf("want %v,get %v", want, got)
	}
}
func TestGetExtraCharge(t *testing.T) {
	wantName := "a"
	wantMultiplier := 3
	got, _ := NewExtraCharge(wantName, wantMultiplier)
	gotName := got.Name()
	gotMultiplier := got.Multiplier()
	if wantName != gotName {
		t.Errorf("want %v,get %v", wantName, gotName)
	}
	if wantMultiplier != gotMultiplier {
		t.Errorf("want %v,get %v", wantMultiplier, gotMultiplier)
	}
}

func TestSetMiltiplierExtraCharge(t *testing.T) {
	wantMultiplier := 3
	got, _ := NewExtraCharge("a", 1)
	got.SetMiltiplier(wantMultiplier)
	if gotMultiplier := got.Multiplier(); gotMultiplier != wantMultiplier {
		t.Errorf("want %v,get %v", wantMultiplier, gotMultiplier)
	}
}
func TestCalculateLineAdvertisementCostSymbolCount(t *testing.T) {
	tagsForTest := []Tag{{name: "Tag A", cost: 100}, {name: "Tag B", cost: 200}, {name: "Tag C", cost: 300}}
	extraChargesForTest := []ExtraCharge{{name: "Charge A", multiplier: 10}, {name: "Charge B", multiplier: 20}, {name: "Charge C", multiplier: 70}}
	lineForTest := AdvertisementLine{
		advertisement: advertisement{
			id:           1,
			orderId:      1,
			reseaseCount: 1,
			cost:         23,
			tags:         []string{"Tag A", "Tag B", "Tag C"},
			extraCharge:  []string{"Charge A", "Charge B", "Charge C"},
			releaseDates: []time.Time{time.Now()},
		},
		message: "Text",
	}
	c, err := NewCostRate("dsa", 4, 66, false)
	if err != nil {
		t.Fatal(err)
	}
	cost, err := c.CalculateLineCost(lineForTest, tagsForTest, extraChargesForTest)
	if err != nil {
		t.Fatal(err)
	}
	if cost != 1232 {
		t.Fatalf("want cost 1232,got %d", cost)
	}
}

func TestCalculateLineAdvertisementCostWordCount(t *testing.T) {
	tagsForTest := []Tag{{name: "Tag A", cost: 100}, {name: "Tag B", cost: 200}, {name: "Tag C", cost: 300}}
	extraChargesForTest := []ExtraCharge{{name: "Charge A", multiplier: 10}, {name: "Charge B", multiplier: 20}, {name: "Charge C", multiplier: 70}}
	lineForTest := AdvertisementLine{
		advertisement: advertisement{
			id:           1,
			orderId:      1,
			reseaseCount: 1,
			cost:         23,
			tags:         []string{"Tag A", "Tag B", "Tag C"},
			extraCharge:  []string{"Charge A", "Charge B", "Charge C"},
			releaseDates: []time.Time{time.Now()},
		},
		message: "Text",
	}
	c, err := NewCostRate("dsa", 4, 66, true)
	if err != nil {
		t.Fatal(err)
	}
	cost, err := c.CalculateLineCost(lineForTest, tagsForTest, extraChargesForTest)
	if err != nil {
		t.Fatal(err)
	}
	if cost != 1208 {
		t.Fatalf("want cost 1208, got %d", cost)
	}
}

func TestCalculateBlockAdvertisementCost(t *testing.T) {
	tagsForTest := []Tag{{name: "Tag A", cost: 100}, {name: "Tag B", cost: 200}, {name: "Tag C", cost: 300}}
	extraChargesForTest := []ExtraCharge{{name: "Charge A", multiplier: 10}, {name: "Charge B", multiplier: 20}, {name: "Charge C", multiplier: 70}}
	blockForTest := AdvertisementBlock{
		advertisement: advertisement{
			id:           1,
			orderId:      1,
			reseaseCount: 1,
			cost:         23,
			tags:         []string{"Tag A", "Tag B", "Tag C"},
			extraCharge:  []string{"Charge A", "Charge B", "Charge C"},
			releaseDates: []time.Time{time.Now()},
		},
		comment: "Text",
		size:    1,
	}
	c, err := NewCostRate("dsa", 4, 66, true)
	if err != nil {
		t.Fatal(err)
	}
	cost, err := c.CalculateBlockCost(blockForTest, tagsForTest, extraChargesForTest)
	if err != nil {
		t.Fatal(err)
	}
	if cost != 1332 {
		t.Fatalf("want cost 1332, got %d", cost)
	}
}

func TestCalculateOrderCost(t *testing.T) {
	blockForTest := AdvertisementBlock{
		advertisement: advertisement{
			id:           1,
			orderId:      1,
			reseaseCount: 1,
			cost:         200,
			tags:         []string{"Tag A", "Tag B", "Tag C"},
			extraCharge:  []string{"Charge A", "Charge B", "Charge C"},
			releaseDates: []time.Time{time.Now()},
		},
		comment: "Text",
		size:    1,
	}
	lineForTest := AdvertisementLine{
		advertisement: advertisement{
			id:           1,
			orderId:      1,
			reseaseCount: 1,
			cost:         400,
			tags:         []string{"Tag A", "Tag B", "Tag C"},
			extraCharge:  []string{"Charge A", "Charge B", "Charge C"},
			releaseDates: []time.Time{time.Now()},
		},
		message: "Text",
	}

	c, err := NewCostRate("dsa", 4, 66, true)
	if err != nil {
		t.Fatal(err)
	}
	orderDtoForTest := AdvertisementOrder{id: 0, clientId: "Вася", cost: 123, createdDate: time.Now()}

	cost := c.CalculateOrderCost(orderDtoForTest, []AdvertisementBlock{blockForTest}, []AdvertisementLine{lineForTest})
	if err != nil {
		t.Fatal(err)
	}
	if cost != 600 {
		t.Fatalf("want cost 600, got %d", cost)
	}
}
