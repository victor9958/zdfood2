package routers

import (
	"zdfood2/controllers"
	"github.com/astaxie/beego"
	"zdfood2/controllers/order"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/v1/admin/order/index", &order.OrderController{},"get:Index")
    beego.Router("/getcode", &controllers.MainController{},"get:GetCode")
}
