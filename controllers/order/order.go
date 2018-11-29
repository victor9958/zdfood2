package order

import (
	"github.com/astaxie/beego"
	"zdfood2/controllers"
	"strconv"
	"zdfood2/models"
)

type OrderController struct {
	controllers.BaseController
}

func (this *OrderController)Index(){

	engine := models.Engine

	//condition := map[string]interface{}{}
	//获取参数
	if campusId :=this.GetString("campus_id");campusId!=""{
		beego.Info(campusId)
		id ,_:= strconv.Atoi(campusId)

		engine.Where("order.campus_id=?",id)

	}
	//campusId := this.GetString("campus_id")
	//if campusId==""{
	//	//this.Ctx.WriteString("校区id为空")
	//	campusId = "0"
	//}

	//id,err := strconv.Atoi(campusId)
	//if err != nil {
	//	this.ReturnJson(map[string]interface{}{"message":"校区id转换换错误"},400)
	//	//this.Ctx.WriteString("校区id装换错误")
	//}
	//nameOrMobile := this.GetString("name_or_mobile")

	var orders []models.Order
	//err2 := models.Engine.Where("order.campus_id=?",id).Limit(10).Find(&orders)
	err2 := engine.Limit(10).Find(&orders)

	if err2 != nil {
		//this.Ctx.WriteString("查询出错"+err2.Error())
		this.ReturnJson(map[string]interface{}{"message":"查询出错"+err2.Error()},400)
		//this.Ctx.WriteString(err2.Error())
	}



	//this.Data["json"] = map[string]interface{}{"data":&orders}
	//this.ServeJSON()
	this.ReturnJson(map[string]interface{}{"data":orders},200)


	//c.Data["json"] = order
	//c.ServeJSON()


}


func(this *OrderController)IndexConditon(){

}