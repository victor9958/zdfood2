package controllers

import (
	"net/url"
	"net/http"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetCode() {
	u,_:= url.Parse("http://www.baidu.com")
	q := u.Query()
	u.RawQuery = q.Encode()
	res,err := http.Get(u.String())
	if err != nil {
		c.Ctx.WriteString("0"+" err")
		return
	}
	resCode :=res.StatusCode
	res.Body.Close()
	//if  {
	//
	//}

	c.ReturnJson(200,resCode)
}
