package models
import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Area struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	AdminId int			`xorm:"not null INT(11)" json:"admin_id"`
	Name string		`xorm:"not null VARCHAR(255)" json:"name"`
	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`

}

type AreaData struct {
	Area
	Data []interface{} `json:"data"`
}

