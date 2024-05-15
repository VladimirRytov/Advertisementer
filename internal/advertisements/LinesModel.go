package advertisements

import (
	"strings"
	"unicode"
)

type AdvertisementLine struct {
	advertisement
	message string
}

func NewAdvertisementLine() AdvertisementLine {
	AdvertisementLine := AdvertisementLine{advertisement: newAdvertisement()}
	return AdvertisementLine
}

func (al *AdvertisementLine) SetContentOld(c string) {
	var newString []rune = make([]rune, 0, len(c))
	for _, v := range c {
		if v == rune('\n') {
			newString = append(newString, ' ')
			continue
		}
		newString = append(newString, v)
	}
	al.message = strings.TrimSpace(string(newString))
}

func (al *AdvertisementLine) SetContent(c string) {
	var (
		prevIsSpace bool
		newString   strings.Builder
	)
	newString.Grow(len(c))
	for _, v := range c {
		switch {
		case unicode.IsSpace(v):
			if prevIsSpace {
				continue
			}
			prevIsSpace = true
			newString.WriteRune(' ')
		default:
			prevIsSpace = false
			newString.WriteRune(v)
		}
	}
	al.message = strings.TrimSpace(newString.String())
}

func (al *AdvertisementLine) Content() string {
	return string(al.message)
}

func (al *AdvertisementLine) WordsCount() int {
	if len(al.message) == 0 {
		return 0
	}
	var words int = 1
	for _, v := range al.message {
		if v == ' ' {
			words++
		}
	}
	return words
}

func (al *AdvertisementLine) SymbolsCount() int {
	return len([]rune(al.message))
}
