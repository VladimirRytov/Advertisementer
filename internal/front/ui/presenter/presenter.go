package presenter

import (
	"strconv"
	"strings"
)

func (dc *DataConverter) YearMonthDayToString(year, month, day uint) string {
	var (
		m string
		d string
	)
	if month < 10 {
		m = "0" + strconv.Itoa(int(month))
	} else {
		m = strconv.Itoa(int(month))
	}
	if day < 10 {
		d = "0" + strconv.Itoa(int(day))
	} else {
		d = strconv.Itoa(int(day))
	}
	return d + "." + m + "." + strconv.Itoa(int(year))
}

func (dc *DataConverter) SelectedReleaseDatesToString(releaseDates []string) string {
	var strBuilder strings.Builder
	strBuilder.Grow(12 * len(releaseDates))
	for _, v := range releaseDates {
		strBuilder.WriteString(v + ", ")
	}
	return strings.TrimSuffix(strBuilder.String(), ", ")
}

func (dc *DataConverter) SelectedTagsToString(tags []SelectedTagDTO) string {
	var strBuilder strings.Builder
	for _, v := range tags {
		if v.Selected {
			strBuilder.WriteString(v.TagName + ", ")
		}
	}
	return strings.TrimSuffix(strBuilder.String(), ", ")
}

func (dc *DataConverter) SelectedTagsToStringOld(tags []SelectedTagDTO) string {
	var tagStr string
	for _, v := range tags {
		if v.Selected {
			tagStr += v.TagName + ", "
		}
	}
	return strings.TrimSuffix(tagStr, ", ")
}

func (dc *DataConverter) SelectedExtraChargeToString(charges []SelectedExtraChargeDTO) string {
	var strBuilder strings.Builder
	for _, v := range charges {
		if v.Selected {
			strBuilder.WriteString(v.ChargeName + ", ")
		}
	}
	return strings.TrimSuffix(strBuilder.String(), ", ")
}

func (dc *DataConverter) SelectedTagsToDTO(tags string) []SelectedTagDTO {
	tagsArr := strings.Split(tags, ", ")
	selectedTags := make([]SelectedTagDTO, 0, len(tagsArr))
	for _, v := range tagsArr {
		selectedTags = append(selectedTags, SelectedTagDTO{Selected: true, TagName: v})
	}
	return selectedTags
}

func (dc *DataConverter) SelectedExtraChargesToDTO(charges string) []SelectedExtraChargeDTO {
	chargesArr := strings.Split(charges, ", ")
	selectedCharges := make([]SelectedExtraChargeDTO, 0, len(chargesArr))
	for _, v := range chargesArr {
		selectedCharges = append(selectedCharges, SelectedExtraChargeDTO{Selected: true, ChargeName: v})
	}
	return selectedCharges
}
