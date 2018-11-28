package models
import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Order struct {
	Id int             		`xorm:"not null pk autoincr INT(11)"`
	CampusId int     		`xorm:"not null INT(11)"`
	CanteenId int     		`xorm:"not null INT(11)"`
	SignUnitId int   		`xorm:"not null INT(11)"`
	SignName string   		`xorm:"not null VARCHAR(255)"`
	SignRemark string 		`xorm:"not null VARCHAR(255)"`
	UserId int        		`xorm:"not null INT(11)"`
	BuildId int		 	`xorm:"not null INT(11)"`
	UserMobile string		`xorm:"not null VARCHAR(255)"`
	UserName string		`xorm:"not null VARCHAR(255)"`
	OrderSn string			`xorm:"not null VARCHAR(255)"`
	OriginalPrice string	`xorm:"not null VARCHAR(255)"`
	DiscountPrice string	`xorm:"not null VARCHAR(255)"`
	MealType int		 	`xorm:"not null INT(11)"`
	AddressId int			`xorm:"not null INT(11)"`
	Address string			`xorm:"not null VARCHAR(255)"`
	SignAt string			`xorm:"not null VARCHAR(255)"`
	Floor int				`xorm:"not null INT(11)"`
	PayType int			`xorm:"not null INT(11)"`
	PayStatus int			`xorm:"not null INT(11)"`
	RiderId int			`xorm:"not null INT(11)"`
	RiderName string		`xorm:"not null VARCHAR(255)"`
	RiderMobile string		`xorm:"not null VARCHAR(255)"`
	AreaId int				`xorm:"not null INT(11)"`
	Status int				`xorm:"not null INT(11)"`
	TakeOutType string	`xorm:"not null VARCHAR(255)"`
	CodeUrl string			`xorm:"not null VARCHAR(255)"`
	PayAt time.Time			`xorm:"VARCHAR(255)"`
	OutId string				`xorm:"not null VARCHAR(255)"`
	CardNo string			`xorm:"not null VARCHAR(255)"`
	RepastDate string		`xorm:"VARCHAR(255)"`
	CreatedAt time.Time			`xorm:"created"`
	EatStartAt time.Time			`xorm:"VARCHAR(255)"`
	TakeAt time.Time				`xorm:"VARCHAR(255)"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)"`
	Count int				`xorm:"not null INT(10)"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)"`

}