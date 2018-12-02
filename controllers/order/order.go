package order

import (
	"github.com/astaxie/beego"
	"strconv"
	"zdfood2/controllers"
	"zdfood2/models"
)

type OrderController struct {
	controllers.BaseController
}

func (this *OrderController)Index(){
	var arr  []interface{}
	var conditionStr []string

	engine := models.Engine.Limit(10)

	if campusId :=this.GetString("campus_id");campusId!=""{
		beego.Info(campusId)
		id ,_:= strconv.Atoi(campusId)
		conditionStr = append(conditionStr,"campus_id = ?")
		arr = append(arr,id)
	}
	if canteenId :=this.GetString("canteen_id");canteenId!=""{
		beego.Info(canteenId)
		id ,_:= strconv.Atoi(canteenId)
		conditionStr = append(conditionStr,"canteen_id = ?")
		arr = append(arr,id)
	}
	if eatType :=this.GetString("eat_type");eatType!=""{
		beego.Info(eatType)
		eatType ,_:= strconv.Atoi(eatType)
		//condition["order.eat_type"] = eatType
		conditionStr = append(conditionStr,"eat_type = ?")
		arr = append(arr,eatType)
	}
	if payType :=this.GetString("pay_type");payType!=""{
		beego.Info(payType)
		payType ,_:= strconv.Atoi(payType)
		//condition["order.pay_type"] = payType
		conditionStr = append(conditionStr,"pay_type = ?")
		arr = append(arr,payType)
	}

	if status :=this.GetString("status");status!=""{
		beego.Info(status)
		status ,_:= strconv.Atoi(status)
		//condition["order.status"] = status
		conditionStr = append(conditionStr,"status = ?")
		arr = append(arr,status)
	}

	if signUnitId :=this.GetString("sign_unit_id");signUnitId!=""{
		beego.Info(signUnitId)
		signUnitId ,_:= strconv.Atoi(signUnitId)
		//condition["order.sign_unit_id"] = signUnitId
		conditionStr = append(conditionStr,"sign_unit_id = ?")
		arr = append(arr,signUnitId)
	}

	if payStatus :=this.GetString("pay_status");payStatus!=""{
		beego.Info(payStatus)
		payStatus ,_:= strconv.Atoi(payStatus)
		//condition["order.pay_status"] = payStatus
		conditionStr = append(conditionStr,"pay_status = ?")
		arr = append(arr,payStatus)
	}

	if repastDate :=this.GetString("repast_date");repastDate!=""{
		beego.Info(repastDate)
		//condition["order.repast_date"] = repastDate
		conditionStr = append(conditionStr,"repast_date = ?")
		arr = append(arr,repastDate)

	}
	if areaId :=this.GetString("area_id");areaId!=""{
		beego.Info(areaId)
		areaId ,_:= strconv.Atoi(areaId)
		//condition["order.area_id"] = areaId
		conditionStr = append(conditionStr,"area_id = ?")
		arr = append(arr,areaId)
	}
	if nameOrMobile :=this.GetString("name_or_mobile");nameOrMobile!=""{
		beego.Info(nameOrMobile)
		//condition["order.user_name like"] = nameOrMobile
		nameOrMobileStr := "%"+nameOrMobile+"%"
		engine.Where("user_name like ? OR user_mobile like ?",nameOrMobileStr,nameOrMobileStr)
	}
	startTime :=this.GetString("start_time")
	endTime:=this.GetString("end_time")
	if startTime!="" && endTime!=""{
		engine.Where("repast_date BETWEEN ? AND ?",startTime ,endTime)
	}
	if conditionNum := len(conditionStr); conditionNum !=0 {
		conditionString := this.DisposeConditionStr(conditionStr...)
		engine.Where(conditionString,arr...)
	}
	engineC := *engine
	var orders []models.Order
	var order models.Order
	count,_ := engineC.Count(order)
	err2 := engine.Find(&orders)
	//
	if err2 != nil {
		//this.Ctx.WriteString("查询出错"+err2.Error())
		this.ReturnJson(map[string]interface{}{"message":"查询出错"+err2.Error()},400)
		//this.Ctx.WriteString(err2.Error())
	}
	//


	//this.Data["json"] = map[string]interface{}{"data":&orders}
	//this.ServeJSON()
	this.ReturnJson(map[string]interface{}{"data":orders,"count":count},200)


	//c.Data["json"] = order
	//c.ServeJSON()



}


func(this *OrderController)Ceshi(){


	//var orders []models.Order
	//engine :=models.Engine.Where("id = ?",542)
	//engine.Where("repast_date between ? AND ?","2018-11-29","2018-11-30")
	//engine2 := *engine
	//var order models.Order
	//count,_:=engine2.Count(order)
	//err2 := engine.Limit(10).Find(&orders)
	//if err2 != nil {
	//	this.ReturnJson(map[string]interface{}{"err2":err2,},400)
	//}
	//this.ReturnJson(map[string]interface{}{"data":orders,"count":count},200)

	var canteen []models.Canteen
	err := models.Engine.Select("id,name").Limit(10).Find(&canteen)
	canteenData := make(map[int]string)
	canteenData[78] = "vcitor"
	for _,v := range canteen{
		beego.Info(v.Id)
		canteenData[int(v.Id)] = v.Name
	}

	if err != nil {
		this.ReturnJson(err,200)
	}
	//beego.Info(canteen)
	this.ReturnJson(map[string]interface{}{"data":canteenData},200)
}

type Name struct {
	Id int
	Name string
}



func(this *OrderController)Ceshi2(){
	var arr  []interface{}
	var conditionStr []string

	engine := models.Engine

	//condition := map[string]interface{}{}
	//condition := map[string]interface{}{}


	if campusId :=this.GetString("campus_id");campusId!=""{
		beego.Info(campusId)
		id ,_:= strconv.Atoi(campusId)
		conditionStr = append(conditionStr,"campus_id = ?")
		arr = append(arr,id)
	}
	if canteenId :=this.GetString("canteen_id");canteenId!=""{
		beego.Info(canteenId)
		id ,_:= strconv.Atoi(canteenId)
		conditionStr = append(conditionStr,"campus_id = ?")
		arr = append(arr,id)
	}





	//获取参数
	//if campusId :=this.GetString("campus_id");campusId!=""{
	//	beego.Info(campusId)
	//	id ,_:= strconv.Atoi(campusId)
	//	condition["order.campus_id"] = id
	//}
	//if canteenId :=this.GetString("canteen_id");canteenId!=""{
	//	beego.Info(canteenId)
	//	id ,_:= strconv.Atoi(canteenId)
	//	condition["order.canteen_id"] = id
	//}
	if eatType :=this.GetString("eat_type");eatType!=""{
		beego.Info(eatType)
		eatType ,_:= strconv.Atoi(eatType)
		//condition["order.eat_type"] = eatType
		conditionStr = append(conditionStr,"eat_type = ?")
		arr = append(arr,eatType)
	}
	if payType :=this.GetString("pay_type");payType!=""{
		beego.Info(payType)
		payType ,_:= strconv.Atoi(payType)
		//condition["order.pay_type"] = payType
		conditionStr = append(conditionStr,"pay_type = ?")
		arr = append(arr,payType)
	}

	if status :=this.GetString("status");status!=""{
		beego.Info(status)
		status ,_:= strconv.Atoi(status)
		//condition["order.status"] = status
		conditionStr = append(conditionStr,"status = ?")
		arr = append(arr,status)
	}

	if signUnitId :=this.GetString("sign_unit_id");signUnitId!=""{
		beego.Info(signUnitId)
		signUnitId ,_:= strconv.Atoi(signUnitId)
		//condition["order.sign_unit_id"] = signUnitId
		conditionStr = append(conditionStr,"sign_unit_id = ?")
		arr = append(arr,signUnitId)
	}

	if payStatus :=this.GetString("pay_status");payStatus!=""{
		beego.Info(payStatus)
		payStatus ,_:= strconv.Atoi(payStatus)
		//condition["order.pay_status"] = payStatus
		conditionStr = append(conditionStr,"pay_status = ?")
		arr = append(arr,payStatus)
	}

	if repastDate :=this.GetString("repast_date");repastDate!=""{
		beego.Info(repastDate)
		//condition["order.repast_date"] = repastDate
		conditionStr = append(conditionStr,"repast_date = ?")
		arr = append(arr,repastDate)

	}
	if areaId :=this.GetString("area_id");areaId!=""{
		beego.Info(areaId)
		areaId ,_:= strconv.Atoi(areaId)
		//condition["order.area_id"] = areaId
		conditionStr = append(conditionStr,"area_id = ?")
		arr = append(arr,areaId)
	}
	if nameOrMobile :=this.GetString("name_or_mobile");nameOrMobile!=""{
		beego.Info(nameOrMobile)
		//condition["order.user_name like"] = nameOrMobile
		conditionStr = append(conditionStr,"user_name like ?")
		arr = append(arr,"%"+nameOrMobile+"%")

	}
	//nameOrMobile := "gogoc"
	startTime :=this.GetString("start_time");
	endTime:=this.GetString("end_time");
	if startTime!="" && endTime!=""{
		//condition["order.area_id"] = areaId
	}
	conditionString := this.DisposeConditionStr(conditionStr...)


	engineC := engine


	var orders []models.Order
	var order models.Order
	count,_ := engineC.Where(conditionString,arr...).Count(order)
	//err2 := models.Engine.Where("order.campus_id=?",id).Limit(10).Find(&orders)

	err2 := engine.Where(conditionString,arr...).Limit(10).Find(&orders)
	//
	if err2 != nil {
		//this.Ctx.WriteString("查询出错"+err2.Error())
		this.ReturnJson(map[string]interface{}{"message":"查询出错"+err2.Error()},400)
		//this.Ctx.WriteString(err2.Error())
	}
	//


	//this.Data["json"] = map[string]interface{}{"data":&orders}
	//this.ServeJSON()
	this.ReturnJson(map[string]interface{}{"data":orders,"count":count},200)


	//c.Data["json"] = order
	//c.ServeJSON()
}