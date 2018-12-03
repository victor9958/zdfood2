package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Building struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	CampusId int     		`xorm:"not null INT(11)" json:"campus_id"`
	AreaId int     			`xorm:"not null INT(11)" json:"area_id"`
	Name string   			`xorm:"not null VARCHAR(50)" json:"name"`
	Code string   			`xorm:"not null VARCHAR(255)" json:"code"`
	Floor string 			`xorm:"not null CHAR(20)" json:"floor"`
	Origin string			`xorm:"not null TINYINT(1)" json:"origin"`
	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`
}