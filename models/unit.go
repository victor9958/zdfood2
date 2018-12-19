package models
import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Unit struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	Name string     		`xorm:"not null VARCHAR(50)" json:"name"`
	ContactName string     		`xorm:"not null VARCHAR(50)" json:"contact_name"`
	ContactPhone string   		`xorm:"not null VARCHAR(20)" json:"contact_phone"`
	Remark string		`xorm:"not null VARCHAR(255)" json:"remark"`
	MasterUserId int			`xorm:"not null INT(11)" json:"master_user_id"`
	Summary string			`xorm:"not null VARCHAR(50)" json:"summary"`


	CampusId int		 	`xorm:"not null INT(11)" json:"campus_id"`

	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`

}

