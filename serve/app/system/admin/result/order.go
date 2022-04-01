package result

import "github.com/gogf/gf/os/gtime"

// 订单信息
type OrderList struct {
	OrderId      int64       `json:"id"`           // 订单编号
	NickName     string      `json:"nickName"`     //
	OrderNum     string      `json:"orderNum"`     // 订单编号
	OrderType    int         `json:"orderType"`    //
	OrderMoney   float64     `json:"orderMoney"`   //
	PaymentMoney float64     `json:"paymentMoney"` //
	PayMethod    int         `json:"payMethod"`    //
	Status       int         `json:"status"`       //
	CreateTime   *gtime.Time `json:"createTime"`   // 创建时间
	PayTime      *gtime.Time `json:"payTime"`      // 创建时间
}

// 订单信息
type OrderInfo struct {
	OrderNum         string      `json:"orderNum"`         //
	NickName         string      `json:"nickName"`         //
	AuthorName       string      `json:"authorName"`       //
	PayMethod        int         `json:"payMethod"`        //
	OrderMoney       float64     `json:"orderMoney"`       //
	DistrictMoney    float64     `json:"districtMoney"`    //
	AuthorMoney      float64     `json:"authorMoney"`      //
	ServiceMoney     float64     `json:"serviceMoney"`     //
	PaymentMoney     float64     `json:"paymentMoney"`     //
	OrderPoint       float64     `json:"orderPoint"`       //
	OrderType        int         `json:"orderType"`        //
	OrderMode        int         `json:"orderMode"`        //
	ShippingMoney    float64     `json:"shippingMoney"`    //
	ShippingAddress  string      `json:"shippingAddress"`  //
	ShippingCompName string      `json:"shippingCompName"` //
	ShippingPhone    string      `json:"shippingPhone"`    //
	ShippingName     string      `json:"shippingName"`     //
	ShippingTime     *gtime.Time `json:"shippingTime"`     //
	Title            string      `json:"title"`            //
	DetailModule     string      `json:"detailModule"`     //
	ReceiveTime      *gtime.Time `json:"receiveTime"`      //
	Status           int         `json:"status"`           //
	CreateTime       *gtime.Time `json:"createTime"`       // 创建时间
	PayTime          *gtime.Time `json:"payTime"`          // 创建时间
}
