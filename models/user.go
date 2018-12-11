package models
import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type User struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	Uuid string     		`xorm:"not null VARCHAR(32)" json:"uuid"`
	Name string     		`xorm:"not null VARCHAR(50)" json:"name"`
	Email string   		`xorm:"not null VARCHAR(50)" json:"email"`
	Mobile string   		`xorm:"not null VARCHAR(20)" json:"mobile"`
	ZdMobile string   		`xorm:"not null VARCHAR(20)" json:"zd_mobile"`
	Password string 		`xorm:"not null CHAR(32)" json:"password"`
	StudentNum string        		`xorm:"not null CHAR(50)" json:"student_num"`
	Remark string		`xorm:"not null VARCHAR(255)" json:"remark"`
	Nickname string		`xorm:"not null CHAR(100)" json:"nickname"`
	CardNo string			`xorm:"not null VARCHAR(255)" json:"card_no"`
	Sort int		 	`xorm:"not null INT(11)" json:"sort"`
	Sex int			`xorm:"not null TINYINT(1)" json:"sex"`
	Openid string			`xorm:"not null VARCHAR(255)" json:"openid"`
	Role int			`xorm:"not null INT(11)" json:"role"`
	Headimgurl string				`xorm:"not null VARCHAR(500)" json:"headimgurl"`
	UnitId int			`xorm:"not null INT(11)" json:"unit_id"`
	IsAuth int			`xorm:"not null TINYINT(11)" json:"is_auth"`
	Position string			`xorm:"not null VARCHAR(50)" json:"position"`
	Origin string		`xorm:"not null VARCHAR(255)" json:"origin"`
	Address string		`xorm:"not null VARCHAR(255)" json:"address"`
	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`
	Flag int			`xorm:"TINYINT(1)" json:"flag"`

}

