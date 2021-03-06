package order

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/go-xorm/xorm"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"
	"zdfood2/controllers"
	"zdfood2/models"
)

type OrderController struct {
	controllers.BaseController
}
func(this *OrderController)condition(engine *xorm.Session){
	var arr  []interface{}
	var conditionStr []string
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
}

func (this *OrderController)Index(){

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

	this.condition(engine)

	engineC := *engine
	var orders []*models.Order
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
		ordersNew = append(ordersNew,&(models.OrderNew{v,(canteen[v.CanteenId]).(string),campus[v.CampusId].(string),building[v.BuildId].(string),"","",models.EatType[v.EatType],models.MealType[v.MealType],models.PayStatus[v.PayStatus],models.Status[v.Status],models.PayType[v.PayType]}))
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
	this.ReturnJson(map[string]interface{}{"data":models.OrderNew{&order,canteenName,campusName,buildingName,areaName,"",models.EatType[order.EatType],models.MealType[order.MealType],models.PayStatus[order.PayStatus],models.Status[order.Status],models.PayType[order.PayType]},"goods":goods},200)
}



type CeshiJson struct {
	GoodsIds []int
}

func(this *OrderController)ExcelHeader(path string){
	this.Ctx.Output.Header("Accept-Ranges", "bytes")
	this.Ctx.Output.Header("Content-Disposition", "attachment; filename="+fmt.Sprintf("%s", path))//文件名
	this.Ctx.Output.Header("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	this.Ctx.Output.Header("Pragma", "no-cache")
	this.Ctx.Output.Header("Expires", "0")
	http.ServeFile(this.Ctx.ResponseWriter, this.Ctx.Request, "Book.xlsx")
}

func(this *OrderController)Ceshi(){
	var name Name
	name.Id = 111
	name.Name = "vart"
	object := reflect.ValueOf(&name)
	myref := object.Elem()

	for i:=0;i<myref.NumField() ;i++{
		field := myref.Field(i)
		beego.Info(field)
		beego.Info(field.Interface())
	}

	this.ReturnJson("结束",200)


	//var orders []*models.OrderExcel
	////err := engineTmp.Where("campus_id = ?",user.CampusId).Limit(50,(i-1)*50).Find(&orders)
	//err := models.Engine.Table("order").Join("INNER","carts","carts.order_id = order.id").
	//	Select("`order`.*,`carts`.num,`carts`.goods_name,`carts`.final_price").
	//	Where("order.campus_id = ?",1).Find(&orders)
	//beego.Info(err)
	//this.ReturnJson(map[string]interface{}{"data":orders},200)




	//ex := excelize.NewFile()
	//ex.SetCellValue("sheet1","A1","sssssss")
	//buff,err11 := ex.WriteToBuffer()
	//if err11 != nil {
	//	this.ReturnJson(map[string]string{"message":"已过期请重新登录"},400)
	//}
	//
	////测试下载
	////this.Ctx.Output.Header("Accept-Ranges", "bytes")
	//this.Ctx.Output.Header("Content-Type", "application/octet-stream")
	//this.Ctx.Output.Header("Content-Disposition", "attachment; filename="+fmt.Sprintf("%s", "Book.xlsx"))//文件名
	//this.Ctx.Output.Header("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	//this.Ctx.Output.Header("Pragma", "no-cache")
	//this.Ctx.Output.Header("Expires", "0")
	//_,err12 := io.Copy(this.Ctx.ResponseWriter,buff)
	//if err12 != nil{
	//	this.ReturnJson(map[string]string{"message":"写入header出错"},400)
	//}
	//this.ReturnJson(map[string]string{"message":"写入完成"},200)
	//http.ServeFile(this.Ctx.ResponseWriter, this.Ctx.Request, "Book.xlsx")
	//this.ExcelHeader("Book.xlsx")
	//最主要的一句


	//t := strconv.Itoa(int(time.Now().Unix()))
	//w := map[string]string{"first":"您的订单已结账成功",
	//"keyword1":"keyword1",
	//"keyword2":"keyword2",
	//"keyword3":"keyword3",
	//"time":t,
	//"touser":"oW9idwtwDrOEE8w2ZgZPL8b9nNY0"}
	////oW9idwtwDrOEE8w2ZgZPL8b9nNY0
	//err := this.SendMessageTmplate2(w)
	//if err!= nil {
	//	this.ReturnJson(map[string]string{"message":err.Error()},400)
	//}
	//bb := "肯定及时飞机上的浪费"
	//bbb := strings.Replace(bb,"及时","aaa",-1)
	//bbb :=base64.StdEncoding.EncodeToString([]byte(bb))  //base64 编码
	//urlencode 是把 字符串转变成16进制然后 在每个字符前面加上 %
	//this.ReturnJson(map[string]string{"message":"成功"},400)


	//aa := "%E8%82%AF%E5%AE%9A%E5%B0%B1%E6%98%AF%E9%A3%9E%E6%9C%BA%E4%B8%8A%E7%9A%84%E6%B5%AA%E8%B4%B9"
	//bb := "肯定及时飞机上的浪费"
	//bbb := url.QueryEscape(bb)
	////bbb :=base64.StdEncoding.EncodeToString([]byte(bb))  //base64 编码
	////urlencode 是把 字符串转变成16进制然后 在每个字符前面加上 %
	//this.ReturnJson(map[string]string{"message":bbb,"2222":aa},400)

	//t := make(map[string]string)
	//
	//t["message"] = "这的说法是打开了房间"
	//str,err := url.Parse(t["message"])
	//if err != nil {
	//		this.ReturnJson(map[string]string{"message":err.Error()},400)
	//}
	//this.ReturnJson(map[string]string{"message":str.Query().Encode()},400)


	//var log models.Log
	//log.Value = "sss"
	//log.Key = "sssss"
	//_,err := models.Engine.Insert(&log)
	//if err != nil {
	//	this.ReturnJson(map[string]string{"message":err.Error()},400)
	//
	//}
	//this.ReturnJson(map[string]interface{}{"message":log},400)
	//this.ReturnJson(map[string]string{"message":time.Now().Format("2006-01-02 15:04:05")},400)
	//timeStr ,_:= time.Parse("2006-01-02 15:04:05","2018-11-13 19:00:00")
	//num := timeStr.Add(3600*12*1e9)
	//this.ReturnJson(map[string]interface{}{"message":num},400)

	//if  goodsIds := this.GetStrings("goods_ids");len(goodsIds)>0{
	//	beego.Info("goods_ids")
	//	this.ReturnJson(map[string]interface{}{"data":goodsIds},200)
	//}
	//
	//if goodsIds2:=this.GetStrings("goods_ids[]");len(goodsIds2)>0 {
	//	beego.Info("goods_ids[]")
	//	this.ReturnJson(map[string]interface{}{"message":goodsIds2},200)
	//}
	//var ob CeshiJson
	//json.Unmarshal(this.Ctx.Input.RequestBody,&ob)
	//for _,v := range
	//this.ReturnJson(map[string]interface{}{"json_data":ob},400)
	//
	//this.ReturnJson(map[string]string{"message":"没有进入任何选项"},400)

	//var users []*models.Carts
	//err := models.Engine.Where("order_id = ?",347).Find(&users)
	//beego.Info(users)
	//if err!= nil {
	//	this.ReturnJson(map[string]string{"message":err.Error()},400)
	//}
	//
	////
	//this.ReturnJson(map[string]interface{}{"data":users},200)

	//arr := this.GetStrings("aa")
	//
	//this.ReturnJson(map[string]interface{}{"columns":arr},200)

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
	//canteenData := make(map[int]interface{})
	//var canteenData []string
	//err := models.Engine.Cols("id").Limit(10).Find(&canteenData)
	//err := models.Engine.Table("canteen").Limit(10).Find(&canteenData)




	//if err != nil {
	//	this.ReturnJson(err,200)
	//}
	////engine := models.Engine.Limit(10)
	//var names []string
	//err3 := models.Engine.Table("canteen").Cols("id").Limit(10).Find(&names)
	//if err3 != nil {
	//	this.ReturnJson(err3,200)
	//}
	//res := make(map[int]interface{})

	//for _,v := range names{
	//	id,_ := strconv.Atoi(v)
	//	res[id] = "sss"
	//	//beego.Info(v)
	//}


	//res ,err4 := models.Pluck(models.Engine.Table("canteen"),"name")
	//if err4 != nil{
	//	this.ReturnJson(err4,200)
	//}


	//this.ReturnJson(map[string]interface{}{"columns":names,"data":res},200)
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
	 _,err4 := session.Id(order.Id).Update(&order)
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
	 for _,v := range goods{
	 	v.Status = 3
		 _,err7 := session.Id(v.Id).Update(&v)
		 if err7 != nil {
			 session.Rollback()
			 this.ReturnJson(map[string]string{"message":"修改菜品信息错误"},400)
		 }
	 }

	 err8 := session.Commit()
	 if err8 != nil {
		 this.ReturnJson(map[string]string{"message":"事务提交失败"},400)
	 }

 }


 /*
 	今日订单


  */
  func(this *OrderController)Today(){
  	var orders []*models.Order
  	t := time.Now()
  	engine := models.Engine.Where("repast_date = ?",t.Format("2006-01-02"))
  	this.condition(engine)
	  if isFloor := this.GetString("is_floor");isFloor != "" {
		 engine.OrderBy("floor")
	  }
  	//engineC := *engine







	  var page int =0
	  var err7 error
	  if pageStr :=this.GetString("page");pageStr!=""{
		  beego.Info(pageStr)
		  page,err7 = strconv.Atoi(pageStr)
		  if err7 !=nil {

		  }
		  page  = (page-1)*10

	  }



	  engineC := *engine
	  var order models.Order
	  count,_ := engineC.Count(&order)
	  err2 := engine.Find(&orders)
	  if err2 != nil {
		 //this.Ctx.WriteString("查询出错"+err2.Error())
		 this.ReturnJson(map[string]interface{}{"message":"查询出错"+err2.Error()},400)
		 //this.Ctx.WriteString(err2.Error())
	  }



	  ////其他数据

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

	  this.ReturnJson(map[string]interface{}{"data":orders,"count":count},200)
	  var ordersNew []*models.OrderNew
	  for _,v:= range orders{
		 //	ordern :=
		 ordersNew = append(ordersNew,&(models.OrderNew{v,(canteen[v.CanteenId]).(string),campus[v.CampusId].(string),building[v.BuildId].(string),"","",models.EatType[v.EatType],models.MealType[v.MealType],models.PayStatus[v.PayStatus],models.Status[v.Status],models.PayType[v.PayType]}))
	  }


	  this.ReturnJson(map[string]interface{}{"data":ordersNew,"count":count},200)



  }

  func(this *OrderController)BatchSign(){
  	orderIds := this.GetStrings("order_ids")
	  if len(orderIds)==0 {
		  this.ReturnJson(map[string]string{"message":"请输入订单id"},400)

	  }
  	var orderIdsNew []int
  	for _,v := range  orderIds{
  		id ,_ :=strconv.Atoi(v)
  		orderIdsNew = append(orderIdsNew,id)
	}
  	beego.Info(orderIdsNew)
  	var orders []*models.Log
  	err := models.Engine.In("id",orderIdsNew).Find(&orders)
	  if err != nil {
		  this.ReturnJson(map[string]string{"message":"查询出错"+err.Error()},400)
	  }

  	session := models.Engine.NewSession()

	  for _,v := range orders{
	  	beego.Info(v)
	  	v.Value= "vvvvv"
		  beego.Info(v)
		  _,err2 := session.Id(v.Id).Update(v)
		  if err2 != nil {
			  this.ReturnJson(map[string]string{"message":"更新出错"+err.Error()},400)
		  }
	  }



	  beego.Info("更新结束")
	  this.ReturnJson(map[string]string{"message":"更新成功"},200)



  }

  func (this *OrderController)CreateOrder(){
  	var order models.Order
  	//var carts []*models.Carts
  	goods := make(map[int64]models.Goods)
  	var user models.User
  	var goodsInput models.GoodsInput
  	var campus models.Campus
  	var canteen models.Canteen
  	var build models.Building
  	var area models.Area
  	var admin models.Admin





	  if userId := this.GetString("user_id");userId == ""{
		  this.ReturnJson(map[string]string{"message":"缺少用户id"},400)
	  }else{
	  	userIdInt,err := strconv.Atoi(userId)
		  if err != nil {
			  this.ReturnJson(map[string]string{"message":"输入的用户id不是数字"},400)
		  }

	  	e,err2 := models.Engine.Id(userIdInt).Get(&user)
		  if err2 != nil || !e {
			  this.ReturnJson(map[string]string{"message":"用户不存在"},400)
		  }

	  	order.UserId = userIdInt
	  }

	  err4 := json.Unmarshal(this.Ctx.Input.RequestBody,&goodsInput)
	  if err4 != nil {
		  this.ReturnJson(map[string]string{"message":"商品信息错误"},400)
	  }

	  err3 := models.Engine.In("id",goodsInput.GoodsId).Find(&goods)
	  if err3 != nil {
		  this.ReturnJson(map[string]string{"message":"无可售商品"+err3.Error()+"err3"},400)
	  }
	  //this.ReturnJson(map[string]interface{}{"data":goods},400)

	  beego.Info(goodsInput)
	  beego.Info(goods)

	  for  _,v := range goodsInput.Num{
		 order.Count += v
	  }




	  nowTime:= time.Now().Format("2006-01-02 15:04:05")
	  beego.Info(nowTime)


	  //if goodsStr := this.GetStrings("goods_ids");len(goodsStr)==0{
	  //	this.ReturnJson(map[string]string{"message":"请选择商品"},400)
	  //}else{
	  //	//查询到了商品
	  //	for _,v := range goodsStr.num{
	  //		goodsCount += v
		//}
	  //
	  //}
	  //就餐时间
	  if order.RepastDate =this.GetString("repast_date");order.RepastDate == "" {
		  this.ReturnJson(map[string]string{"message":"缺少就餐时间"},400)
	  }
	  //就餐方式
	  if eayType :=this.GetString("eat_type");eayType != "" {
	  		eatTypeInt,err5 := strconv.Atoi(eayType)
		  if err5 != nil {
			  this.ReturnJson(map[string]string{"message":"就餐方式传参错误"},400)
		  }
		  order.EatType = eatTypeInt
	  }else{
		  this.ReturnJson(map[string]string{"message":"就餐方式必选"},400)
	  }
	  //支付方式
	  if payType:=this.GetString("pay_type");payType != "" {
		  payTypeInt,err6 := strconv.Atoi(payType)
		  if err6 != nil {
			  this.ReturnJson(map[string]string{"message":"支付方式传参错误"},400)
		  }
		  order.PayType = payTypeInt
	  }else{
		  this.ReturnJson(map[string]string{"message":"缺少支付方式"},400)
	  }
		//签单单位id
	  if order.PayType ==2 {
		  signUnitId :=this.GetString("sign_unit_id")
		  if signUnitId == "" {
			  this.ReturnJson(map[string]string{"message":"缺少签单单位id"},400)
		  }
		  signUnitIdInt,err7 := strconv.Atoi(signUnitId)

		  if err7 != nil {
			  this.ReturnJson(map[string]string{"message":"请输入正确的签单单位id"},400)
		  }
		  order.SignUnitId = signUnitIdInt

		  //签单人
		  order.SignName = user.Name
	  }
	  //就餐方式
	  if mealType:=this.GetString("meal_type");mealType != "" {
		  mealTypeInt,err7 := strconv.Atoi(mealType)
		  if err7 != nil {
			  this.ReturnJson(map[string]string{"message":"餐次传参错误"},400)
		  }
		  order.MealType = mealTypeInt
	  }else{
		  this.ReturnJson(map[string]string{"message":"缺少餐次信息"},400)
	  }
	  //校区
	  if campusId :=this.GetString("campus_id");campusId != "" {
		  campusIdInt,err8 := strconv.Atoi(campusId)
		  if err8 != nil {
			  this.ReturnJson(map[string]string{"message":"校区信息传参错误"},400)
		  }
		  campusE,err9 := models.Engine.Id(campusIdInt).Get(&campus)
		  if err9 == nil || !campusE {
			  this.ReturnJson(map[string]string{"message":"校区信息不存在"},400)
		  }
		  order.CampusId = campusIdInt
	  }else{
		  this.ReturnJson(map[string]string{"message":"缺少校区信息"},400)
	  }
	//食堂
	  if canteenId :=this.GetString("canteen_id");canteenId != "" {
		  canteenIdInt,err8 := strconv.Atoi(canteenId)
		  if err8 != nil {
			  this.ReturnJson(map[string]string{"message":"食堂信息传参错误"},400)
		  }

		  canteenE,err10 := models.Engine.Id(canteenIdInt).Get(&canteen)
		  if err10 != nil && !canteenE {
			  this.ReturnJson(map[string]string{"message":"食堂信息不存在"},400)
		  }
		  order.CanteenId = canteenIdInt
	  }else{
		  this.ReturnJson(map[string]string{"message":"缺少食堂信息"},400)
	  }
	  //用户信息
	  order.UserId = user.Id
	  order.UserName = user.Name
	  order.UserMobile = user.Mobile

	  //楼宇 外卖
	  if order.EatType==1 {
		  buildId := this.GetString("build_id")
		  if buildId =="" {
			  this.ReturnJson(map[string]string{"message":"缺少楼宇信息"},400)
		  }
		  buildIdInt,err11 := strconv.Atoi(buildId)
		  if err11 != nil {
			  this.ReturnJson(map[string]string{"message":"请输入正确的楼宇id"},400)
		  }
		  buildE,err12 := models.Engine.Id(buildIdInt).Get(&build)
		  if err12!= nil|| !buildE {
			  this.ReturnJson(map[string]string{"message":"查无次楼宇信息"},400)
		  }
		  order.BuildId = buildIdInt

		  floor := this.GetString("floor")
		  if floor =="" {
			  this.ReturnJson(map[string]string{"message":"缺少楼层信息"},400)
		  }
		  floorInt,err12 := strconv.Atoi(floor)
		  if err12 != nil {
			  this.ReturnJson(map[string]string{"message":"请输入正确的楼宇id"},400)
		  }
		  order.Floor = floorInt

		  //详细地址

		  order.Address = campus.Name+" "+canteen.Name+ " "+floor+"楼 "+this.GetString("address")
		  //区域
		  areaE,err13 := models.Engine.Id(build.AreaId).Get(&area)
		  if err13!=nil|| !areaE {
			  this.ReturnJson(map[string]string{"message":"查无此区域"},400)
		  }
		  order.AreaId = build.AreaId
		  adminE,err14 := models.Engine.Id(area.AdminId).Get(&admin)
		  if err14 != nil || !adminE {
			  this.ReturnJson(map[string]string{"message":"骑手信息错误"},400)
		  }
		  order.RiderId = admin.Id
		  order.RiderName = admin.Name
		  order.RiderMobile = admin.Mobile


	  }
	  var timeEatStart,timeEatEnd time.Time
	  var err16,err17 error


	  //如果是午餐
	  if order.MealType ==2 {
	  	timeEatStart,err16 =time.Parse(models.BaseFormat,order.RepastDate+" "+canteen.LunchStartAt)
		  if err16 != nil {
			  this.ReturnJson(map[string]string{"message":"时间转换出错"+err16.Error()},400)
		  }
	  	timeEatEnd,err17 =time.Parse(models.BaseFormat,order.RepastDate+" "+canteen.LunchEndAt)
		  if err17 != nil {
			  this.ReturnJson(map[string]string{"message":"时间转换出错"},400)
		  }
	  }
	  //如果是午餐
	  if order.MealType ==3 {
		  timeEatStart,err16 =time.Parse(models.BaseFormat,order.RepastDate+" "+canteen.DinnerStartAt)
		  if err16 != nil {
			  this.ReturnJson(map[string]string{"message":"时间转换出错"},400)
		  }
		  timeEatEnd,err17 =time.Parse(models.BaseFormat,order.RepastDate+" "+canteen.DinnerStartAt)
		  if err17 != nil {
			  this.ReturnJson(map[string]string{"message":"时间转换出错"},400)
		  }
	  }
	  if len(goodsInput.GoodsId)==0 || len(goodsInput.Num)==0  {
		  this.ReturnJson(map[string]string{"message":"请输入商品参数"},400)
	  }

	  //验证菜品
	  for _,v := range goodsInput.GoodsId{
		  if goods[int64(v)].CanteenId != order.CanteenId  {
			  this.ReturnJson(map[string]string{"message":"菜品不属于该食堂"},400)
		  }
	  }
	  beego.Info("验证菜品结束")
	  order.EatStartAt = timeEatStart
	  order.EatEndAt = timeEatEnd.Add(time.Duration(canteen.OrderExpire*60*1e9))

	  order.CardNo  = user.CardNo

	  //时间
	  beego.Info("时间开始")
	  t := time.Now()
	  order.CreatedAt = t
	  order.UpdatedAt = t
	  //time.Now()

	  order.DiscountPrice = "0.00"




	  session := models.Engine.NewSession()
	  defer session.Close()
		beego.Info("插入开始1")
	  err18 := session.Begin()
	  if err18 !=  nil {
		  this.ReturnJson(map[string]string{"message":"只有未支付订单才可取消订单"},400)
	  }
	  beego.Info("插入开始2")
	  insertNum,err19 := session.InsertOne(&order)
	  if err19 != nil || insertNum ==0  {
		  session.Rollback()
		  this.ReturnJson(map[string]string{"message":"添加订单失败"+err19.Error()},400)
	  }
	  beego.Info("插入开始3")
	  var cartInsert []*models.Carts
	  var orderFinalPrice float64
	  for k,v:= range goodsInput.GoodsId{
	  	  goodsPrice ,err20 := strconv.ParseFloat(goods[int64(v)].Price,64)
		  if err20 != nil {
			  this.ReturnJson(map[string]string{"message":"菜品价格有误"},400)
		  }
		  finalPrice := goodsPrice*float64(goodsInput.Num[k])
		  orderFinalPrice += finalPrice

		  var cart models.Carts
		  cart.UserId = user.Id
		  cart.GoodsId = v
		  cart.CanteenId = canteen.Id
		  cart.TakeOutType = order.TakeOutType
		  cart.AreaId = order.AreaId
		  cart.EatType = order.EatType
		  cart.Num = goodsInput.Num[k]
		  cart.FinalPrice = strconv.FormatFloat(finalPrice,'f',2,64)
		  cart.OrderId = order.Id
		  cart.Status=1
		  cart.GoodsImage = goods[int64(k)].ImageUrl
		  cart.GoodsName = goods[int64(k)].Name
		  cart.CreatedAt = t
		  cart.UpdatedAt = t
		  cartInsert = append(cartInsert,&cart)

	  }
	  _,err21 := session.Insert(&cartInsert)
	  if err21 != nil {
		  session.Rollback()
		  this.ReturnJson(map[string]string{"message":"事务提交失败"+err21.Error()+"err21"},400)
	  }
	  //修改最终价格
	  order.DiscountPrice = strconv.FormatFloat(orderFinalPrice,'f',2,64)
	  _,err22 := session.Id(order.Id).Update(&order)
	  if err22 != nil {
	  		session.Rollback()
		  this.ReturnJson(map[string]string{"message":"事务提交失败"+err22.Error()+"err22"},400)
	  }

	  err23 := session.Commit()
	  if err23 != nil {
		  this.ReturnJson(map[string]string{"message":"事务提交失败"+err23.Error()+"err23"},400)
	  }

	  this.ReturnJson(map[string]string{"message":"生成订单成功"},400)

  }

  func(this *OrderController)TakeOutList(){
  	orderId := this.GetString("order_id")
	  if orderId=="" {
		  this.ReturnJson(map[string]string{"message":"缺少订单id"},400)
	  }
  	orderIdInt,err3 := strconv.Atoi(orderId)
	  if err3 != nil  {
		  this.ReturnJson(map[string]string{"message":"订单id是数字"},400)
	  }
  	var order models.Order
  	orderE,err := models.Engine.Id(orderIdInt).Get(&order)
  	if err != nil || !orderE {
		this.ReturnJson(map[string]string{"message":"订单已失效"},400)
	}

  	var list []*models.Admin
  	campusId := this.GetString("campus_id")

	  if orderId=="" {
		  this.ReturnJson(map[string]string{"message":"缺少订单id"},400)
	  }
	  campusIdInt,err4 := strconv.Atoi(campusId)

	  if err4 != nil {
		  this.ReturnJson(map[string]string{"message":"校区参数错误"},400)
	  }
  	err2 := models.Engine.Where("campus_id = ?",campusIdInt).Find(&list)
	  if err2 != nil {
		  this.ReturnJson(map[string]string{"message":"外卖员查询错误"},400)
	  }
  	var AdminList []*models.AdminValue
	  for _,v:= range list{
	  	var value int
		  if order.RiderId == v.Id {
			  value = 1
		  }else{
		  	  value = 0
		  }
	  	AdminList = append(AdminList,&models.AdminValue{v,value})
	  }


	  this.ReturnJson(map[string]interface{}{"data":AdminList},200)
  }


  func(this *OrderController)ChangeTakeOut(){

	  orderId := this.GetString("order_id")
	  if orderId=="" {
		  this.ReturnJson(map[string]string{"message":"缺少订单id"},400)
	  }
	  orderIdInt,err3 := strconv.Atoi(orderId)
	  if err3 != nil {
		  this.ReturnJson(map[string]string{"message":"订单id是数字"},400)
	  }
	  var order models.Order
	  orderE,err := models.Engine.Id(orderIdInt).Get(&order)
	  if err != nil || orderE {
		  this.ReturnJson(map[string]string{"message":"订单已失效"},400)
	  }

	  adminId := this.GetString("admin_id")
	  if adminId == "" {
		  this.ReturnJson(map[string]string{"message":"缺少外卖员id"},400)
	  }
	  adminIdInt,err4 := strconv.Atoi(adminId)
	  if err4 != nil {
		  this.ReturnJson(map[string]string{"message":"外卖员id不是数字"},400)
	  }
	  var admin models.Admin
	  adminE,err5 := models.Engine.Id(adminIdInt).Get(&admin)
	  if err5 != nil || !adminE {
		  this.ReturnJson(map[string]string{"message":"该外卖员信息缺失"},400)
	  }
	  order.RiderId = admin.Id
	  order.RiderName = admin.Name
	  order.RiderMobile = admin.Mobile
	  num,err6 := models.Engine.Id(order.Id).Update(&order)
	  if err6 != nil || num==0 {
		  this.ReturnJson(map[string]string{"message":"修改失败"},400)
	  }
	  this.ReturnJson(map[string]string{"message":"修改成功"},200)

  }

  func(this *OrderController)SignList(){
  	engine := models.Engine.NewSession()
  	this.condition(engine)
  	engine.Where("pay_type = ? AND status BETWEEN ? AND ?",2,1,7)
  	enginePrice := *engine
  	engineCount := *engine
  	var orderPrice models.Order
  	orderPriceF64,err := enginePrice.Sum(orderPrice,"discount_price")
	  if err != nil {
		  this.ReturnJson(map[string]string{"message":err.Error()},400)
	  }
  	beego.Info(orderPriceF64)

  	//count
  	orderCount,err8 := engineCount.Count(orderPrice)
	  if err8 != nil {
		  this.ReturnJson(map[string]string{"message":err8.Error()},400)
	  }
  	beego.Info(orderCount)


	  var page int =1
	  var err7 error
	  if pageStr :=this.GetString("page");pageStr!=""{
		  beego.Info(pageStr)
		  page,err7 = strconv.Atoi(pageStr)
		  if err7 !=nil {

		  }
		  beego.Info("strconv",page)
		  page  = (page-1)*10

	  }

	  var orders []*models.Order
	  err2 := engine.Limit(10,page).Find(&orders)
	  if err2 != nil {
		  this.ReturnJson(map[string]string{"message":"分页查询错误"},400)
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

	  //签单单位
	  unit,err7 := models.Pluck(models.Engine.Table("unit"),"name")
	  unit[0] = "未知"
	  if err7 != nil{
		  this.ReturnJson(err7,400)
	  }



	  for _,v:= range orders{
		  //	ordern :=
		  ordersNew = append(ordersNew,&(models.OrderNew{v,(canteen[v.CanteenId]).(string),campus[v.CampusId].(string),building[v.BuildId].(string),"",unit[v.SignUnitId].(string),models.EatType[v.EatType],models.MealType[v.MealType],models.PayStatus[v.PayStatus],models.Status[v.Status],models.PayType[v.PayType]}))
	  }

	  this.ReturnJson(map[string]interface{}{"data":ordersNew,"count":orderCount,"price_count":orderPriceF64},200)


  }

  type OrderIds struct {
  	OrderIds []int
  }


  func(this *OrderController) BatchSignNotified(){
  	userId := this.GetSession("user_id")
	  if userId != nil {
		  this.ReturnJson(map[string]string{"message":"找不到用户"},400)
	  }
  	var admin models.Admin
  	adminE,err := models.Engine.Id(userId).Get(&admin)
	  if err != nil || !adminE {
		  this.ReturnJson(map[string]string{"message":"操作人信息不存在"+err.Error()},400)
	  }

  	var orderIds OrderIds

	  if err2 := json.Unmarshal(this.Ctx.Input.RequestBody,&orderIds);err2 != nil {
		  this.ReturnJson(map[string]string{"message":"订单id错误"},400)
	  }

	  t := int(time.Now().Unix())
	  //orderDatas
		orders := make(map[int64]*models.Order)
		err3 := models.Engine.In("id",orderIds.OrderIds).Find(&orders)
	  if err3 != nil {
		  this.ReturnJson(map[string]string{"message":"订单查询错误"},400)
	  }
		var userIds []int
		for _,v := range orders{
			if v.PayType ==2 && v.PayStatus == 1 {
				userIds = append(userIds,v.UserId)
			}
		}
		users := make(map[int64]models.User)
		err4 := models.Engine.In("id",userIds).Find(&users)
		if err4 != nil{
			this.ReturnJson(map[string]string{"message":"用户查询错误"},400)
		}

	  //签单单位
	  unit,err7 := models.Pluck(models.Engine.Table("unit"),"name")
	  unit[0] = "未知"
	  if err7 != nil{
		  this.ReturnJson(err7,400)
	  }

	  for _,v := range orders{
		  if _,ok := users[int64(v.UserId)]; ok && v.PayType ==2 && v.PayStatus == 1 {
			  this.SendMessageTmplate2(map[string]string{"first":"您的签单订单还未结转,请及时结转",
			  "keyword1":"笔数:1笔,金额:"+v.DiscountPrice,
			  "keyword2":time.Now().Format(models.BaseFormat),
			  "keyword3":unit[v.SignUnitId].(string),
			  "time":strconv.Itoa(t),
			  "tourer":users[int64(v.UserId)].Openid,
			  })
		  }
	  }
  }



  /*
  			订单到处execl
   */
   func(this *OrderController)OrderExecl(){
   	var userId int = 7
	   //if userId := this.GetSession("user_id");userId == nil{
		//   this.ReturnJson(map[string]string{"message":"已过期请重新登录"},400)
	   //}

   	//当前操作人员信息
   	var user models.Admin
   	userE,err9 := models.Engine.Id(userId).Get(&user)
   	beego.Info(user)
   	if err9!= nil || !userE {
		this.ReturnJson(map[string]string{"message":"用户不存在"+err9.Error()},400)
	}



	   //其他数据
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
	   building ,err6 := models.Pluck(models.Engine.Table("building"),"name")
	   building[0] = "未知"
	   if err5 != nil{
		   this.ReturnJson(err6,400)
	   }

	   beego.Info(building)

	   //签单单位
	   area,err7 := models.Pluck(models.Engine.Table("area"),"name")
	   area[0] = "未知"
	   if err7 != nil{
		   this.ReturnJson(err7,400)
	   }

	   //形成一个slise 包着 slice
	   var orderExecl [][]string
	   orderExecl = append(orderExecl,[]string{"订单号","就餐日期","客户姓名","手机号","就餐方式","就餐餐次","套餐","数量","金额","地址","支付方式","配送员","配送员电话"})

	   engine := models.Engine.NewSession()

	   this.condition(engine)

	   var i = 2

	   aa:
	   for  {
	   		engineTmp := *engine
	   		var orders []*models.OrderExcel
	   		//err := engineTmp.Where("campus_id = ?",user.CampusId).Limit(50,(i-1)*50).Find(&orders)
	   		err := engineTmp.Table("order").Join("INNER","carts","carts.order_id = order.id").
	   			Select("`order`.*,`carts`.num,`carts`.goods_name,`carts`.final_price").
	   			Where("order.campus_id = ?",user.CampusId).
	   			Limit(50,(i-1)*50).Find(&orders)

	   		//err := models.Engine.Table("order").Join("INNER", "carts", "carts.order_id = order.id").Limit(10).Find(&orderGoods)
		   if err != nil {
			   this.ReturnJson(map[string]string{"message":"分页查询错误"+err.Error()},400)
		   }

	   		if len(orders) ==0 {
	   			break aa
			}
	   		for _,v := range orders{
	   			//var goods []*models.Carts

	   			address := canteen[v.CampusId].(string)+building[v.BuildId].(string)+strconv.Itoa(v.Floor)+"楼"+v.Address
	   			//address := ""
	   			//err10 := models.Engine.Where("order_id = ?",v.Id).Find(&goods)
				//if err10 != nil {
				//	this.ReturnJson(map[string]string{"message":"查询商品错误"},400)
				//}
	   			//if len(goods) == 0{
				//	orderExecl = append(orderExecl,[]string{v.OrderSn,v.RepastDate,v.UserName,v.UserMobile,models.EatType[v.EatType],
				//		models.MealType[v.MealType],"","","",
				//		address,models.PayType[v.PayType],v.RiderName,v.RiderMobile})
				//}else {
				//	for _,g := range goods{
				//		orderExecl = append(orderExecl,[]string{v.OrderSn,v.RepastDate,v.UserName,v.UserMobile,models.EatType[v.EatType],
				//			models.MealType[v.MealType],g.GoodsName,strconv.Itoa(g.Num),g.FinalPrice,
				//		address,models.PayType[v.PayType],v.RiderName,v.RiderMobile})
				//	}
				//}

				orderExecl = append(orderExecl,[]string{v.OrderSn,v.RepastDate,v.UserName,v.UserMobile,models.EatType[v.EatType],
						models.MealType[v.MealType],v.GoodsName,strconv.Itoa(v.Num),v.FinalPrice,
						address,models.PayType[v.PayType],v.RiderName,v.RiderMobile})

			}

	   		i++
	   }
	   //this.ReturnJson(map[string]interface{}{"data":"sss"},200)

		beego.Info(orderExecl)
	   ex := excelize.NewFile()
	   for k,v := range orderExecl{
		   for key,val := range v{
			   ex.SetCellValue("sheet1",controllers.ExeclMap[key]+strconv.Itoa(k+1),val)
			   //beego.Info(controllers.ExeclMap[key]+strconv.Itoa(k))
		   }

	   }
	  buff,err11 := ex.WriteToBuffer()
	   if err11 != nil {
		   this.ReturnJson(map[string]string{"message":"已过期请重新登录"},400)
	   }

	   this.Ctx.Output.Header("Content-Type", "application/octet-stream")
	   this.Ctx.Output.Header("Content-Disposition", "attachment; filename="+fmt.Sprintf("%s", "Book.xlsx"))//文件名
	   this.Ctx.Output.Header("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	   this.Ctx.Output.Header("Pragma", "no-cache")
	   this.Ctx.Output.Header("Expires", "0")
	   _,err12 := io.Copy(this.Ctx.ResponseWriter,buff)
	   if err12 != nil{
		   this.ReturnJson(map[string]string{"message":"写入header出错"},400)
	   }
   }


