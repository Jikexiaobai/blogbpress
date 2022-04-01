package result

type Vip struct {
	VipId    int64   `json:"id"`
	Title    string  `json:"title"`
	Icon     string  `json:"icon"`
	Day      int     `json:"day"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount"`
	Color    string  `json:"color"`
}
