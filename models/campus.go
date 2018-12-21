package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
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

func PluckCampusName(x *xorm.Session)(map[int]interface{},error){
	campus := map[int64]Campus{}
	err:= x.Select("id,name").Find(&campus)
	if err != nil{
		return map[int]interface{}{},err
	}
	res := map[int]interface{}{}
	for _,v := range campus{
		res[v.Id] = v.Name
	}
	res[0] = "未知"
	return res,nil
}