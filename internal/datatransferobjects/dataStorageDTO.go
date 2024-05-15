package datatransferobjects

import "io"

type LocalDSN struct {
	Path string `json:"filepath"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type NetworkDataBaseDSN struct {
	Source   string
	DataBase string
	UserName string
	Password string
	SSLMode  bool
	Port     uint
}

type ServerDSN struct {
	Source   string
	UserName string
	Password string
	SSLMode  bool
	Port     uint
}

type Files struct {
	FileNames []string `json:"fileNames"`
}

type FileDTO struct {
	Name string `json:"name"`
	Size int64  `json:"size,omitempty"`
	Data []byte `json:"data,omitempty"`
}

type FileStream struct {
	Name string
	Size int64
	Data io.ReadWriteCloser
}

type ResponceMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	AccessToken string `json:"accessToken"`
}

type BlockAdvertisementReport struct {
	Block BlockAdvertisementDTO
	Paid  bool
}

type LineAdvertisementReport struct {
	Line LineAdvertisementDTO
	Paid bool
}
