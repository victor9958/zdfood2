package controllers

import "github.com/astaxie/beego"

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

