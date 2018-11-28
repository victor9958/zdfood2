package order

import (
	"zdfood2/controllers"
	"strconv"
	"zdfood2/models"
)

type OrderController struct {
	controllers.BaseController
}

func (this *OrderController)Index(){
	//获取参数
	campusId := this.GetString("campus_id")
	if campusId==""{
		this.Ctx.WriteString("校区id为空")
		return
	}
	id,err := strconv.Atoi(campusId)
	if err != nil {
		this.Ctx.WriteString("校区id装换错误")
		return
	}
	var orders []models.Order
	err2 := models.Engine.Where("order.campus_id=?",id).Find(&orders)

	if err2 != nil {
		this.Ctx.WriteString("查询出错"+err2.Error())
		//this.Ctx.WriteString(err2.Error())
		return
	}



	this.Data["json"] = map[string]interface{}{"data":&orders}
	this.ServeJSON()


	//c.Data["json"] = order
	//c.ServeJSON()


}