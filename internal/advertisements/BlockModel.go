package advertisements

import (
	"errors"
	"strings"
)

var (
	errSize     = errors.New("площадь должна быть больше 0")
	errFileName = errors.New("имя файла содержит недопустимые символы")
)

func ErrSize() error { return errSize }

func ErrFileName() error { return errFileName }

type AdvertisementBlock struct {
	size     int16
	comment  string
	fileName string
	advertisement
}

// Блочные объявления
func NewAdvertisementBlock() AdvertisementBlock {
	advertisementBlock := AdvertisementBlock{advertisement: newAdvertisement()}
	return advertisementBlock
}

func (ab *AdvertisementBlock) Size() int16 {
	return ab.size
}

func (ab *AdvertisementBlock) SetSize(s int16) error {
	if s == 0 {
		return errSize
	}
	ab.size = s
	return nil
}

func (ab *AdvertisementBlock) SetComment(c string) {
	ab.comment = c
}

func (ab *AdvertisementBlock) Comment() string {
	return ab.comment
}

func (ab *AdvertisementBlock) FileName() string {
	return ab.fileName
}

func (ab *AdvertisementBlock) SetFileName(c string) error {
	if c == `..` || c == `.` || strings.Contains(c, "/") || strings.Contains(c, `\`) {
		return errFileName
	}
	ab.fileName = c
	return nil
}
