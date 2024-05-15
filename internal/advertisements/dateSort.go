package advertisements

import "time"

type Dates struct {
	ReleaseDates []time.Time
}

func (d *Dates) Len() int {
	return len(d.ReleaseDates)
}

func (d *Dates) Less(i, j int) bool {
	return d.ReleaseDates[i].Before(d.ReleaseDates[j])
}

func (d *Dates) Swap(i, j int) {
	d.ReleaseDates[i], d.ReleaseDates[j] = d.ReleaseDates[j], d.ReleaseDates[i]
}
