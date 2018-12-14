package models
import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type Goods struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	ZdDishId int             		`xorm:"INT(11)" json:"zd_dish_id"`
	Code string 		`xorm:"not null CHAR(50)" json:"code"`
	CanteenId int     		`xorm:"not null INT(11)" json:"canteen_id"`
	Name string		`xorm:"not null VARCHAR(255)" json:"name"`
	Describe string		`xorm:"not null VARCHAR(255)" json:"describe"`
	Stock int   		`xorm:"not null INT(11)" json:"stock"`
	Price string	`xorm:"not null VARCHAR(255)" json:"price"`
	Sales int 		`xorm:"not null INT(11)" json:"sales"`
	OrderLimit int        		`xorm:"not null INT(11)" json:"order_limit"`
	LessLimit int		 	`xorm:"not null INT(11)" json:"less_limit"`
	Type int		 	`xorm:"not null TINYINT(11)" json:"type"`
	MealType int		 	`xorm:"not null TINYINT(11)" json:"meal_type"`
	IsHot int		 	`xorm:"not null TINYINT(11)" json:"is_hot"`
	Origin string	`xorm:"not null VARCHAR(255)" json:"origin"`
	ImageUrl string		`xorm:"not null VARCHAR(255)" json:"image_url"`
	SaleDate string			`xorm:"not null VARCHAR(255)" json:"sale_date"`
	Detail string			`xorm:"not null VARCHAR(255)" json:"detail"`
	IsSpecialty int				`xorm:"not null TINYINT(1)" json:"is_specialty"`
	SpecialtySort int			`xorm:"not null INT(11)" json:"specialty_sort"`
	Sort int			`xorm:"not null INT(11)" json:"sort"`
	Material string			`xorm:"not null VARCHAR(255)"  json:"material"`
	Nutritive string			`xorm:"not null VARCHAR(255)" json:"nutritive"`//	`xorm:"not null VARCHAR(255)"
	CreatedAt time.Time			`xorm:"created" json:"created_at"`
	UpdatedAt time.Time			`xorm:"VARCHAR(255)" json:"updated_at"`
	DeletedAt time.Time			`xorm:"VARCHAR(255)" json:"deleted_at"`

}
type GoodsInput struct {
	GoodsId []int `json:"goods_id"`
	Num []int `json:"num"`
	
}
