package routers

import (
	"zdfood2/controllers"
	"github.com/astaxie/beego"
    "zdfood2/controllers/finance"
    "zdfood2/controllers/order"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/v1/admin/order/index", &order.OrderController{},"get:Index")
    beego.Router("/v1/admin/order/detail", &order.OrderController{},"get:Detail")
    beego.Router("/getcode", &controllers.MainController{},"get:GetCode")
    beego.Router("/ceshi", &order.OrderController{},"get:Ceshi")
    beego.Router("/ceshi", &order.OrderController{},"post:Ceshi")
    beego.Router("/ceshi2", &order.OrderController{},"get:Ceshi2")
    beego.Router("/v1/admin/order/cancel", &order.OrderController{},"put:Cancel")
    beego.Router("/v1/admin/order/today", &order.OrderController{},"get:Today")
    beego.Router("/v1/admin/order/batch-sign", &order.OrderController{},"put:BatchSign")
    beego.Router("/v1/admin/order/create-order", &order.OrderController{},"post:CreateOrder")
    beego.Router("/v1/admin/order/sign-list", &order.OrderController{},"get:SignList")
    beego.Router("/v1/admin/order/order-all-execl", &order.OrderController{},"get:OrderExecl")
    beego.Router("/v1/admin/order/order-statistics", &finance.FinanceController{},"get:OrderStatistics")
}
