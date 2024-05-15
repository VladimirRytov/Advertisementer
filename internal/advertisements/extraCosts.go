package advertisements

import (
	"errors"
)

type CostRate struct {
	calcForOneWord   bool
	name             string
	forOneWordSymbol int
	forOneSquare     int
}

func NewCostRate(name string, costWS, cm int, calcForOneWord bool) (CostRate, error) {
	if len(name) < 2 {
		return CostRate{}, errors.New("имя тарифа должно должно состоять из 3 или больше символов")
	}
	return CostRate{forOneSquare: cm, forOneWordSymbol: costWS, calcForOneWord: calcForOneWord, name: name}, nil
}

func (cr *CostRate) CalsForOneWord() bool {
	return cr.calcForOneWord
}

func (cr *CostRate) SetCalsForOneWord(calc bool) {
	cr.calcForOneWord = calc
}

func (cr *CostRate) Name() string {
	return cr.name
}

func (cr *CostRate) SetName(name string) error {
	if len(name) < 2 {
		return errors.New("имя тарифа должно быть больше 3 символов")
	}
	cr.name = name
	return nil
}

func (cr *CostRate) CostForOnecm2() int {
	return cr.forOneSquare
}

func (cr *CostRate) SetCostForOnecm2(cost int) error {
	if cost < 0 {
		return errors.New("триф за слово болжен быть больше 0")
	}
	cr.forOneSquare = cost
	return nil
}

func (cr *CostRate) CostForWordOrSymbol() int {
	return cr.forOneWordSymbol
}

func (cr *CostRate) SetCostForWordOrSymbol(cost int) {
	cr.forOneWordSymbol = cost
}

func (cr *CostRate) CalculateBlockCost(blockAdv AdvertisementBlock, tags []Tag, charges []ExtraCharge) (int, error) {
	var (
		tagCosts    int
		chargeCosts int = 100
	)
	for i := range tags {
		tagCosts += tags[i].Cost()
	}
	for i := range charges {
		chargeCosts += charges[i].Multiplier()
	}
	total := (int(blockAdv.Size())*cr.forOneSquare + tagCosts) * chargeCosts * int(blockAdv.ReseaseCount())
	if total < 0 {
		return 0, errors.New("стоимость не может быть ниже 0")
	}
	return total / 100, nil
}

func (cr *CostRate) CalculateLineCost(lineAdv AdvertisementLine, tags []Tag, charges []ExtraCharge) (int, error) {
	var (
		tagCosts    int
		chargeCosts int = 100
		total       int
	)
	for i := range tags {
		tagCosts += tags[i].Cost()
	}
	for i := range charges {
		chargeCosts += charges[i].Multiplier()
	}
	if cr.calcForOneWord {
		total = (int(lineAdv.WordsCount()*cr.forOneWordSymbol) + tagCosts)
	} else {
		total = (int(lineAdv.SymbolsCount()*cr.forOneWordSymbol) + tagCosts)
	}
	total *= chargeCosts * int(lineAdv.reseaseCount)
	if total < 0 {
		return 0, errors.New("стоимость не может быть ниже 0")

	}
	return total / 100, nil
}

func (cr *CostRate) CalculateOrderCost(orderAdv AdvertisementOrder, blocks []AdvertisementBlock, lines []AdvertisementLine) int {
	var (
		lineCost  int
		blockCost int
	)
	for i := range blocks {
		blockCost += blocks[i].Cost()
	}
	for i := range lines {
		lineCost += lines[i].Cost()
	}
	return lineCost + blockCost
}

type Tag struct {
	name string
	cost int
}

func NewTag(name string, cost int) (Tag, error) {
	if len(name) == 0 {
		return Tag{}, errors.New("имя тэга должно быть заполнено")
	}
	return Tag{name: name, cost: cost}, nil
}

func (t *Tag) Name() string {
	return t.name
}

func (t *Tag) Cost() int {
	return t.cost
}
func (t *Tag) SetCost(c int) {
	t.cost = c
}

type ExtraCharge struct {
	name       string
	multiplier int
}

func NewExtraCharge(name string, multiplier int) (ExtraCharge, error) {
	if len(name) == 0 {
		return ExtraCharge{}, errors.New("имя наценки должно быть заполнено")
	}
	return ExtraCharge{name: name, multiplier: multiplier}, nil
}

func (ex *ExtraCharge) Multiplier() int {
	return ex.multiplier
}

func (ex *ExtraCharge) Name() string {
	return ex.name
}
func (ex *ExtraCharge) SetMiltiplier(c int) {
	ex.multiplier = c
}
