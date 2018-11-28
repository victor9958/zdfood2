package models
import (
	_ "github.com/go-sql-driver/mysql"
)


type Log struct {
	Id int  `xorm:"not null pk autoincr INT(11)"`
	Value string `xorm:"not null VARCHAR(255)"`
	Key string `xorm:"not null VARCHAR(255)"`
}