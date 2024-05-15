package advertisements

import (
	"errors"
	"slices"
	"sort"
	"time"
)

// Общие параметры объявлений
type advertisement struct {
	id           int
	reseaseCount int16
	orderId      int
	cost         int
	releaseDates []time.Time
	extraCharge  []string
	tags         []string
}

func newAdvertisement() advertisement {
	return advertisement{}
}

func (a *advertisement) AppendReleaseDate(date time.Time) {
	for _, v := range a.releaseDates {
		if v == date {
			return
		}
	}
	sort.Slice(a.releaseDates, func(i, j int) bool {
		return a.releaseDates[i].Before(a.releaseDates[j])
	})
	sortedDates := Dates{ReleaseDates: []time.Time{date}}
	sortedDates.ReleaseDates = append(sortedDates.ReleaseDates, a.releaseDates...)

	sort.Sort(&sortedDates)
	a.releaseDates = sortedDates.ReleaseDates
	a.reseaseCount++
}

func (a *advertisement) SetReleaseDates(dates []time.Time) error {
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})

	a.releaseDates = slices.CompactFunc(dates, func(t1, t2 time.Time) bool { return t1.Equal(t2) })
	a.reseaseCount = int16(len(a.releaseDates))
	return nil
}

func (a *advertisement) AppendReleaseDateNew(date time.Time) {
	for _, v := range a.releaseDates {
		if v == date {
			return
		}
	}
	a.releaseDates = append(a.releaseDates, date)
	sort.Slice(a.releaseDates, func(i, j int) bool {
		return a.releaseDates[i].Before(a.releaseDates[j])
	})

	a.reseaseCount++
}

func (a *advertisement) RemoveReleaseDate(date time.Time) {
	newReleaseDates := make([]time.Time, 0, len(a.releaseDates))
	for i := range a.releaseDates {
		if a.releaseDates[i] == date {
			newReleaseDates = append(newReleaseDates, a.releaseDates[:i]...)
			if len(a.releaseDates)-1 > i {
				newReleaseDates = append(newReleaseDates, a.releaseDates[i+1:]...)
			}
			a.releaseDates = newReleaseDates
			return
		}
	}
	if a.reseaseCount != 0 {
		a.reseaseCount--
	}
}

func (a *advertisement) SetOrderId(i int) error {
	if i < 0 {
		return errors.New("заказ должен быть больше или равен 0")
	}
	a.orderId = i
	return nil
}
func (a *advertisement) SetId(i int) error {
	if i < 0 {
		return errors.New("заказ должен быть больше или равен 0")
	}
	a.id = i
	return nil
}

func (a *advertisement) SetCost(cost int) error {
	if cost < 0 {
		return errors.New("стоимость заказа должно быть больше или равен 0")
	}
	a.cost = cost
	return nil
}

func (a *advertisement) Id() int {
	return a.id
}

func (a *advertisement) OrderId() int {
	return a.orderId
}

func (a *advertisement) ReleaseDates() []time.Time {
	return a.releaseDates
}

func (a *advertisement) Cost() int {
	return a.cost
}

func (a *advertisement) Tags() []string {
	return a.tags
}

func (a *advertisement) ReseaseCount() int16 {
	return a.reseaseCount
}

func (a *advertisement) AppendTag(name string) {
	for _, v := range a.tags {
		if v == name {
			return
		}
	}
	a.tags = append(a.tags, name)
}

func (a *advertisement) SetTags(tags []string) {
	slices.Sort(tags)
	a.tags = slices.Compact(tags)
}

func (a *advertisement) RemoveTag(name string) {
	NewTags := make([]string, 0, len(a.tags))
	for i := range a.tags {
		if a.tags[i] == name {
			NewTags = append(NewTags, a.tags[:i]...)
			if len(a.tags)-1 > i {
				NewTags = append(NewTags, a.tags[i+1:]...)
			}
			a.tags = NewTags
			return
		}
	}
}
func (a *advertisement) ExtraCharge() []string {
	return a.extraCharge
}

func (a *advertisement) AppendExtraCharge(name string) {
	for _, v := range a.extraCharge {
		if v == name {
			return
		}
	}
	a.extraCharge = append(a.extraCharge, name)
}

func (a *advertisement) SetExtraCharges(extraCharges []string) {
	slices.Sort(extraCharges)
	a.extraCharge = slices.Compact(extraCharges)
}

func (a *advertisement) RemoveExtraCharge(name string) {
	NewExtraCharge := make([]string, 0, len(a.extraCharge))
	for i := range a.extraCharge {
		if a.extraCharge[i] == name {
			NewExtraCharge = append(NewExtraCharge, a.extraCharge[:i]...)
			if len(a.extraCharge)-1 > i {
				NewExtraCharge = append(NewExtraCharge, a.extraCharge[i+1:]...)
			}
			a.extraCharge = NewExtraCharge
			return
		}
	}
}
