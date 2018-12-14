package models
import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Admin struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	Uuid string     		`xorm:"not null VARCHAR(32)" json:"uuid"`
	Name string     		`xorm:"not null VARCHAR(50)" json:"name"`
	Email string   		`xorm:"not null VARCHAR(50)" json:"email"`
	Mobile string   		`xorm:"not null VARCHAR(20)" json:"mobile"`
	JobNo string   		`xorm:"not null CHAR(50)" json:"job_no"`
	Password string 		`xorm:"not null CHAR(32)" json:"password"`
	Remark string		`xorm:"not null VARCHAR(255)" json:"remark"`
	RoleId int			`xorm:"not null INT(11)" json:"role_id"`
	Sort int		 	`xorm:"not null INT(11)" json:"sort"`
	Sex int			`xorm:"not null TINYINT(1)" json:"sex"`
	Super int			`xorm:"not null TINYINT(1)" json:"super"`
	Status int			`xorm:"not null TINYINT(1)" json:"status"`
	Position string			`xorm:"not null VARCHAR(50)" json:"position"`


	CampusId int		 	`xorm:"not null INT(11)" json:"campus_id"`
	CanteenId int		 	`xorm:"not null INT(11)" json:"canteen_id"`

	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`

}

