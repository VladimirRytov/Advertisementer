package presenter

type ClientDTO struct {
	Name                  string
	Phones                string
	Email                 string
	AdditionalInformation string
}

type OrderDTO struct {
	ID            int
	ClientName    string
	Cost          string
	PaymentType   string
	CreatedDate   string
	PaymentStatus bool
}

type Advertisement struct {
	ID             int
	OrderID        int
	ReleaseCount   int
	ClosestRelease string
	ReleaseDates   string
	Cost           string
	Text           string
	Tags           string
	ExtraCharge    string
}

type BlockAdvertisementDTO struct {
	Advertisement
	Size     int
	FileName string
}

type LineAdvertisementDTO struct {
	Advertisement
}

type TagDTO struct {
	TagName string
	TagCost string
}

type ExtraChargeDTO struct {
	ChargeName string
	Multiplier string
}

type SelectedTagDTO struct {
	Selected bool
	TagName  string
}

type SelectedExtraChargeDTO struct {
	Selected   bool
	ChargeName string
}

type LastDatabaseConnection struct {
	DatabaseType string
	ConfigSaved  bool
	AutoLogin    bool
}

type LocalDSN struct {
	Name string
	Path string
	Type string
}

type NetworkDataBaseDSN struct {
	DatabaseName string
	Source       string
	DataBase     string
	UserName     string
	Password     string
	SSLMode      bool
	Port         string
}

type ServerDSN struct {
	DatabaseName string
	Source       string
	UserName     string
	Password     string
	SSLMode      bool
	Port         string
}

type CostRateDTO struct {
	Name            string
	OneWordOrSymbol string
	Onecm2          string
	CalcForOneWord  bool
}

type EnabledJournal struct {
	Journal string
}

type ShowData struct {
	Actual bool
}

type ImportParams struct {
	ActualClients bool

	AllBlocks       bool
	AlllLines       bool
	AllTags         bool
	AllExtraCharges bool
	AllCostRates    bool
	IgnoreErrors    bool
	ThickMode       bool
}

type ReportParams struct {
	ReportType       string
	FromDate, ToDate string
	BlocksFolderPath string
	DeployPath       string
}

type File struct {
	Name string
	Size string
	Data []byte
}

type ConnectionInfo struct {
	Addres string
	Path   string
	Token  string
}
