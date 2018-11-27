package routers

import (
	"zdfood2/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/v1/admin/order/index", &controllers.MainController{},"get:Index")
}
