package orm

import (
	"errors"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datastorage"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"gorm.io/gorm"
)

func fetchTagsName(t []Tag) []string {
	logging.Logger.Debug("orm: Fetching Tag name")

	var tagSlice = make([]string, 0, len(t))
	for _, v := range t {
		tagSlice = append(tagSlice, v.Name)
	}
	return tagSlice
}

func fetchExtraChargesName(t []ExtraCharge) []string {
	logging.Logger.Debug("orm: Fetching ExtraCharge name")
	var extraChargeSlice = make([]string, 0, len(t))
	for _, v := range t {
		extraChargeSlice = append(extraChargeSlice, v.Name)
	}
	return extraChargeSlice
}

func fetchReleaseDates(t []ReleaseDates) []time.Time {
	logging.Logger.Debug("orm: Release date")
	var releaseSlice = make([]time.Time, 0, len(t))
	for _, v := range t {
		releaseSlice = append(releaseSlice,
			time.Date(v.ReleaseDate.Year(), v.ReleaseDate.Month(), v.ReleaseDate.Day(), 0, 0, 0, 0, time.UTC))
	}
	return releaseSlice
}
func handleError(err error) error {
	switch err {
	case nil:
		return err
	case gorm.ErrDuplicatedKey:
		return datastorage.ErrDuplicate
	case gorm.ErrForeignKeyViolated:
		return datastorage.ErrViolatesForeignKey
	case gorm.ErrRecordNotFound:
		return datastorage.ErrNotFound
	default:
		return errors.Join(datastorage.ErrDatabase, err)
	}
}
