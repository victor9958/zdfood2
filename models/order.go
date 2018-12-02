package models
import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Order struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	CampusId int     		`xorm:"not null INT(11)" json:"campus_id"`
	CanteenId int     		`xorm:"not null INT(11)" json:"canteen_id"`
	SignUnitId int   		`xorm:"not null INT(11)" json:"sign_unit_id"`
	SignName string   		`xorm:"not null VARCHAR(255)" json:"sign_name"`
	SignRemark string 		`xorm:"not null VARCHAR(255)" json:"sign_remark"`
	UserId int        		`xorm:"not null INT(11)" json:"user_id"`
	BuildId int		 	`xorm:"not null INT(11)" json:"build_id"`
	UserMobile string		`xorm:"not null VARCHAR(255)" json:"user_mobile"`
	UserName string		`xorm:"not null VARCHAR(255)" json:"user_name"`
	OrderSn string			`xorm:"not null VARCHAR(255)" json:"order_sn"`
	OriginalPrice string	`xorm:"not null VARCHAR(255)" json:"orginal_price"`
	DiscountPrice string	`xorm:"not null VARCHAR(255)" json:"discount_price"`
	MealType int		 	`xorm:"not null INT(11)" json:"meal_type"`
	AddressId int			`xorm:"not null INT(11)" json:"address_id"`
	Address string			`xorm:"not null VARCHAR(255)" json:"address"`
	SignAt string			`xorm:"not null VARCHAR(255)" json:"sign_at"`
	Floor int				`xorm:"not null INT(11)" json:"floor"`
	PayType int			`xorm:"not null INT(11)" json:"pay_type"`
	PayStatus int			`xorm:"not null INT(11)" json:"pay_status"`
	RiderId int			`xorm:"not null INT(11)" json:"rider_id"`
	RiderName string		`xorm:"not null VARCHAR(255)" json:"rider_name"`
	RiderMobile string		`xorm:"not null VARCHAR(255)" json:"rider_mobile"`
	AreaId int				`xorm:"not null INT(11)" json:"area_id"`
	Status int				`xorm:"not null INT(11)" json:"status"`
	TakeOutType string	`xorm:"not null VARCHAR(255)" json:"take_out_type"`
	CodeUrl string			`xorm:"not null VARCHAR(255)" json:"code_url"`
	PayAt time.Time			`xorm:"VARCHAR(255)" json:"pay_at"`
	OutId string				`xorm:"not null VARCHAR(255)" json:"out_id"`
	CardNo string			`xorm:"not null VARCHAR(255)" json:"card_no"`
	RepastDate string		`xorm:"VARCHAR(255)" json:"repast_date"`
	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	EatStartAt time.Time			`xorm:"VARCHAR(255)" json:"eat_start_at"`
	TakeAt time.Time				`xorm:"VARCHAR(255)" json:"take_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	Count int				`xorm:"not null INT(10)" json:"count"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`

}