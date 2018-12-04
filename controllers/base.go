package controllers

import (
	"github.com/astaxie/beego"
	"reflect"
)

type BaseController struct {
	beego.Controller
	isLogin bool
}

type Page struct {
	PageLimit int
	PageOffest int
	Count int
}

func (c *BaseController) Prepare()  {

	//userLogin := c.GetSession("userLogin")
	//if userLogin != nil{
	//	c.isLogin = false
	//	beego.Info("false")
	//}else{
	//	c.isLogin = true
	//	beego.Info("true")
	//}

	c.Data["isLogin"] = true

}

func(c *BaseController)ReturnJson(data interface{},status int){

	c.Ctx.Output.Status=status
	c.Ctx.Output.JSON(data,true,false)
}


func(c *BaseController)DisposeConditionStr(conditionArr... string)(string){
	var conditionStr string
	for _,v :=range conditionArr {
		conditionStr += v+" AND "
	}
	return  string(conditionStr[0:len(conditionStr)-5])
}

/*
结构体 转化 map
 */
func StructToMap(obj interface{}) map[string]interface{}{
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

