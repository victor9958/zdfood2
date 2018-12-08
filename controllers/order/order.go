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
	var page int =0
	var err7 error
	if pageStr :=this.GetString("page");pageStr!=""{
		beego.Info(pageStr)
		page,err7 = strconv.Atoi(pageStr)
		if err7 !=nil {

		}
		beego.Info("strconv",page)
		page  = (page-1)*10

	}
	beego.Info(page)
	engine := models.Engine.Limit(10,page)

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
	count,_ := engineC.Count(&order)
	err2 := engine.Find(&orders)
	if err2 != nil {
		//this.Ctx.WriteString("查询出错"+err2.Error())
		this.ReturnJson(map[string]interface{}{"message":"查询出错"+err2.Error()},400)
		//this.Ctx.WriteString(err2.Error())
	}

	//其他数据
	var ordersNew []*models.OrderNew
	//食堂
	canteen ,err4 := models.Pluck(models.Engine.Table("canteen"),"name")
	canteen[0] = "未知"
	if err4 != nil{
		this.ReturnJson(err4,400)
	}
	//校区
	campus ,err5 := models.Pluck(models.Engine.Table("campus"),"name")
	campus[0] = "未知"
	if err5 != nil{
		this.ReturnJson(err5,400)
	}

	//楼幢
	building ,err6 := models.Pluck(models.Engine.Table("campus"),"name")
	building[0] = "未知"
	if err5 != nil{
		this.ReturnJson(err6,400)
	}

	for _,v:= range orders{
	//	ordern :=
		ordersNew = append(ordersNew,&(models.OrderNew{v,(canteen[v.CanteenId]).(string),campus[v.CampusId].(string),building[v.BuildId].(string),"",models.EatType[v.EatType],models.MealType[v.MealType],models.PayStatus[v.PayStatus],models.Status[v.Status],models.PayType[v.PayType]}))
	}


	//

	//


	//this.Data["json"] = map[string]interface{}{"data":&orders}
	//this.ServeJSON()
	this.ReturnJson(map[string]interface{}{"data":ordersNew,"count":count},200)


	//c.Data["json"] = order
	//c.ServeJSON()



}

/*
	订单详情
 */
func(this *OrderController)Detail(){
	//var goods []*models.Carts
	//err := models.Engine.Limit(10).Find(&goods)
	//if err != nil {
	//	this.ReturnJson(err,200)
	//}
	//this.ReturnJson(map[string]interface{}{"data":goods},200)
	var order models.Order
	var goods []*models.Carts
	var id int = 0
	//var err3 error
	/*************************连表查询语句******************************************/
	//err := models.Engine.Table("order").Join("INNER", "carts", "carts.order_id = order.id").Limit(10).Find(&orderGoods)
	/*************************连表查询语句******************************************/
	idStr := this.GetString("id")
	if idStr == "" {
		this.ReturnJson(map[string]string{"message":"请输入订单id"},400)
	}else{
		id,_ = strconv.Atoi(idStr)
	}

	res,err := models.Engine.Id(id).Get(&order)


	if err != nil || !res {
		if !res {
			this.ReturnJson(map[string]string{"message":"订单不存在"},400)
		}
		this.ReturnJson(map[string]string{"message":err.Error()},400)
	}

	//其他数据
	//食堂
	var canteenName string = ""
	if order.CanteenId != 0{
		num ,_ := models.Engine.Table("canteen").Id(order.CanteenId).Cols("name").Get(&canteenName)
		if !num  {
			this.ReturnJson(map[string]string{"message":"无此食堂"},400)
		}
	}
	//校区
	var campusName string = ""
	if order.CampusId != 0{
		num ,_ := models.Engine.Table("campus").Id(order.CampusId).Cols("name").Get(&campusName)
		if !num  {
			this.ReturnJson(map[string]string{"message":"无此校区"},400)
		}
	}


	//楼幢
	var buildingName string = ""
	if order.CampusId != 0{
		num ,_ := models.Engine.Table("building").Id(order.BuildId).Cols("name").Get(&buildingName)
		if !num  {
			this.ReturnJson(map[string]string{"message":"无此楼撞"},400)
		}
	}


	//区域
	var areaName string = ""
	if order.AreaId != 0{
		num ,_ := models.Engine.Table("area").Id(order.AreaId).Cols("name").Get(&areaName)
		if !num  {
			this.ReturnJson(map[string]string{"message":"无此区域"},400)
		}
	}






	//orderMap := controllers.StructToMap(order)
	//beego.Info(orderMap)

	err2 := models.Engine.Where("order_id = ?",id).Find(&goods)
	if err2 != nil {
		this.ReturnJson(map[string]string{"message":err2.Error()},400)
	}
	this.ReturnJson(map[string]interface{}{"data":models.OrderNew{order,canteenName,campusName,buildingName,areaName,models.EatType[order.EatType],models.MealType[order.MealType],models.PayStatus[order.PayStatus],models.Status[order.Status],models.PayType[order.PayType]},"goods":goods},200)
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

	//var canteen []models.Canteen
	//err := models.Engine.Select("id,name").Limit(10).Find(&canteen)
	//canteenData := make(map[int]string)
	//canteenData[78] = "vcitor"
	//for _,v := range canteen{
	//	beego.Info(v.Id)
	//	canteenData[int(v.Id)] = v.Name
	//}
	//var canteenData []models.Canteen
	canteenData := make(map[int]interface{})
	//var canteenData []string
	//err := models.Engine.Cols("id").Limit(10).Find(&canteenData)
	err := models.Engine.Table("canteen").Limit(10).Find(&canteenData)




	if err != nil {
		this.ReturnJson(err,200)
	}
	//engine := models.Engine.Limit(10)
	var names []string
	err3 := models.Engine.Table("canteen").Cols("id").Limit(10).Find(&names)
	if err3 != nil {
		this.ReturnJson(err3,200)
	}
	res := make(map[int]interface{})

	//for _,v := range names{
	//	id,_ := strconv.Atoi(v)
	//	res[id] = "sss"
	//	//beego.Info(v)
	//}


	res ,err4 := models.Pluck(models.Engine.Table("canteen"),"name")
	if err4 != nil{
		this.ReturnJson(err4,200)
	}


	this.ReturnJson(map[string]interface{}{"columns":names,"data":res},200)
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
/*
	取消订单
 */
 func(this *OrderController)Cancel(){
 	var err error
 	var orderId int
	 if orderIdStr := this.GetString("order_id");orderIdStr != "" {
	 	orderId,err = strconv.Atoi(orderIdStr)
		 if err != nil {
			 this.ReturnJson(map[string]string{"message":"请输入正确的订单id"},400)
		 }
		 //this.ReturnJson(map[string]string{"message":"请输入orderId"},200)
	 }
	 var order models.Order
	 res,err2 := models.Engine.Id(orderId).Get(&order)
	 if err2 != nil || !res{
		 this.ReturnJson(map[string]string{"message":"订单不存在"},400)
	 }
	 if order.Status !=1 {
		 this.ReturnJson(map[string]string{"message":"当前订单状态不可取消"},400)
	 }
	 if order.PayType !=1 {
		 this.ReturnJson(map[string]string{"message":"只有一卡通才能取消订单"},400)
	 }
	 if order.PayStatus !=1 {
		 this.ReturnJson(map[string]string{"message":"只有未支付订单才可取消订单"},400)
	 }


	 session := models.Engine.NewSession()
	 defer session.Close()

	 err3 := session.Begin()
	 if err3 !=  nil {
		 this.ReturnJson(map[string]string{"message":"只有未支付订单才可取消订单"},400)
	 }
	 _,err4 := session.Update(&order)
	 if err4 != nil {
	 	session.Rollback()
		 this.ReturnJson(map[string]string{"message":"修改订单信息失败"},400)
	 }
	 var goods []*models.Carts
	 err6 := models.Engine.Where("order_id = ?",order.Id).Find(&goods)
	 if err6!=nil {
		 this.ReturnJson(map[string]string{"message":"该订单的菜品信息错误"},400)
	 }
	 for k,_ := range goods{
		 goods[k].Status = 3
	 }

	 num,err7 := session.Update(&goods)
	 beego.Info(num)
	 if err7 != nil {
	 	session.Rollback()
		 this.ReturnJson(map[string]string{"message":"修改菜品信息错误"},400)
	 }

	 err8 := session.Commit()
	 if err8 != nil {
		 this.ReturnJson(map[string]string{"message":"事务提交失败"},400)
	 }

 }