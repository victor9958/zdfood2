package finance

import (
	"github.com/go-xorm/xorm"
	"strconv"
	"zdfood2/controllers"
	"zdfood2/models"
)

type OrderStatistics struct {
	OrderAllCount int `json:"order_all_count"`
	PayAllCount string `json:"pay_all_count"`
	SignAll string `json:"sign_all"`
	ElectronicAll string `json:"electronic_all"`
	CampusId int `json:"campus_id"`
	CanteenId int `json:"canteen_id"`
}

type OrderStatisticsNew struct {
	*OrderStatistics
	CampusName string `json:"campus_name"`
	CanteenName string `json:"canteen_name"`
}

type FinanceController struct {
	controllers.BaseController
}

func(this *FinanceController)OrderStatistics(){
	var finances []*OrderStatistics
	engine := models.Engine.NewSession()
	this.GetCondition(engine)
	err := engine.Table("finance").
		Select("sum(order_count) as order_all_count,sum(day_pay_count) as pay_all_count,sum(sign_pay_count) as sign_all,sum(electronic_pay_count) as electronic_all,campus_id,canteen_id").
		GroupBy("campus_id,canteen_id").
		Where("meal_type = ?",4).
		Find(&finances)
	if err != nil {
		this.ReturnJson(map[string]string{"message":err.Error()},400)
	}

	//var campus map[string]models.Campus
	campus,err2:= models.PluckCampusName(models.Engine.NewSession())
	if err2 != nil {
		this.ReturnJson(map[string]string{"message":err2.Error()},400)
	}
	canteen,err3:= models.PluckCanteenName(models.Engine.NewSession())
	if err3 != nil {
		this.ReturnJson(map[string]string{"message":err3.Error()},400)
	}


	orderAllCount := 0
	var  payAllCount, signAll ,electronicAll float64
	var campusName string = "合计"
	canteenName := ""

	var financesNew []*OrderStatisticsNew
	for _,v := range finances{
		financesNew = append(financesNew,&OrderStatisticsNew{v,campus[v.CampusId].(string),canteen[v.CanteenId].(string)})
		orderAllCount +=v.OrderAllCount

		payAllCountTemp,err4 := strconv.ParseFloat(v.PayAllCount,64)
		if err4 != nil {
			this.ReturnJson(map[string]string{"message":err4.Error()},400)
		}
		payAllCount += payAllCountTemp

		signAllTemp,err5 := strconv.ParseFloat(v.SignAll,64)
		if err5 != nil {
			this.ReturnJson(map[string]string{"message":err5.Error()},400)
		}
		signAll += signAllTemp

		electronicAllTemp,err6 := strconv.ParseFloat(v.ElectronicAll,64)
		if err6 != nil {
			this.ReturnJson(map[string]string{"message":err6.Error()},400)
		}
		electronicAll += electronicAllTemp
	}

	financesNew = append(financesNew,&OrderStatisticsNew{&OrderStatistics{OrderAllCount:orderAllCount,PayAllCount:strconv.FormatFloat(payAllCount,'f',2,64),SignAll:strconv.FormatFloat(signAll,'f',2,64),ElectronicAll:strconv.FormatFloat(electronicAll,'f',2,64),CanteenId:0,CampusId:0},campusName,canteenName})

	this.ReturnJson(map[string]interface{}{"data":financesNew},200)

}
func(this *FinanceController)GetCondition(session *xorm.Session){

}
