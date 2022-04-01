package result

type SystemInfo struct {
	Base *Base `json:"base"`
	File *File `json:"file"`
	Pay  *Pay  `json:"pay"`
}
type Base struct {
	Title          string `json:"title"`
	ChildTitle     string `json:"childTitle"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Logo           string `json:"logo"`
	Icon           string `json:"icon"`
	CurrencySymbol string `json:"currencySymbol"`
	Language       string `json:"language"`
	RecordNumber   string `json:"recordNumber"`
}
type File struct {
	FileSize  string   `json:"fileSize"`
	ImageType []string `json:"imageType"`
	AudioType []string `json:"audioType"`
	VideoType []string `json:"videoType"`
}
type Pay struct {
	Recharge           []int   `json:"recharge"`
	PayMode            []int   `json:"payMode"`
	CashMin            float64 `json:"cashMin"`
	CashServicePercent float64 `json:"cashServicePercent"`
}
type Auth struct {
	RegisterMode string   `json:"registerMode"`
	PolicyUrl    string   `json:"policyUrl"`
	ProtocolUrl  string   `json:"protocolUrl"`
	Social       []string `json:"social"`
}

type SystemVipAndGradeList struct {
	Grade []*Grade `json:"grade"`
	Vip   []*Vip   `json:"vip"`
}
type Grade struct {
	GradeId     string   `json:"id"`
	Title       string   `json:"title"`       // 角色名称
	Icon        string   `json:"icon"`        // 角色图标
	Integral    int      `json:"integral"`    // 角色图标
	CreateGroup int      `json:"createGroup"` // 角色图标
	PostsModule []string `json:"postsModule"` // 角色图标
	CommonAuth  []string `json:"commonAuth"`  // 角色图标
}
type Vip struct {
	VipId    string  `json:"id"`
	Title    string  `json:"title"`
	Icon     string  `json:"icon"`
	Color    string  `json:"color"`
	Day      int     `json:"day"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
}
