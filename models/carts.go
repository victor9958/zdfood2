package models
import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Carts struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	CanteenId int     		`xorm:"not null INT(11)" json:"canteen_id"`
	UserId int        		`xorm:"not null INT(11)" json:"user_id"`
	GoodsId int        		`xorm:"not null INT(11)" json:"goods_id"`
	OrderId int        		`xorm:"not null INT(11)" json:"order_id"`
	Num int        		`xorm:"not null INT(11)" json:"num"`
	StyleId int        		`xorm:"not null INT(11)" json:"style_id"`
	StyleName string		`xorm:"not null VARCHAR(255)" json:"style_name"`
	StyleImage string		`xorm:"not null VARCHAR(255)" json:"style_image"`
	GoodsName string		`xorm:"not null VARCHAR(255)" json:"goods_name"`
	FinalPrice string	`xorm:"not null DECIMAL(10,2)" json:"final_price"`

	AreaId int				`xorm:"not null INT(11)" json:"area_id"`
	Status int				`xorm:"not null INT(11)" json:"status"`
	TakeOutType string	`xorm:"not null VARCHAR(50)" json:"take_out_type"`

	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`
	EatType int			`xorm:"TINYINT(1)" json:"eat_type"`

}

//type OrderNew struct {
//	Order
//	CanteenName 	string `json:"canteen_name"`
//	CampusName 		string `json:"campus_name"`
//	BuildingName 	string `json:"building_name"`
//	EatTypeName 	string `json:"eat_type_name"`
//	MealTypeName 	string `json:"meal_type_name"`
//	PayStatusName 	string `json:"pay_status_name"`
//	StatusName 		string `json:"status_name"`
//	PayTypeName 	string `json:"pay_type_name"`
//}

//type OrderGoods struct {
//
//}
//var EatType map[int]string = map[int]string{
//	0:"未知",
//	1:"外卖",
//	2:"堂食",
//}
//var MealType map[int]string = map[int]string{
//	0:"未知",
//	1:"早餐",
//	2:"中餐",
//	3:"晚餐",
//}
//var	PayStatus  map[int]string = map[int]string{
//	0:"未知",
//	1:"未支付",
//	2:"已支付",
//	3:"支付失败",
//}
//
//var	Status  map[int]string = map[int]string{
//	0:"未知",
//	1:"菜品生产中",
//	2:"待核销",
//	3:"已过期",
//	4:"已核销",
//	5:"带配送",
//	6:"配送中",
//	7:"配送完成",
//	8:"已取消",
//	//1菜品生产中  2待核销 3已过期  4已核销 5带配送 6配送中 7配送完成 8已取消
//}
//var	PayType  map[int]string = map[int]string{
//	0:"未知",
//	1:"一卡通",
//	2:"签单",
//}
