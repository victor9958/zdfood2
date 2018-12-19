package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

type WechatMessage struct {
	C 			string 	`json:"c"`
	A			string 	`json:"a"`
	Template 	string 	`json:"template"`
	First 		string 	`json:"first"`
	Keyword1 	string 	`json:"keyword1"`
	Keyword2 	string 	`json:"keyword2"`
	Keyword3 	string 	`json:"keyword3"`
	Remark 		string 	`json:"remark"`
	Time 		int  	`json:"time"`
	Url 		string 	`json:"url"`
	Touser 		string 	`json:"touser"`


}

var ExeclMap map[int]string = map[int]string{
	0:"A",
	1:"B",
	2:"C",
	3:"D",
	4:"E",
	5:"F",
	6:"G",
	7:"H",
	8:"I",
	9:"J",
	10:"K",
	11:"L",
	12:"M",
	13:"N",
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


/*
	错误的 由于 golang map是无序的
 */
func(this *BaseController)SendMessageTmplate(t map[string]string)error{
	t["a"] = "publicSendMsgv2"
	t["c"] = "smart"
	if _,ok := t["template_id"];!ok {
		t["template_id"] = "NO0sydE7PgWcHcQjXzg0G5HBFCG1KImZXm0GYTM8t_4"   //一个默认的模板id
	}
	if _,ok := t["url"];ok {
		t["url"] = url.QueryEscape(t["url"])
	}
	//if _,ok := t["first"];ok {
	//	t["first"] = url.QueryEscape(t["first"])
	//}
	if _,ok := t["remark"];ok {
		t["remark"] = url.QueryEscape(t["remark"])
	}
	if _,ok := t["keyword1"];ok {
		t["keyword1"] = strings.Replace(t["keyword1"]," ",",",-1)
	}
	if _,ok := t["keyword2"];ok {
		t["keyword2"] = strings.Replace(t["keyword2"]," ",",",-1)
	}

	if _,ok := t["keyword3"];ok {
		t["keyword3"] = strings.Replace(t["keyword3"]," ",",",-1)
	}

	var url  string = "http://zulg.zju.edu.cn/api.php?"
	var key string

	for k,v := range t{
		key += k +"="+v
		url += k + "=" + v + "&"
	}

	key += "smart123456"
	beego.Info(key)
	md5Str := md5.New()
	md5Str.Write([]byte(key))
	md5Key := hex.EncodeToString(md5Str.Sum(nil))
	url += "sig=" + md5Key
	beego.Info(url)
	re,err3 := http.Get(url)
	if err3 != nil {
		return err3
	}
	defer  re.Body.Close()
	body ,err2 := ioutil.ReadAll(re.Body)

	if err2 != nil {
		return err2
	}
	beego.Info(string(body))

	return nil


}

func(this *BaseController)SendMessageTmplate2(t map[string]string)error{
	var urlStr  string = "http://zulg.zju.edu.cn/api.php?"
	var key string
	t["a"] = "publicSendMsgv2"
	t["c"] = "smart"

	urlStr +="c="+t["c"]+"&a="+t["a"]
	key +="c="+t["c"]+"a="+t["a"]



	if _,ok := t["first"];ok {
		//t["first"] = url.QueryEscape(t["first"])
		urlStr += "&first=" +t["first"]
		key += "first=" +t["first"]
	}

	if _,ok := t["keyword1"];ok {
		t["keyword1"] = strings.Replace(t["keyword1"]," ",",",-1)
		urlStr += "&keyword1=" +t["keyword1"]
		key += "keyword1=" +t["keyword1"]

	}
	if _,ok := t["keyword2"];ok {
		t["keyword2"] = strings.Replace(t["keyword2"]," ",",",-1)
		urlStr += "&keyword2=" +t["keyword2"]
		key += "keyword2=" +t["keyword2"]
	}

	if _,ok := t["keyword3"];ok {
		t["keyword3"] = strings.Replace(t["keyword3"]," ",",",-1)
		urlStr += "&keyword3=" +t["keyword3"]
		key += "keyword3=" +t["keyword3"]
	}
	if _,ok := t["remark"];ok {
		t["remark"] = url.QueryEscape(t["remark"])
		urlStr += "&remark=" +t["remark"]
		key += "remark=" +t["remark"]
	}
	if _,ok := t["time"];ok {
		urlStr += "&time=" +t["time"]
		key += "time=" +t["time"]
	}
	if _,ok := t["url"];ok {
		t["url"] = url.QueryEscape(t["url"])
		urlStr += "&url=" +t["url"]
		key += "url=" +t["url"]
	}

	if _,ok := t["template_id"];!ok {
		t["template_id"] = "NO0sydE7PgWcHcQjXzg0G5HBFCG1KImZXm0GYTM8t_4"   //一个默认的模板id
		urlStr += "&template_id=" +t["template_id"]
		key += "template_id=" +t["template_id"]
	}

	if _,ok := t["touser"];ok {
		urlStr += "&touser=" +t["touser"]
		key += "touser=" +t["touser"]
	}



	//for k,v := range t{
	//	key += k +"="+v
	//	url += k + "=" + v + "&"
	//}

	key += "smart123456"
	beego.Info(key)
	md5Str := md5.New()
	md5Str.Write([]byte(key))
	md5Key := hex.EncodeToString(md5Str.Sum(nil))
	urlStr += "&sig=" + md5Key
	beego.Info(urlStr)
	re,err3 := http.Get(urlStr)
	if err3 != nil {
		return err3
	}
	defer  re.Body.Close()
	body ,err2 := ioutil.ReadAll(re.Body)

	if err2 != nil {
		return err2
	}
	beego.Info(string(body))

	return nil


}

