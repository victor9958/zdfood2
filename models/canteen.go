package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Canteen struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	CampusId int     		`xorm:"not null INT(11)" json:"campus_id"`
	Code string     		`xorm:"not null VARCHAR(50)" json:"code"`
	Name string   				`xorm:"not null VARCHAR(50)" json:"name"`
	Address string   		`xorm:"not null VARCHAR(255)" json:"address"`
	Phone string 			`xorm:"not null CHAR(20)" json:"phone"`
	Description string        	`xorm:"not null VARCHAR(200)" json:"description"`
	SubDays string			`xorm:"not null VARCHAR(255)" json:"sub_days"`
	NoSubDate string		`xorm:"not null VARCHAR(255)" json:"no_sub_date"`
	NoSubTime string		`xorm:"not null VARCHAR(255)" json:"no_sub_time"`
	OrderExpire int			`xorm:"not null INT(11)" json:"order_expire"`
	CantSubMinute int		`xorm:"not null INT(11)" json:"cant_sub_minute"`
	Latitude string			`xorm:"not null VARCHAR(255)" json:"latitude"`
	Longitude int		 	`xorm:"not null VARCHAR(255)" json:"longitude"`
	Email int				`xorm:"not null VARCHAR(11)" json:"email"`
	Origin string			`xorm:"not null TINYINT(1)" json:"origin"`
	LunchStartAt string		`xorm:"not null VARCHAR(29)" json:"lunch_start_at"`
	LunchEndAt string		`xorm:"not null VARCHAR(11)" json:"lunch_end_at"`
	DinnerStartAt string	`xorm:"not null VARCHAR(11)" json:"dinner_start_at"`
	DinnerEndAt string			`xorm:"not null VARCHAR(11)" json:"dinner_end_at"`
	Accounts string			`xorm:"not null VARCHAR(50)" json:"accounts"`
	ServerTimeField string	`xorm:"not null VARCHAR(255)" json:"server_time_field"`
	Sort int				`xorm:"not null INT(11)" json:"sort"`
	ImageUrl string			`xorm:"not null VARCHAR(255)" json:"image_url"`
	EatType int				`xorm:"not null TINYINT(255)" json:"eat_type"`
	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`

}