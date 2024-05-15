package orm

import "time"

type Client struct {
	Name                  string `gorm:"primaryKey;size:100"`
	Phones                string
	Email                 string
	AdditionalInformation string
	Orders                []Order `gorm:"foreignKey:ClientName;constraint:OnDelete:CASCADE"`
}
type Order struct {
	ID                  int
	ClientName          string `gorm:"size:100"`
	Cost                int
	CreatedDate         time.Time
	PaymentType         string
	PaymentStatus       bool
	AdvertisementsLines []AdvertisementLine  `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	AdvertisementBlocks []AdvertisementBlock `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}
type Advertisement struct {
	ID           int
	OrderID      int
	ReleaseCount int16
	Cost         int
	Text         string
}

type AdvertisementBlock struct {
	Advertisement
	Size         int16
	FileName     string
	Tags         []Tag          `gorm:"many2many:advertisementsblock_tags;foreignKey:ID;References:Name;constraint:OnDelete:CASCADE"`
	ExtraCharges []ExtraCharge  `gorm:"many2many:advertisementsblock_extracharge;foreignKey:ID;References:Name;constraint:OnDelete:CASCADE"`
	ReleaseDates []ReleaseDates `gorm:"many2many:advertisementsblock_releasedates;foreignKey:ID;References:ReleaseDate;constraint:OnDelete:CASCADE"`
}
type AdvertisementLine struct {
	Advertisement
	Tags         []Tag          `gorm:"many2many:advertisementsline_tags;foreignKey:ID;References:Name;constraint:OnDelete:CASCADE"`
	ExtraCharges []ExtraCharge  `gorm:"many2many:advertisementsline_extracharge;foreignKey:ID;References:Name;constraint:OnDelete:CASCADE"`
	ReleaseDates []ReleaseDates `gorm:"many2many:advertisementsline_releasedates;foreignKey:ID;References:ReleaseDate;constraint:OnDelete:CASCADE"`
}

type ReleaseDates struct {
	ReleaseDate time.Time `gorm:"primaryKey"`
}

type Tag struct {
	Name string `gorm:"primaryKey;size:20"`
	Cost int
}

type ExtraCharge struct {
	Name       string `gorm:"primaryKey;size:20"`
	Multiplier int
}

type CostRate struct {
	Name            string `gorm:"primaryKey;size:20"`
	OneWordOrSymbol int
	OneSquare       int
	CalcForOneWord  bool
}

type User struct {
	Name        string `gorm:"primaryKey;size:20"`
	Password    []byte
	Permissions byte
}
