package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Campus struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	Code string     		`xorm:"not null VARCHAR(50)" json:"code"`
	Name string   				`xorm:"not null VARCHAR(50)" json:"name"`
	Address string   		`xorm:"not null VARCHAR(255)" json:"address"`
	Latitude string			`xorm:"not null VARCHAR(255)" json:"latitude"`
	Longitude int		 	`xorm:"not null VARCHAR(255)" json:"longitude"`
	Origin string			`xorm:"not null TINYINT(1)" json:"origin"`
	Sort int				`xorm:"not null INT(11)" json:"sort"`
	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`

}