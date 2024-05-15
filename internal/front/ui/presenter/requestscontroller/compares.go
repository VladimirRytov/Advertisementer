package requestscontroller

import (
	"strings"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

var LocaleSorter = collate.New(language.Russian)

func (r *RequestsHandler) ValueInString(arr, val string) bool {
	tagsArray := strings.Split(arr, ", ")
	for _, v := range tagsArray {
		if val == v {
			return true
		}
	}
	return false
}

func (r *RequestsHandler) CompareStringTime(a, b string) int {
	selectedTime, err := r.converter.DateToDto(a)
	if err != nil {
		logging.Logger.Warn("cannot convert time string to time", "error", err)
	}
	compareTime, err := r.converter.DateToDto(b)
	if err != nil {
		logging.Logger.Warn("cannot convert time string to time", "error", err)
	}
	return selectedTime.Compare(compareTime)
}

func (r *RequestsHandler) ReleasesInTimeRange(release, from, to string) bool {
	var (
		err      error
		fromTime = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)
		toTime   = time.Date(9999, 1, 1, 0, 0, 0, 0, time.Local)
	)
	if len(from) > 0 {
		fromTime, err = r.converter.DateToDto(from)
		if err != nil {
			logging.Logger.Warn("cannot convert time string to time", "error", err)
		}
	}
	if len(to) > 0 {
		toTime, err = r.converter.DateToDto(to)
		if err != nil {
			logging.Logger.Warn("cannot convert time string to time", "error", err)
		}
	}

	stringReleaseDates := strings.Split(release, ", ")
	releaseDates := make([]time.Time, 0, len(stringReleaseDates))
	for _, v := range stringReleaseDates {
		releaseDate, err := r.converter.DateToDto(v)
		if err != nil {
			continue
		}
		releaseDates = append(releaseDates, releaseDate)
	}

	closest, err := r.converter.ClosestRelease(releaseDates, fromTime)
	if err != nil {
		return false
	}
	return closest.Before(toTime)
}

func (r *RequestsHandler) CompareSelected(a, b bool) int {
	switch {
	case a == b:
		return 0
	case a:
		return -1
	default:
		return 1
	}
}

func (r *RequestsHandler) CheckCostString(cost string) bool {
	return len(r.costRegex.FindString(cost)) > 0
}

func (r *RequestsHandler) CheckAdvCostString(cost string) bool {
	return len(r.costAdvRegex.FindString(cost)) > 0
}

func (r *RequestsHandler) CompareInts(a, b int) int {
	return a - b
}

func (r *RequestsHandler) CompareStringsOld(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func (r *RequestsHandler) CompareStrings(a, b string) int {
	return r.collator.Compare([]byte(a), []byte(b))
}

func (r *RequestsHandler) CompareCosts(a, b string) int {
	intA, err := r.converter.CostToDto(a)
	if err != nil {
		logging.Logger.Error("CompareCosts: cannot convert costA to int", "error", err)
		return 0
	}

	intB, err := r.converter.CostToDto(b)
	if err != nil {
		logging.Logger.Error("CompareCosts: cannot convert costB to int", "error", err)
		return 0
	}
	return intA - intB
}
