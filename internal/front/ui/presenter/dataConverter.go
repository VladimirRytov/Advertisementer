package presenter

import (
	"errors"
	"math"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
)

type DataConverter struct {
	costRegex *regexp.Regexp
	dateRegex *regexp.Regexp
}

func NewDataConverter() *DataConverter {
	return &DataConverter{
		costRegex: regexp.MustCompile(`^([+-])?([0-9]+)[\.|\,]?([0-9]{1,2})?$`),
		dateRegex: regexp.MustCompile(`^([0-9]{2})\.([0-9]{2})\.([0-9]{4,})$`),
	}
}

func (dc *DataConverter) CostToView(cost int) string {
	logging.Logger.Debug("converter: Converting cost val to string")
	remainder := cost % 100
	if (remainder % 10) == remainder {
		return strconv.Itoa(cost/100) + "," + "0" + strconv.Itoa(remainder)
	}
	return strconv.Itoa(cost/100) + "," + strconv.Itoa(remainder)
}

func (dc *DataConverter) DateToView(date time.Time) string {
	logging.Logger.Debug("converter: Converting date val to string")
	var strBuilder strings.Builder
	strBuilder.Grow(12)
	if day := strconv.Itoa(date.Day()); len(day) == 1 {
		strBuilder.WriteString("0" + day + ".")
	} else {
		strBuilder.WriteString(day + ".")
	}

	if month := strconv.Itoa(int(date.Month())); len(month) == 1 {
		strBuilder.WriteString("0" + month + ".")
	} else {
		strBuilder.WriteString(month + ".")
	}
	strBuilder.WriteString(strconv.Itoa(date.Year()))

	return strBuilder.String()
}

func (dc *DataConverter) DateToDto(date string) (time.Time, error) {
	logging.Logger.Debug("converter: Converting date string to date val")
	match := dc.dateRegex.FindStringSubmatch(date)
	if len(match) != 4 {
		return time.Time{}, errors.New("шаблон даты не найден")
	}

	year, _ := strconv.Atoi(match[3])
	month, _ := strconv.Atoi(match[2])
	day, _ := strconv.Atoi(match[1])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func (dc *DataConverter) CostToDto(cost string) (int, error) {
	logging.Logger.Debug("converter: Converting cost string to cost val")
	match := dc.costRegex.FindStringSubmatch(cost)
	if len(match) == 0 {
		return 0, errors.New("шаблон стоимости не найден")
	}
	var sign = 1
	if match[1] == "-" {
		sign = -1
	}
	num, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, dc.handleError(err)
	}

	var remainder int
	if len(match[3]) != 0 {
		if len(match[3]) == 1 {
			match[3] += "0"
		}

		remainder, err = strconv.Atoi(match[3])
		if err != nil {
			return 0, dc.handleError(err)
		}
	}

	return (num*100 + remainder) * sign, nil
}

func (dc *DataConverter) ArrayToString(arr []string) string {
	var strBuilder strings.Builder

	for _, v := range arr {
		strBuilder.WriteString(v + ", ")
	}
	return strings.TrimSuffix(strBuilder.String(), ", ")
}

func (dc *DataConverter) ParsePath(path string) string {
	if runtime.GOOS == "windows" {
		after, found := strings.CutPrefix(path, "/")
		if found {
			return filepath.Clean(after)
		}
	}
	return filepath.Clean(path)
}

func (dc *DataConverter) calcFileSize(byteLen int) string {
	switch {
	case byteLen < KiB:
		return strconv.Itoa(byteLen) + " " + "Byte"
	case byteLen < MiB:
		i := dc.convertFileSize(byteLen, KiB)
		return i + " " + "Kib"
	case byteLen < GiB:
		i := dc.convertFileSize(byteLen, MiB)
		return i + " " + "Mib"
	}
	i := dc.convertFileSize(byteLen, GiB)

	return i + " " + "Gib"
}

func (dc *DataConverter) convertFileSize(byteLen int, size int) string {
	rou := math.Pow10(2)
	got := math.Round((float64(byteLen) / float64(size) * rou)) / rou
	return strconv.FormatFloat(got, 'f', -1, 64)
}

func (dc *DataConverter) convertFileSizeOld(byteLen int, size int) string {
	var su string
	if byteLen%100 != 0 {
		su = strconv.Itoa(byteLen / size % 100)
	}
	if len(su) == 0 {
		return strconv.Itoa(byteLen / (100 * size))
	}
	return strconv.Itoa(byteLen/(100*size)) + "," + su
}

func (dc *DataConverter) ClosestRelease(arr []time.Time, date time.Time) (time.Time, error) {
	if len(arr) == 0 {
		return date, errors.New("closestRelease: array is empty")
	}
	left := 0
	right := len(arr)
	if date.Before(arr[left]) {
		return arr[left], nil
	}

	for left < right {
		mid := (right + left) / 2

		if date.Equal(arr[mid]) {
			return arr[mid], nil
		}

		if date.After(arr[mid]) {
			if mid < len(arr)-1 {
				if date.Before(arr[mid+1]) {
					return arr[mid+1], nil
				}
			}
			left = mid + 1
			continue
		}
		if mid > 0 {
			if date.After(arr[mid-1]) {
				return arr[mid], nil
			}
		}
		right = mid
	}
	return date, errors.New("closestRelease: closest release not found")
}

func (dc *DataConverter) splitStr(str string) []string {
	if len(str) == 0 {
		return nil
	}
	return strings.Split(str, ", ")
}

func (dc *DataConverter) handleError(err error) error {
	switch {
	case errors.Is(err, strconv.ErrRange):
		return errors.New("введённое значение выходит за рамки доступных границ")
	case errors.Is(err, strconv.ErrSyntax):
		return errors.New("невозможно перевести введённое значение в числовой формат")
	}
	return err
}
