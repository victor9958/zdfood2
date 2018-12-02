package moban

import (
	"encoding/json"
	"fmt"
	"github.com/henrylee2cn/faygo"
	l "server/jianzhuang_logic"
	"server/logic"
	u "server/units"
	"strconv"
	"strings"
	"time"
)

var GetBuildIndex = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取可用功能失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用功能失败,user_id必须为整型"))
	}
	var project_id int64
	if c.Param("project_id") == "" {
		//获取默认的项目编号
		project_id = l.GetLastProjectId(user_id)
		if project_id == 0 {
			project_id = 0
			//return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用功能失败,您没有被授权的项目!"))
		}
	} else {
		project_id, err = strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用功能失败,project_id必须为整型"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用功能失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	mi, err := l.GetBuildIndex(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(mi)
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddSecurityCheck = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加安全检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加安全检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查失败,project_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加安全检查失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查失败,company_id必须为整型"))
	}
	var assign_to, img_list []string
	if c.Param("assign_to") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加安全检查失败,请传入assign_to!"))
	}
	assign_to = strings.Split(c.Param("assign_to"), ",")

	if strings.Trim(c.Param("limit_time"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加安全检查失败,请传入limit_time!"))
	}
	//assign_to, err := strconv.ParseInt(c.Param("assign_to"), 10, 64)
	//if err != nil {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查失败,assign_to必须为整型"))
	//}

	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}

	var param string
	if strings.Trim(c.Param("param"), " ") != "" {
		param = strings.Trim(c.Param("param"), " ")
	}
	var materialsListId int64 = 0
	if strings.Trim(c.Param("materials_list_id"), " ") != "" {
		materialsListId, err = strconv.ParseInt(strings.Trim(c.Param("materials_list_id"), " "), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查失败,materials_list_id必须为整型!"))
		}
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddSecurityCheck(user_id, project_id, company_id, materialsListId, bind_type, c.Param("check_item"), c.Param("check_content"),
		strings.Trim(c.Param("limit_time"), " "), assign_to, img_list, param)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加安全检查失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加安全检查成功"))
})

var AddQualityTesting = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加质量检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加质量检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查失败,project_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加质量检查失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查失败,company_id必须为整型"))
	}
	if c.Param("assign_to") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加质量检查失败,请传入assign_to!"))
	}
	//assign_to, err := strconv.ParseInt(c.Param("assign_to"), 10, 64)
	//if err != nil {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查失败,assign_to必须为整型"))
	//}

	if strings.Trim(c.Param("limit_time"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加质量检查失败,请传入limit_time!"))
	}

	assign_to := strings.Split(c.Param("assign_to"), ",")
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var materialsListId int64 = 0
	if strings.Trim(c.Param("materials_list_id"), " ") != "" {
		materialsListId, err = strconv.ParseInt(strings.Trim(c.Param("materials_list_id"), " "), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查失败,materials_list_id必须为整型!"))
		}
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	var param string
	if strings.Trim(c.Param("param"), " ") != "" {
		param = strings.Trim(c.Param("param"), " ")
	}

	b, err := l.AddQualityTesting(user_id, project_id, company_id, materialsListId, bind_type, c.Param("check_item"), c.Param("check_content"),
		strings.Trim(c.Param("limit_time"), " "), assign_to, img_list, param)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加质量检查失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加质量检查成功"))
})

var AddFieldInspection = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加协作巡查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加协作巡查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查失败,project_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加协作巡查失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查失败,company_id必须为整型"))
	}
	if c.Param("assign_to") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加协作巡查失败,请传入assign_to!"))
	}

	if strings.Trim(c.Param("limit_time"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加协作巡查失败,请传入limit_time!"))
	}
	assign_to, err := strconv.ParseInt(c.Param("assign_to"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查失败,assign_to必须为整型"))
	}
	var re_assign_to []string
	if c.Param("re_assign_to") == "" {
		re_assign_to = strings.Split(c.Param("re_assign_to"), ",")
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var materialsListId int64 = 0
	if strings.Trim(c.Param("materials_list_id"), " ") != "" {
		materialsListId, err = strconv.ParseInt(strings.Trim(c.Param("materials_list_id"), " "), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查失败,materials_list_id必须为整型!"))
		}
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	var param string
	if strings.Trim(c.Param("param"), " ") != "" {
		param = strings.Trim(c.Param("param"), " ")
	}
	b, err := l.AddFieldInspection(user_id, project_id, company_id, assign_to, materialsListId, bind_type, c.Param("check_item"),
		c.Param("check_content"), strings.Trim(c.Param("limit_time"), " "), re_assign_to, img_list, param)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加协作巡查失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加协作巡查成功"))
})

var EditSecurityCheck = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改安全检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改安全检查失败,user_id必须为整型"))
	}
	if c.Param("sc_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改安全检查失败,请传入sc_id!"))
	}
	sc_id, err := strconv.ParseInt(c.Param("sc_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改安全检查失败,sc_id必须为整型"))
	}
	if c.Param("status") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改安全检查失败,请传入status!"))
	}
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改安全检查失败,status必须为整型"))
	}
	if c.Param("assign_to") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改安全检查失败,请传入assign_to!"))
	}
	//assign_to, err := strconv.ParseInt(c.Param("assign_to"), 10, 64)
	//if err != nil {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改安全检查失败,assign_to必须为整型"))
	//}
	assign_to := strings.Split(c.Param("assign_to"), ",")
	var materialsListId int64 = 0
	if strings.Trim(c.Param("materials_list_id"), " ") != "" {
		materialsListId, err = strconv.ParseInt(strings.Trim(c.Param("materials_list_id"), " "), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改安全检查失败,materials_list_id必须为整型!"))
		}
	}
	b, err := l.EditSecurityCheck(user_id, sc_id, materialsListId, c.Param("check_item"), c.Param("check_content"), status, assign_to)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改安全检查失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改安全检查成功"))
})

var EditQualityTesting = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改质量检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改质量检查失败,user_id必须为整型"))
	}
	if c.Param("qt_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改质量检查失败,请传入qt_id!"))
	}
	qt_id, err := strconv.ParseInt(c.Param("qt_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改质量检查失败,qt_id必须为整型"))
	}
	if c.Param("status") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改质量检查失败,请传入status!"))
	}
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改质量检查失败,status必须为整型"))
	}
	if c.Param("assign_to") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改质量检查失败,请传入assign_to!"))
	}
	//assign_to, err := strconv.ParseInt(c.Param("assign_to"), 10, 64)
	//if err != nil {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改质量检查失败,assign_to必须为整型"))
	//}
	assign_to := strings.Split(c.Param("assign_to"), ",")
	var materialsListId int64 = 0
	if strings.Trim(c.Param("materials_list_id"), " ") != "" {
		materialsListId, err = strconv.ParseInt(strings.Trim(c.Param("materials_list_id"), " "), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改质量检查失败,materials_list_id必须为整型!"))
		}
	}

	b, err := l.EditQualityTesting(user_id, qt_id, materialsListId, c.Param("check_item"), c.Param("check_content"), status, assign_to)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改质量检查失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改质量检查成功"))
})

var EditFieldInspection = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改协作巡查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改协作巡查失败,user_id必须为整型"))
	}
	if c.Param("fi_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改协作巡查失败,请传入fi_id!"))
	}
	fi_id, err := strconv.ParseInt(c.Param("fi_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改协作巡查失败,fi_id必须为整型"))
	}
	if c.Param("status") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改协作巡查失败,请传入status!"))
	}
	status, err := strconv.Atoi(c.Param("status"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改协作巡查失败,status必须为整型"))
	}
	if c.Param("assign_to") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改协作巡查失败,请传入assign_to!"))
	}
	assign_to, err := strconv.ParseInt(c.Param("assign_to"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改协作巡查失败,assign_to必须为整型"))
	}
	var re_assign_to []string
	if c.Param("re_assign_to") != "" {
		re_assign_to = strings.Split(c.Param("re_assign_to"), ",")
		//return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改协作巡查失败,请传入assign_to!"))
	}
	var materialsListId int64 = 0
	if strings.Trim(c.Param("materials_list_id"), " ") != "" {
		materialsListId, err = strconv.ParseInt(strings.Trim(c.Param("materials_list_id"), " "), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改协作巡查失败,materials_list_id必须为整型!"))
		}
	}
	b, err := l.EditFieldInspection(user_id, fi_id, assign_to, materialsListId, c.Param("check_item"), c.Param("check_content"), status, re_assign_to)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改协作巡查失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改协作巡查成功"))
})

var GetSecurityCheckList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取安全检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取安全检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败,project_id必须为整型"))
	}
	var page_num, count int64
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，count必须为整型"))
		}
	}
	var keyword string
	if c.Param("keyword") != "" {
		keyword = strings.Trim(c.Param("keyword"), " ")
	}
	var status int
	if c.Param("status") != "" {
		status, err = strconv.Atoi(c.Param("status"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，status必须为整型"))
		}
	} else {
		status = -1
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}

	sc, count, err := l.GetSecurityCheck(user_id, project_id, int((page_num-1)*count), int(count), status, bind_type, keyword)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(sc)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s,"total_count":%d}`, code, msg, string(b), count))
})

var GetQualityTestingList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取质量检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取质量检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取质量检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取质量检查失败,project_id必须为整型"))
	}
	var page_num, count int64
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取质量检查失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取质量检查失败，count必须为整型"))
		}
	}
	var keyword string
	if c.Param("keyword") != "" {
		keyword = strings.Trim(c.Param("keyword"), " ")
	}
	var status int
	if c.Param("status") != "" {
		status, err = strconv.Atoi(c.Param("status"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，status必须为整型"))
		}
	} else {
		status = -1
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	qt, count, err := l.GetQualityTesting(user_id, project_id, int((page_num-1)*count), int(count), status, bind_type, keyword)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(qt)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s,"total_count":%d}`, code, msg, string(b), count))
})

var GetFieldInspectionList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取协作巡查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取协作巡查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取协作巡查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取协作巡查失败,project_id必须为整型"))
	}
	var page_num, count int64
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取协作巡查失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取协作巡查失败，count必须为整型"))
		}
	}
	var keyword string
	if c.Param("keyword") != "" {
		keyword = strings.Trim(c.Param("keyword"), " ")
	}
	var status int
	if c.Param("status") != "" {
		status, err = strconv.Atoi(c.Param("status"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，status必须为整型"))
		}
	} else {
		status = -1
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	fi, count, err := l.GetFieldInspection(user_id, project_id, int((page_num-1)*count), int(count), status, bind_type, keyword)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(fi)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s,"total_count":%d}`, code, msg, string(b), count))
})

var GetBuildFindIndex = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取建装端发现数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取建装端发现数据失败,user_id必须为整型"))
	}
	var project_id int64
	if c.Param("project_id") != "" {
		project_id, err = strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取建装端发现数据失败,project_id必须为整型"))
		}
		//return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取建装端发现数据失败,请传入project_id!"))
	} else {
		project_id = 0
	}

	if c.Param("city_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取建装端发现数据失败,缺少参数city_id"))
	}
	city_id, err := strconv.ParseInt(c.Param("city_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取建装端发现数据失败,city_id必须为整型"))
	}
	data, err := l.GetBuildFindIndex(user_id, project_id, city_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetBuildNotice = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取建装端发现数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取建装端发现数据失败,user_id必须为整型"))
	}
	if c.Param("city_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取建装端发现数据失败,缺少参数city_id"))
	}
	city_id, err := strconv.ParseInt(c.Param("city_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取建装端发现数据失败,city_id必须为整型"))
	}
	var page_num, count int64
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取建装端发现数据失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取建装端发现数据失败，count必须为整型"))
		}
	}
	data, err := l.GetBuildNotice(user_id, city_id, int((page_num-1)*count), int(count))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetCompanyShow = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取展示主页数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取展示主页数据失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取展示主页数据失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取展示主页数据失败,company_id必须为整型"))
	}
	data, err := l.GetCompanyShow(user_id, company_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var SetCompanyInfo = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "设置公司信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置公司信息失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "设置公司信息失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置公司信息失败,company_id必须为整型"))
	}

	b, err := l.SetCompanyInfo(user_id, company_id, strings.Trim(c.Param("name"), " "),
		strings.Trim(c.Param("desc"), " "), strings.Trim(c.Param("addr"), " "),
		strings.Trim(c.Param("phone"), " "), strings.Trim(c.Param("biz_area"), " "),
		strings.Trim(c.Param("logo"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "success"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "success"))
})

var GetCompanyOrderList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取建装端发现数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取建装端发现数据失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加质量检查失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查失败,company_id必须为整型"))
	}
	data, err := l.GetCompanyOrderList(user_id, company_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetConstructionCommentList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取施工讲评台数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台数据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取施工讲评台数据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台数据失败,project_id必须为整型"))
	}
	var page_num, count int64
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台失败，count必须为整型"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetConstructionCommentList(user_id, project_id, int((page_num-1)*count), int(count), bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetConstructionCommentTotal = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取施工讲评台数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台数据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取施工讲评台数据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台数据失败,project_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	count, err := l.GetConstructionCommentTotal(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}

	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","count":%d}`, code, msg, count))
})

var GetConstructionComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取施工讲评台数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台数据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取施工讲评台数据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台数据失败,project_id必须为整型"))
	}
	if c.Param("comment_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取施工讲评台数据失败,请传入comment_id!"))
	}
	comment_id, err := strconv.ParseInt(c.Param("comment_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工讲评台数据失败,comment_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}

	data, err := l.GetConstructionComment(user_id, project_id, comment_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddConstructionComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加施工讲评台失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工讲评台失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加施工讲评台失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工讲评台失败,project_id必须为整型"))
	}
	if project_id < 1 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工讲评台失败,请先创建项目"))
	}
	if strings.Trim(c.Param("title"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加施工讲评台失败,请传入title!"))
	}
	if strings.Trim(c.Param("content"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加施工讲评台失败,请传入content!"))
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工讲评台失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddConstructionComment(user_id, project_id, bind_type, strings.Trim(c.Param("title"), " "), strings.Trim(c.Param("content"), " "), img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加施工讲评台失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加施工讲评台成功"))
})

var ModifyConstructionComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改施工讲评台失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改施工讲评台失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改施工讲评台失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改施工讲评台失败,project_id必须为整型"))
	}
	if c.Param("comment_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改施工讲评台失败,请传入comment_id!"))
	}
	comment_id, err := strconv.ParseInt(c.Param("comment_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改施工讲评台失败,comment_id必须为整型"))
	}

	if strings.Trim(c.Param("content"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改施工讲评台失败,请传入content!"))
	}
	if strings.Trim(c.Param("title"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改施工讲评台失败,请传入title!"))
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改施工讲评台失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.ModifyConstructionComment(user_id, project_id, comment_id, bind_type, strings.Trim(c.Param("title"), " "),
		strings.Trim(c.Param("content"), " "), img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改施工讲评台失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改施工讲评台成功"))
})

var DeleteConstructionComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "删除施工讲评台失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工讲评台失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "删除施工讲评台失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工讲评台失败,project_id必须为整型"))
	}
	if c.Param("comment_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "删除施工讲评台失败,请传入project_id!"))
	}
	comment_id, err := strconv.ParseInt(c.Param("comment_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工讲评台失败,project_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改施工讲评台失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.DeleteConstructionComment(user_id, project_id, comment_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除施工讲评台失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除施工讲评台成功"))
})

var GetEngineeringDirectory = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取工程名录失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工程名录失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取工程名录失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工程名录失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetEngineeringDirectory(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetRoleList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取工程名录失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工程名录失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取工程名录失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工程名录失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}

	var parent_id int64
	if c.Param("parent_id") == "" {
		parent_id = 0
	} else {
		parent_id, err = strconv.ParseInt(c.Param("parent_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工程名录失败,parent_id必须为整型"))
		}
	}
	if parent_id > 0 {
		data, err := l.GetRoleList(user_id, project_id, parent_id)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}
	data, err := l.GetRoleAllList(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetDefaultRoleList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取默认角色失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取默认角色失败,user_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取默认角色失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetDefaultRoleList(user_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddCustomRole = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加自定义角色失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加自定义角色失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加自定义角色失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加自定义角色失败,project_id必须为整型"))
	}
	if c.Param("parent_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加自定义角色失败,请传入project_id!"))
	}
	parent_id, err := strconv.ParseInt(c.Param("parent_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加自定义角色失败,parent_id必须为整型"))
	}
	if strings.Trim(c.Param("role_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加自定义角色失败,role_name不能为空!"))
	}
	is_construction_team := false
	if c.Param("is_construction_team") != "" {
		is_construction_team, err = strconv.ParseBool(c.Param("is_construction_team"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加自定义角色失败,is_construction_team必须为布尔型"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddCustomRole(user_id, project_id, parent_id, bind_type, strings.Trim(c.Param("role_name"), " "), is_construction_team)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加自定义角色失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加自定义角色成功"))
})

var ModifyUserRight = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改用户权限失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改用户权限失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改用户权限失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改用户权限失败,project_id必须为整型"))
	}
	if c.Param("modify_object_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改用户权限失败,请传入modify_object_id!"))
	}
	modify_object_id, err := strconv.ParseInt(c.Param("modify_object_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改用户权限失败,modify_object_id必须为整型"))
	}
	if strings.Trim(c.Param("right_value"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改用户权限失败,right_value不能为空!"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.ModifyUserRight(user_id, project_id, modify_object_id, bind_type, strings.Trim(c.Param("right_value"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改用户权限失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改用户权限成功"))
})

var AddUserByDirectory = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,project_id必须为整型"))
	}
	if c.Param("role_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,请传入role_id!"))
	}
	role_id, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,role_id必须为整型"))
	}
	if strings.Trim(c.Param("name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,请传入name!"))
	}
	if strings.Trim(c.Param("phone"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,请传入phone!"))
	}
	if c.Param("is_send_sms") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,请传入phone!"))
	}
	var is_send_sms = true
	if c.Param("is_send_sms") != "" {
		is_send_sms, err = strconv.ParseBool(c.Param("is_send_sms"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录增加用户及权限失败,is_send_sms必须为布尔型"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddUserByDirectory(user_id, project_id, role_id, bind_type, strings.Trim(c.Param("name"), " "),
		strings.Trim(c.Param("phone"), " "), strings.Trim(c.Param("id_card_no"), " "),
		strings.Trim(c.Param("join_date"), " "), strings.Trim(c.Param("bank_card"), " "),
		strings.Trim(c.Param("bank_name"), " "), strings.Trim(c.Param("certificate_name"), " "),
		strings.Trim(c.Param("certificate_code"), " "), strings.Trim(c.Param("certificate_validity"), " "), is_send_sms)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "工程名录增加用户及权限失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "工程名录增加用户及权限成功"))
})

var EditUserByDirectory = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录修改用户信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录修改用户信息失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录修改用户信息失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录修改用户信息失败,project_id必须为整型"))
	}
	if c.Param("object_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录修改用户信息失败,请传入object_id!"))
	}
	object_id, err := strconv.ParseInt(c.Param("object_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "工程名录修改用户信息失败,object_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.EditUserByDirectory(user_id, project_id, object_id, bind_type, strings.Trim(c.Param("name"), " "),
		strings.Trim(c.Param("id_card_no"), " "),
		strings.Trim(c.Param("join_date"), " "), strings.Trim(c.Param("bank_card"), " "),
		strings.Trim(c.Param("bank_name"), " "), strings.Trim(c.Param("certificate_name"), " "),
		strings.Trim(c.Param("certificate_code"), " "), strings.Trim(c.Param("certificate_validity"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "工程名录修改用户信息失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "工程名录修改用户信息成功"))
})

var GetUserRight = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户权限失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户权限失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户权限失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户权限失败,project_id必须为整型"))
	}
	if c.Param("object_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户权限失败,object_id!"))
	}
	object_id, err := strconv.ParseInt(c.Param("object_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改用户权限失败,object_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetUserRight(user_id, project_id, object_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetMyRight = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用权限失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用权限失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用权限失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用权限失败,project_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可用权限失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	if user_id < 1 {
		user_id = 6
		project_id = 1
	}
	data, err := l.GetMyRight(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetRoleRight = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取角色权限失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取角色权限失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取角色权限失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取角色权限失败,project_id必须为整型"))
	}
	if c.Param("role_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取角色权限失败,请传入role_id!"))
	}
	role_id, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取角色权限失败,role_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetRoleRight(user_id, project_id, role_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var DeleteUserByDirectory = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的用户失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的用户失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的用户失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的用户失败,project_id必须为整型"))
	}
	if c.Param("object_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的用户失败,请传入object_id!"))
	}
	object_id, err := strconv.ParseInt(c.Param("object_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的用户失败,object_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的用户失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.DeleteUserByDirectory(user_id, project_id, object_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除工程名录的用户失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除工程名录的用户成功"))
})

var ModifyRoleRight = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改角色权限失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改角色权限失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改角色权限失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改角色权限失败,project_id必须为整型"))
	}
	if c.Param("role_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改角色权限失败,请传入role_id!"))
	}
	role_id, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改角色权限失败,role_id必须为整型"))
	}
	if strings.Trim(c.Param("right_value"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改角色权限失败,right_value不能为空!"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.ModifyRoleRight(user_id, project_id, role_id, bind_type, strings.Trim(c.Param("right_value"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改角色权限失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改角色权限成功"))
})

var DeleteRoleByDirectory = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的角色失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的角色失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的角色失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的角色失败,project_id必须为整型"))
	}
	if c.Param("role_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的角色失败,请传入role_id!"))
	}
	role_id, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除工程名录的角色失败,role_id必须为整型"))
	}
	b, err := l.DeleteRoleByDirectory(user_id, project_id, role_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除工程名录的角色失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除工程名录的角色成功"))
})

var GetCompanyByName = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司名称失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司名称失败,user_id必须为整型"))
	}
	if strings.Trim(c.Param("name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司名称失败,请传入name!"))
	}
	data, err := l.GetCompanyByName(user_id, strings.Trim(c.Param("name"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetProjectByName = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司名称失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司名称失败,user_id必须为整型"))
	}
	if strings.Trim(c.Param("name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司名称失败,请传入name!"))
	}
	data, err := l.GetProjectByName(user_id, strings.Trim(c.Param("name"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var BindProject = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,user_id必须为整型"))
	}
	if c.Param("company_name") == "" && (c.Param("company_id") == "" || strings.Trim(c.Param("company_id"), " ") == "0") {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,请传入company_name或者company_id!"))
	}
	var company_id, project_id int64
	if strings.Trim(c.Param("company_id"), " ") != "" {
		company_id, err = strconv.ParseInt(c.Param("company_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,company_id必须为整型"))
		}
	} else {
		company_id = 0
	}

	if c.Param("project_name") == "" && (c.Param("project_id") == "" || strings.Trim(c.Param("project_id"), " ") == "0") {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,请传入project_name或者project_id!"))
	}
	if strings.Trim(c.Param("project_id"), " ") != "" {
		project_id, err = strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,project_id必须为整型"))
		}
	} else {
		project_id = 0
	}

	if c.Param("project_addr") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,请传入project_addr!"))
	}
	if c.Param("default_role_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,请传入default_role_id!"))
	}
	default_role_id, err := strconv.ParseInt(c.Param("default_role_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,default_role_id必须为整型"))
	}
	var start_date string
	if strings.Trim(c.Param("start_date"), " ") != "" {
		start_date = strings.Trim(c.Param("start_date"), " ")
	} else {
		start_date = "1970-01-01"
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	var lon, lat float64
	if strings.Trim(c.Param("lon"), " ") != "" {
		lon, err = strconv.ParseFloat(strings.Trim(c.Param("lon"), " "), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建楼盘信息失败,lon必须为数字型"))
		}
	}
	if strings.Trim(c.Param("lat"), " ") != "" {
		lat, err = strconv.ParseFloat(strings.Trim(c.Param("lat"), " "), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建楼盘信息失败,lat必须为数字型"))
		}
	}
	data, err := l.BindProject(user_id, company_id, project_id, default_role_id, strings.Trim(c.Param("project_name"), " "),
		strings.Trim(c.Param("company_name"), " "), strings.Trim(c.Param("project_addr"), " "), start_date, bind_type, lon, lat)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddCapitalLedger = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加资金总账失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加资金总账失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加资金总账失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加资金总账失败,project_id必须为整型"))
	}
	if c.Param("amount") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加资金总账失败,请传入amount!"))
	}
	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加资金总账失败,amount必须为数字"))
	}
	if c.Param("payments_type") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加资金总账失败,请传入payments_type!"))
	}
	payments_type, err := strconv.ParseInt(c.Param("payments_type"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加资金总账失败,payments_type必须为整型"))
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var accounting_date string
	if strings.Trim(c.Param("accounting_date"), " ") != "" {
		accounting_date = strings.Trim(c.Param("accounting_date"), " ")
	} else {
		accounting_date = time.Now().Format("2006-01-02")
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddCapitalLedger(user_id, project_id, payments_type, bind_type, amount, strings.Trim(c.Param("memo"), " "), accounting_date, img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加资金总账失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加资金总账成功"))
})

var ModifyCapitalLedger = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改资金总账失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改资金总账失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改资金总账失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改资金总账失败,project_id必须为整型"))
	}
	if c.Param("amount") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改资金总账失败,请传入amount!"))
	}
	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改资金总账失败,amount必须为数字"))
	}
	if c.Param("capital_ledger_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改资金总账失败,请传入capital_ledger_id!"))
	}
	capital_ledger_id, err := strconv.ParseInt(c.Param("capital_ledger_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改资金总账失败,capital_ledger_id必须为整型"))
	}
	var accounting_date string
	if strings.Trim(c.Param("accounting_date"), " ") != "" {
		accounting_date = strings.Trim(c.Param("accounting_date"), " ")
	} else {
		accounting_date = ""
	}
	b, err := l.ModifyCapitalLedger(user_id, project_id, capital_ledger_id, amount, strings.Trim(c.Param("memo"), " "), accounting_date)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改资金总账失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改资金总账成功"))
})

var AddServiceAccount = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务账目失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务账目失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务账目失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务账目失败,project_id必须为整型"))
	}
	if c.Param("amount") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务账目失败,请传入amount!"))
	}
	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务账目失败,amount必须为数字"))
	}
	if c.Param("role_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务账目失败,请传入role_id!"))
	}
	role_id, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务账目失败,role_id必须为整型"))
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var accounting_date string
	if strings.Trim(c.Param("accounting_date"), " ") != "" {
		accounting_date = strings.Trim(c.Param("accounting_date"), " ")
	} else {
		accounting_date = time.Now().Format("2006-01-02")
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddServiceAccount(user_id, project_id, role_id, bind_type, amount, strings.Trim(c.Param("memo"), " "), accounting_date, img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加劳务账目失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加劳务账目成功"))
})

var ModifyServiceAccount = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改劳务账目失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改劳务账目失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改劳务账目失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改劳务账目失败,project_id必须为整型"))
	}
	if c.Param("amount") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改劳务账目失败,请传入amount!"))
	}
	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改劳务账目失败,amount必须为数字"))
	}
	if c.Param("service_account_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改劳务账目失败,请传入service_account_id!"))
	}
	service_account_id, err := strconv.ParseInt(c.Param("service_account_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改劳务账目失败,service_account_id必须为整型"))
	}

	b, err := l.ModifyServiceAccount(user_id, project_id, service_account_id, amount, strings.Trim(c.Param("memo"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改劳务账目失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改劳务账目成功"))
})

var AddPurchasingAccount = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,project_id必须为整型"))
	}
	if c.Param("purchasing_type") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入purchasing_type!"))
	}
	purchasing_type, err := strconv.ParseInt(c.Param("purchasing_type"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,purchasing_type必须为整型"))
	}
	var contract_id, payment_method, material_type int64
	var contract_amount, price float64
	if purchasing_type == 0 {
		if c.Param("contract_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入contract_id!"))
		}
		contract_id, err = strconv.ParseInt(c.Param("contract_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,contract_id必须为整型"))
		}
		if c.Param("payment_method") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入payment_method!"))
		}
		payment_method, err = strconv.ParseInt(c.Param("payment_method"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,payment_method必须为整型"))
		}
		if c.Param("contract_amount") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入contract_amount!"))
		}
		contract_amount, err = strconv.ParseFloat(c.Param("contract_amount"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,contract_amount必须为数字"))
		}
	} else {
		if c.Param("price") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入price!"))
		}
		price, err = strconv.ParseFloat(c.Param("price"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,price必须为数字"))
		}
		if strings.Trim(c.Param("material_type"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入material_type!"))
		}
		material_type, err = strconv.ParseInt(c.Param("material_type"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,material_type必须为整型"))
		}
		if strings.Trim(c.Param("material_name"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入material_name!"))
		}
	}

	if strings.Trim(c.Param("company_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入company_name!"))
	}

	if c.Param("amount") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入amount!"))
	}
	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,amount必须为数字"))
	}
	if c.Param("quantity") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入quantity!"))
	}
	quantity, err := strconv.ParseFloat(c.Param("quantity"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,quantity必须为数字"))
	}
	if c.Param("unit_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入quantity!"))
	}
	unit_id, err := strconv.ParseInt(c.Param("unit_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,unit_id必须为数字"))
	}
	if strings.Trim(c.Param("invoiced"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,请传入invoiced!"))
	}
	invoiced, err := strconv.ParseBool(strings.Trim(c.Param("invoiced"), " "))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加采购账目失败,invoiced必须为布尔型!"))
	}
	var img_list, payment_voucher, invoice_img []string
	if strings.Trim(c.Param("img_list"), " ") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	if strings.Trim(c.Param("payment_voucher"), " ") != "" {
		payment_voucher = strings.Split(c.Param("payment_voucher"), ",")
	}
	if strings.Trim(c.Param("invoice_img"), " ") != "" {
		invoice_img = strings.Split(c.Param("invoice_img"), ",")
	}
	var accounting_date string
	if strings.Trim(c.Param("accounting_date"), " ") != "" {
		accounting_date = strings.Trim(c.Param("accounting_date"), " ")
	} else {
		accounting_date = time.Now().Format("2006-01-02")
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddPurchasingAccount(user_id, project_id, bind_type, contract_amount, amount, quantity, price,
		strings.Trim(c.Param("company_name"), " "),
		strings.Trim(c.Param("material_name"), " "),
		strings.Trim(c.Param("memo"), " "), accounting_date, material_type, contract_id, payment_method, unit_id, invoiced,
		img_list, payment_voucher, invoice_img)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加采购账目失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加采购账目成功"))
})

var ModifyPurchasingAccount = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,project_id必须为整型"))
	}
	if c.Param("amount") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,请传入amount!"))
	}
	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,amount必须为数字"))
	}
	if c.Param("purchasing_account_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,请传入purchasing_account_id!"))
	}
	purchasing_account_id, err := strconv.ParseInt(c.Param("purchasing_account_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,purchasing_account_id必须为整型"))
	}

	b, err := l.ModifyPurchasingAccount(user_id, project_id, purchasing_account_id, amount, strings.Trim(c.Param("memo"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改采购账目失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改采购账目成功"))
})

var GetCapitalAccountingIndex = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改采购账目失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetCapitalAccountingIndex(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetCapitalLedgerIndex = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取资金总账页面数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取资金总账页面数据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取资金总账页面数据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取资金总账页面数据失败,project_id必须为整型"))
	}
	var begin_date, end_date time.Time
	if c.Param("begin_date") != "" {
		begin_date, err = time.Parse("2006-01-02", c.Param("begin_date"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取资金总账页面数据失败,begin_date需要遵循YYYY-MM-DD格式"))
		}
	} else {
		begin_date = time.Now().AddDate(-1, 0, 0)
	}
	if c.Param("end_date") != "" {
		end_date, err = time.Parse("2006-01-02", c.Param("end_date"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取资金总账页面数据失败,end_date需要遵循YYYY-MM-DD格式"))
		}
	} else {
		end_date = time.Now().AddDate(0, 0, 1)
	}
	if begin_date.Unix() > 0 {
		if end_date.Unix() < 0 {
			end_date = time.Now()
		}
		if end_date.Unix() < begin_date.Unix() {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取资金总账页面数据失败,end_date必须大于begin_date"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}

	data, err := l.GetCapitalLedgerIndex(user_id, project_id, bind_type, begin_date, end_date, strings.Trim(c.Param("keyword"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetServiceAccountIndex = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取劳务账目页面数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取劳务账目页面数据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取劳务账目页面数据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取劳务账目页面数据失败,project_id必须为整型"))
	}
	var begin_date, end_date time.Time
	if c.Param("begin_date") != "" {
		begin_date, err = time.Parse("2006-01-02", c.Param("begin_date"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取劳务账目页面数据失败,begin_date需要遵循YYYY-MM-DD格式"))
		}
	} else {
		begin_date = time.Now().AddDate(-1, 0, 0)
	}
	if c.Param("end_date") != "" {
		end_date, err = time.Parse("2006-01-02", c.Param("end_date"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取劳务账目页面数据失败,end_date需要遵循YYYY-MM-DD格式"))
		}
	} else {
		end_date = time.Now().AddDate(0, 0, 1)
	}
	if begin_date.Unix() > 0 {
		if end_date.Unix() < 0 {
			end_date = time.Now()
		}
		if end_date.Unix() < begin_date.Unix() {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取劳务账目页面数据失败,end_date必须大于begin_date"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetServiceAccountIndex(user_id, project_id, bind_type, begin_date, end_date, strings.Trim(c.Param("keyword"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetPurchasingAccountIndex = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取采购账目页面数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取采购账目页面数据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取采购账目页面数据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取采购账目页面数据失败,project_id必须为整型"))
	}
	var begin_date, end_date time.Time
	if c.Param("begin_date") != "" {
		begin_date, err = time.Parse("2006-01-02", c.Param("begin_date"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取采购账目页面数据失败,begin_date需要遵循YYYY-MM-DD格式"))
		}
	} else {
		begin_date = time.Now().AddDate(-1, 0, 0)
	}
	if c.Param("end_date") != "" {
		end_date, err = time.Parse("2006-01-02", c.Param("end_date"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取采购账目页面数据失败,end_date需要遵循YYYY-MM-DD格式"))
		}
	} else {
		end_date = time.Now().AddDate(0, 0, 1)
	}
	if begin_date.Unix() > 0 {
		if end_date.Unix() < 0 {
			end_date = time.Now()
		}
		if end_date.Unix() < begin_date.Unix() {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取采购账目页面数据失败,end_date必须大于begin_date"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetPurchasingAccountIndex(user_id, project_id, bind_type, begin_date, end_date, strings.Trim(c.Param("keyword"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetUnitList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取计量单位列表失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取计量单位列表失败,user_id必须为整型"))
	}
	data, err := l.GetUnitList(user_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetDocList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档列表失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档列表失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档列表失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档列表失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetDocList(user_id, project_id, bind_type, strings.Trim(c.Param("keyword"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetConstructionTeamList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工班组列表失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工班组列表失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工班组列表失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工班组列表失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetConstructionTeamList(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddProjectDoc = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,project_id必须为整型"))
	}
	if c.Param("doc_type") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,请传入doc_type!"))
	}
	doc_type, err := strconv.ParseInt(c.Param("doc_type"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,doc_type必须为整型"))
	}
	if strings.Trim(c.Param("company_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,请传入company_name!"))
	}
	if c.Param("contract_amount") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,请传入contract_amount!"))
	}
	contract_amount, err := strconv.ParseFloat(c.Param("contract_amount"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档失败,contract_amount必须为数字"))
	}
	var img_list []string
	if strings.Trim(c.Param("img_list"), " ") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddProjectDoc(user_id, project_id, doc_type, bind_type, strings.Trim(c.Param("company_name"), " "),
		strings.Trim(c.Param("doc_code"), " "), contract_amount, img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "新增项目文档失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "新增项目文档成功"))
})

var GetDocTypeList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档类型失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档类型失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档类型失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档类型失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetDocTypeList(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddDocType = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档类型失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档类型失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档类型失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档类型失败,project_id必须为整型"))
	}
	if strings.Trim(c.Param("name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增项目文档类型失败,请传入name!"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddDocType(user_id, project_id, bind_type, strings.Trim(c.Param("name"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "新增项目文档类型失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "新增项目文档类型成功"))
})

var ViewProjectDoc = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档失败,project_id必须为整型"))
	}
	if c.Param("doc_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档失败,请传入doc_id!"))
	}
	doc_id, err := strconv.ParseInt(c.Param("doc_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档失败,doc_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ViewProjectDoc(user_id, project_id, doc_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetMaterielCategory = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料类别失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料类别失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料类别失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料类别失败,project_id必须为整型"))
	}
	var parent_id int64
	if strings.Trim(c.Param("parent_id"), " ") == "" {
		parent_id = 0
	} else {
		parent_id, err = strconv.ParseInt(c.Param("parent_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料类别失败,parent_id必须为整型"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetMaterielCategory(user_id, project_id, parent_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetPurchasingDetail = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料采购明细失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料采购明细失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料采购明细失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料采购明细失败,project_id必须为整型"))
	}
	if strings.Trim(c.Param("company_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取物料采购明细失败,请传入company_name!"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}

	data, err := l.GetPurchasingDetail(user_id, project_id, bind_type, strings.Trim(c.Param("company_name"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))

})

var AddServiceRole = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务角色失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务角色失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务角色失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务角色失败,project_id必须为整型"))
	}
	if strings.Trim(c.Param("role_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务角色失败,请传入company_name!"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加劳务角色失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddServiceRole(user_id, project_id, strings.Trim(c.Param("role_name"), " "), bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加劳务角色失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加劳务角色成功"))
})

var DelDocType = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档分类失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档分类失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档分类失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档分类失败,project_id必须为整型"))
	}
	if c.Param("type_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档分类失败,请传入type_id!"))
	}
	type_id, err := strconv.ParseInt(c.Param("type_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档分类失败,type_id必须为整型"))
	}
	b, err := l.DelDocType(user_id, project_id, type_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除文档分类失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除文档分类成功"))
})

var GetMyOrderList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取订单列表失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取订单列表失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取订单列表失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取订单列表失败,company_id必须为整型"))
	}
	var page_num, count int
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.Atoi(c.Param("page_num"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.Atoi(c.Param("count"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败，count必须为整型"))
		}
	}
	data, totalCount, err := l.GetMyOrderList(user_id, company_id, page_num, count)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","total_count":%d,"data":%s}`, code, msg, totalCount, string(b)))
})

var ModifyOrderStatus = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单状态失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单状态失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单状态失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单状态失败,company_id必须为整型"))
	}
	if c.Param("order_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单状态失败,请传入order_id!"))
	}
	order_id, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单状态失败,order_id必须为整型"))
	}
	if c.Param("status") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单状态失败,请传入status!"))
	}
	status, err := strconv.ParseInt(c.Param("status"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单状态失败,status必须为整型"))
	}
	b, err := l.ModifyOrderStatus(user_id, company_id, order_id, status)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改订单状态失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改订单状态成功"))
})

var DrawCashRequest = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "公司申请提现失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "公司申请提现失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "公司申请提现失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "公司申请提现失败,company_id必须为整型"))
	}
	if c.Param("payment_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "公司申请提现失败,请传入payment_id!"))
	}
	payment_id, err := strconv.ParseInt(c.Param("payment_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "公司申请提现失败,payment_id必须为整型"))
	}
	if c.Param("amount") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "公司申请提现失败,请传入amount!"))
	}
	amount, err := strconv.ParseFloat(c.Param("amount"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "公司申请提现失败,amount必须为数字"))
	}

	b, err := l.DrawCashRequest(user_id, company_id, payment_id, amount, strings.Trim(c.Param("memo"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "申请提现失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "申请提现成功"))
})

var GetCompanyAccount = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户信息失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户信息失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户信息失败,company_id必须为整型"))
	}
	data, err := l.GetCompanyAccount(user_id, company_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetCompanyAccountDetail = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户明细失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户明细失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户明细失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户明细失败,company_id必须为整型"))
	}
	data, err := l.GetCompanyAccountDetail(user_id, company_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetCompanyAccountDetail2 = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户明细失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户明细失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户明细失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取公司账户明细失败,company_id必须为整型"))
	}
	var beginMonth, endMonth string
	if strings.Trim(c.Param("begin_time"), " ") != "" {
		beginMonth = strings.Trim(c.Param("begin_time"), " ")
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户账户余额详情失败,begin_time必须为YYYY-MM-DD格式"))
		}
	} else {
		beginMonth = "2017-01"
	}
	if strings.Trim(c.Param("end_time"), " ") != "" {
		endMonth = strings.Trim(c.Param("end_time"), " ")
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户账户余额详情失败,end_time必须为YYYY-MM-DD格式"))
		}
	} else {
		endMonth = time.Now().Format("2006-01")
	}
	data, err := l.GetCompanyAccountDetail2(user_id, company_id, beginMonth, endMonth)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddBankAccount = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加公司账户失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加公司账户失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加公司账户失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加公司账户失败,company_id必须为整型"))
	}
	if c.Param("pay_account") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加公司账户失败,请传入pay_account!"))
	}
	if c.Param("phone") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加公司账户失败,请传入phone!"))
	}
	if c.Param("validate_code") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加公司账户失败,请传入validate_code!"))
	}
	if c.Param("pay_type") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加公司账户失败,请传入pay_type!"))
	}
	pay_type, err := strconv.Atoi(c.Param("pay_type"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "pay_type必须为整型"))
	}
	b, err := l.AddBankAccount(user_id, company_id, c.Param("phone"), c.Param("pay_account"), c.Param("bank_name"), c.Param("validate_code"), pay_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))

})

var DeleteBankAccount = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "解绑银行卡失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "解绑银行卡失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加公司账户失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加公司账户失败,company_id必须为整型"))
	}
	if c.Param("payment_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "解绑银行卡失败,请传入payment_id!"))
	}
	payment_id, err := strconv.ParseInt(c.Param("payment_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "解绑银行卡失败,payment_id必须为整型"))
	}
	if c.Param("pay_pass") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "解绑银行卡失败,请传入pay_pass!"))
	}

	b, err := l.DeleteBankAccount(user_id, company_id, payment_id, c.Param("pay_pass"))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var SetPayPass = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "设置公司支付密码失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置公司支付密码失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置公司支付密码失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置公司支付密码失败,company_id必须为整型"))
	}
	if c.Param("pass") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "设置公司支付密码失败,请传入pass!"))
	}
	if c.Param("validate_code") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "设置公司支付密码失败,请传入validate_code!"))
	}
	b, err := l.SetPayPass(user_id, company_id, c.Param("validate_code"), c.Param("pass"))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var ModifyPayPass = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "修改公司支付密码失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改公司支付密码失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改公司支付密码失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改公司支付密码失败,company_id必须为整型"))
	}
	if c.Param("old_pass") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "修改公司支付密码失败,请传入old_pass!"))
	}
	if c.Param("new_pass") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "修改公司支付密码失败,请传入new_pass!"))
	}
	if c.Param("validate_code") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "修改公司支付密码失败,请传入validate_code!"))
	}
	b, err := l.ModifyPayPass(user_id, company_id, c.Param("validate_code"), c.Param("old_pass"), c.Param("new_pass"))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var ForgetPayPass = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "忘记公司支付密码失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "忘记公司支付密码失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "忘记公司支付密码失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "忘记公司支付密码失败,company_id必须为整型"))
	}
	if c.Param("validate_code") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "忘记公司支付密码失败,请传入validate_code!"))
	}
	if c.Param("new_pay_pass") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "忘记公司支付密码失败,请传入validate_code!"))
	}
	b, err := l.ForgetPayPass(user_id, company_id, c.Param("validate_code"), c.Param("new_pay_pass"))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}

	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s", "setup":%t}`, 0, "success", b))
})

//获取绑定的支付信息
var GetCompanyPaymentInfo = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取企业支付信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取企业支付信息失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取企业支付信息失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取企业支付信息失败,company_id必须为整型"))
	}
	pi, err := l.GetCompanyPaymentInfo(user_id, company_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, err := json.Marshal(pi)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetBuildMyInfo = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取我的信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取我的信息失败,user_id必须为整型"))
	}
	if c.Param("company_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取我的信息失败,请传入company_id!"))
	}
	company_id, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取我的信息失败,company_id必须为整型"))
	}
	data, err := l.GetBuildMyInfo(user_id, company_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetMyBuildQrcode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取自己的二维码以及用户信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取自己的二维码以及用户信息失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取自己的二维码以及用户信息失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取自己的二维码以及用户信息失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取自己的二维码以及用户信息失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetMyBuildQrcode(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetUserBuildQrcode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取用户二维码失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户二维码失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取用户二维码失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户二维码失败,project_id必须为整型"))
	}
	if c.Param("object_user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取用户二维码失败,请传入object_user_id!"))
	}
	object_user_id, err := strconv.ParseInt(c.Param("object_user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户二维码失败,object_user_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取用户二维码失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetUserBuildQrcode(user_id, project_id, object_user_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ViewSecurityCheck = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "查看安全检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,project_id必须为整型"))
	}
	if c.Param("sc_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "查看安全检查失败,请传入sc_id!"))
	}
	sc_id, err := strconv.ParseInt(c.Param("sc_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,sc_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ViewSecurityCheck(user_id, project_id, sc_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ViewQualityTesting = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "查看质量检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,project_id必须为整型"))
	}
	if c.Param("qt_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "查看质量检查失败,请传入qt_id!"))
	}
	qt_id, err := strconv.ParseInt(c.Param("qt_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,qt_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ViewQualityTesting(user_id, project_id, qt_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ViewFieldInspection = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "查看协作巡查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,project_id必须为整型"))
	}
	if c.Param("fi_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "查看协作巡查失败,请传入fi_id!"))
	}
	fi_id, err := strconv.ParseInt(c.Param("fi_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,fi_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ViewFieldInspection(user_id, project_id, fi_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddSecurityCheckComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加安全检查评论失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查评论失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查评论失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查评论失败,project_id必须为整型"))
	}
	if c.Param("sc_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加安全检查评论失败,请传入sc_id!"))
	}
	sc_id, err := strconv.ParseInt(c.Param("sc_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查评论失败,sc_id必须为整型"))
	}
	if c.Param("reply_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加安全检查评论失败,请传入reply_id!"))
	}
	reply_id, err := strconv.ParseInt(c.Param("reply_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查评论失败,reply_id必须为整型"))
	}
	is_rectification := false
	if c.Param("is_rectification") != "" {
		is_rectification, err = strconv.ParseBool(c.Param("is_rectification"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查评论失败,is_rectification必须为布尔型"))
		}
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加安全检查评论失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	var param string
	if strings.Trim(c.Param("param"), " ") != "" {
		param = strings.Trim(c.Param("param"), " ")
	}
	b, err := l.AddSecurityCheckComment(user_id, project_id, sc_id, reply_id, bind_type, strings.Trim(c.Param("content"), " "), is_rectification, img_list, param)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var AddQualityTestingComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加质量检查评论失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查评论失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查评论失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查评论失败,project_id必须为整型"))
	}
	if c.Param("qt_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加质量检查评论失败,请传入qt_id!"))
	}
	qt_id, err := strconv.ParseInt(c.Param("qt_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查评论失败,qt_id必须为整型"))
	}
	if c.Param("reply_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加质量检查评论失败,请传入reply_id!"))
	}
	reply_id, err := strconv.ParseInt(c.Param("reply_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查评论失败,reply_id必须为整型"))
	}
	is_rectification := false
	if c.Param("is_rectification") != "" {
		is_rectification, err = strconv.ParseBool(c.Param("is_rectification"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查评论失败,is_rectification必须为布尔型"))
		}
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加质量检查评论失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	var param string
	if strings.Trim(c.Param("param"), " ") != "" {
		param = strings.Trim(c.Param("param"), " ")
	}
	b, err := l.AddQualityTestingComment(user_id, project_id, qt_id, reply_id, bind_type, strings.Trim(c.Param("content"), " "), is_rectification, img_list, param)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var AddFieldInspectionComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加协作巡查评论失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查评论失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查评论失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查评论失败,project_id必须为整型"))
	}
	if c.Param("fi_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加协作巡查评论失败,请传入fi_id!"))
	}
	fi_id, err := strconv.ParseInt(c.Param("fi_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查评论失败,fi_id必须为整型"))
	}
	if c.Param("reply_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "增加协作巡查评论失败,请传入reply_id!"))
	}
	reply_id, err := strconv.ParseInt(c.Param("reply_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查评论失败,reply_id必须为整型"))
	}
	is_rectification := false
	if c.Param("is_rectification") != "" {
		is_rectification, err = strconv.ParseBool(c.Param("is_rectification"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查评论失败,is_rectification必须为布尔型"))
		}
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加协作巡查评论失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	var param string
	if strings.Trim(c.Param("param"), " ") != "" {
		param = strings.Trim(c.Param("param"), " ")
	}
	b, err := l.AddFieldInspectionComment(user_id, project_id, fi_id, reply_id, bind_type, strings.Trim(c.Param("content"), " "), is_rectification, img_list, param)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var DelSecurityCheckComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "删除安全检查评论失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除安全检查评论失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除安全检查评论失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除安全检查评论失败,project_id必须为整型"))
	}
	if c.Param("sc_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "删除安全检查评论失败,请传入sc_id!"))
	}
	sc_id, err := strconv.ParseInt(c.Param("sc_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除安全检查评论失败,sc_id必须为整型"))
	}

	b, err := l.DelSecurityCheckComment(user_id, project_id, sc_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var DelQualityTestingComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "删除质量检查评论失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除质量检查评论失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除质量检查评论失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除质量检查评论失败,project_id必须为整型"))
	}
	if c.Param("qt_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "删除质量检查评论失败,请传入qt_id!"))
	}
	qt_id, err := strconv.ParseInt(c.Param("qt_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除质量检查评论失败,qt_id必须为整型"))
	}
	b, err := l.DelQualityTestingComment(user_id, project_id, qt_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var DelFieldInspectionComment = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "删除协作巡查评论失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除协作巡查评论失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除协作巡查评论失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除协作巡查评论失败,project_id必须为整型"))
	}
	if c.Param("fi_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "删除协作巡查评论失败,请传入fi_id!"))
	}
	fi_id, err := strconv.ParseInt(c.Param("fi_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除协作巡查评论失败,fi_id必须为整型"))
	}
	b, err := l.DelFieldInspectionComment(user_id, project_id, fi_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var LaunchTeamVisa = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "发起班组签证失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,project_id必须为整型"))
	}
	if c.Param("assign_team") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,请传入assign_team!"))
	}
	assign_team, err := strconv.ParseInt(c.Param("assign_team"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,assign_team必须为整型"))
	}
	if c.Param("visa_type") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,请传入visa_type!"))
	}
	visa_type, err := strconv.ParseInt(c.Param("visa_type"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,visa_type必须为整型"))
	}
	if visa_type > 4 || visa_type < 1 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,签证类型不正确"))
	}
	if c.Param("visa_reason") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,请传入visa_reason!"))
	}
	var main_send_user, cc_user, img_list []string // []int64
	if c.Param("main_send_user") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,请传入main_send_user!"))
	}
	main_send_user = strings.Split(c.Param("main_send_user"), ",")
	if strings.Trim(c.Param("cc_send_user"), " ") != "" {
		cc_user = strings.Split(c.Param("cc_send_user"), ",")
	}
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var visa_date string
	if strings.Trim(c.Param("visa_date"), " ") != "" {
		visa_date = strings.Trim(c.Param("visa_date"), " ")
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起班组签证失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.LaunchTeamVisa(user_id, project_id, assign_team, visa_type, bind_type, strings.Trim(c.Param("visa_reason"), " "),
		strings.Trim(c.Param("visa_desc"), " "), visa_date, main_send_user, cc_user, img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var GetTeamVisaList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取班组签证列表失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败,project_id必须为整型"))
	}
	var assign_team, approval_status, related_me int64
	if c.Param("assign_team") != "" {
		assign_team, err = strconv.ParseInt(c.Param("assign_team"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败,assign_team必须为整型"))
		}
	}
	if c.Param("approval_status") != "" {
		approval_status, err = strconv.ParseInt(c.Param("approval_status"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败,approval_status必须为整型"))
		}
	}
	if c.Param("related_me") != "" {
		related_me, err = strconv.ParseInt(c.Param("related_me"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败,related_me必须为整型"))
		}
	}
	var page_num, count int
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.Atoi(c.Param("page_num"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.Atoi(c.Param("count"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败，count必须为整型"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证列表失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, tcount, err := l.GetTeamVisaList(user_id, project_id, assign_team, approval_status, related_me, page_num, count, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s,"total_count":%d}`, code, msg, string(b), tcount))
})

var GetTeamVisaInfo = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取班组签证信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证信息失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证信息失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证信息失败,project_id必须为整型"))
	}
	if c.Param("tv_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证信息失败,请传入tv_id!"))
	}
	tv_id, err := strconv.ParseInt(c.Param("tv_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证信息失败,tv_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组签证信息失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetTeamVisaInfo(user_id, project_id, tv_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var HandlingTeamVisa = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "处理班组签证失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "处理班组签证失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "处理班组签证失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "处理班组签证失败,project_id必须为整型"))
	}
	if c.Param("tv_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "处理班组签证失败,请传入user_id!"))
	}
	tv_id, err := strconv.ParseInt(c.Param("tv_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "处理班组签证失败,user_id必须为整型"))
	}
	if c.Param("agree_refuse") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "处理班组签证失败,请传入agree_refuse!"))
	}
	agree_refuse, err := strconv.ParseBool(c.Param("agree_refuse"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "处理班组签证失败,agree_refuse必须为布尔型"))
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	var amount float64
	if c.Param("amount") != "" {
		amount, err = strconv.ParseFloat(c.Param("amount"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "处理班组签证失败,amount必须为浮点或者数值型"))
		}
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "处理班组签证失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.HandlingTeamVisa(user_id, project_id, tv_id, agree_refuse, bind_type, c.Param("reason"), amount, img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var LaunchWagesManage = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "发起工资管理失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,project_id必须为整型"))
	}
	if c.Param("wages_type") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,请传入wages_type!"))
	}
	wages_type, err := strconv.Atoi(c.Param("wages_type"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,wages_type必须为整型"))
	}
	if c.Param("wages_month") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,请传入wages_month!"))
	}
	wages_month, err := strconv.Atoi(c.Param("wages_month"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,wages_month必须为整型"))
	}
	if wages_month > 12 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,month必须小于12"))
	}
	if c.Param("total_wages") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "发起工资管理失败,请传入total_wages!"))
	}
	total_wages, err := strconv.ParseFloat(c.Param("total_wages"), 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,total_wages必须为数值型"))
	}
	var wages_year int
	if c.Param("wages_year") != "" {
		wages_year, err = strconv.Atoi(c.Param("wages_year"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,wages_year必须为整型"))
		}
	} else {
		wages_year = time.Now().Local().Year()
	}
	if wages_year-1 > time.Now().Local().Year() {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,不能发放大于当前年份的工资"))
	}
	if c.Param("team_wages_alloc") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "发起工资管理失败,请传入team_wages_alloc!"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发起工资管理失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.LaunchWagesManage(user_id, project_id, wages_type, wages_month, wages_year, bind_type, total_wages, c.Param("team_wages_alloc"))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var GetWagesManageList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取工资管理列表失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工资管理列表失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工资管理列表失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工资管理列表失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetWagesManageList(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetWagesManageDetail = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取工资管理详情失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工资管理详情失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工资管理详情失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工资管理详情失败,project_id必须为整型"))
	}
	if c.Param("wm_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工资管理详情失败,请传入wm_id!"))
	}
	wm_id, err := strconv.ParseInt(c.Param("wm_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取工资管理详情失败,wm_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetWagesManageDetail(user_id, project_id, wm_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetWagesTeamDetail = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取班组工资管理详情失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,project_id必须为整型"))
	}
	if c.Param("wm_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,请传入wm_id!"))
	}
	wm_id, err := strconv.ParseInt(c.Param("wm_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,wm_id必须为整型"))
	}
	if c.Param("tm_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,请传入tm_id!"))
	}
	tm_id, err := strconv.ParseInt(c.Param("tm_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,tm_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组工资管理详情失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetWagesTeamDetail(user_id, project_id, wm_id, tm_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetTeamUserList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取班组成员信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组成员信息失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组成员信息失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组成员信息失败,project_id必须为整型"))
	}
	if c.Param("role_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组成员信息失败,请传入role_id!"))
	}
	role_id, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组成员信息失败,role_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取班组成员信息失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetTeamUserList(user_id, project_id, role_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ModifyWagesManageStatus = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "更新工资表状态失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "更新工资表状态失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "更新工资表状态失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "更新工资表状态失败,project_id必须为整型"))
	}
	if c.Param("wm_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "更新工资表状态失败,请传入wm_id!"))
	}
	wm_id, err := strconv.ParseInt(c.Param("wm_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "更新工资表状态失败,wm_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "更新工资表状态失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.ModifyWagesManageStatus(user_id, project_id, wm_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var PerfectTeamPayroll = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "完善班组工资表失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "完善班组工资表失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "完善班组工资表失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "完善班组工资表失败,project_id必须为整型"))
	}
	if c.Param("wm_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "完善班组工资表失败,请传入wm_id!"))
	}
	wm_id, err := strconv.ParseInt(c.Param("wm_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "完善班组工资表失败,wm_id必须为整型"))
	}
	if c.Param("role_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "完善班组工资表失败,请传入role_id!"))
	}
	role_id, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "完善班组工资表失败,role_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "更新工资表状态失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.PerfectTeamPayroll(user_id, project_id, wm_id, role_id, bind_type, c.Param("team_wages_alloc"))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var DetermineRoles = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "判断用户在项目中的角色失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "判断用户在项目中的角色失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "判断用户在项目中的角色失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "判断用户在项目中的角色失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.DetermineRoles(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetMyProject = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "判断用户在项目中的角色失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "判断用户在项目中的角色失败,user_id必须为整型"))
	}
	data, err := l.GetMyProject(user_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ScanBuildQrcode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "扫一扫二维码获取信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,project_id必须为整型"))
	}
	if strings.ToLower(strings.Trim(c.Param("scene"), " ")) != "get_qrcode_detail" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "功能尚未开放，请等待下一次更新"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ScanBuildQrcode(user_id, project_id, bind_type, c.Param("phone"), strings.Trim(c.Param("scene"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if len(data) > 0 {
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, data))
	} else {
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, "功能尚未开放，请等待下一次更新"))
	}

})

var PayrollVoucherUpload = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "上传工资发放凭证失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传工资发放凭证失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传工资发放凭证失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传工资发放凭证失败,project_id必须为整型"))
	}
	if c.Param("wm_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传工资发放凭证失败,请传入wm_id!"))
	}
	wm_id, err := strconv.ParseInt(c.Param("wm_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传工资发放凭证失败,wm_id必须为整型"))
	}

	if c.Param("payroll_voucher") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传工资发放凭证失败,请传入payroll_voucher!"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传工资发放凭证失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.PayrollVoucherUpload(user_id, project_id, wm_id, bind_type, strings.Trim(c.Param("payroll_voucher"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 500, err))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var GetWagesIsPerfect = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("wi_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取项目施工日志失败,请传入user_id!"))
	}
	wi_id, err := strconv.ParseInt(c.Param("wi_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志失败,user_id必须为整型"))
	}
	b := l.GetWagesIsPerfect(wi_id)

	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "false"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 0, "success"))
})

var GetProjectBuildingLog = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取项目施工日志失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetProjectBuildingLog(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddProjectBuildingLog = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "发表项目施工日志失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发表项目施工日志失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发表项目施工日志失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发表项目施工日志失败,project_id必须为整型"))
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	if strings.Trim(c.Param("log_context"), " ") == "" && len(img_list) == 0 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发表项目施工日志失败,不能发表空日志!"))
	}
	//var node_id int64 = -1
	//if strings.Trim(c.Param("node_id"), " ") != "" {
	//	node_id, err = strconv.ParseInt(c.Param("node_id"), 10, 64)
	//	if err != nil {
	//		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发表项目施工日志失败,node_id必须为整型"))
	//	}
	//}
	if c.Param("node_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发表项目施工日志失败,请传入node_id!"))
	}
	node_id, err := strconv.ParseInt(c.Param("node_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发表项目施工日志失败,node_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddProjectBuildingLog(user_id, project_id, node_id, bind_type, strings.Trim(c.Param("log_context"), " "), img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "发表项目施工日志失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "发表项目施工日志成功"))
})

var GetProjectBuildingLogNode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "获取项目施工日志节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志节点失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志节点失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetProjectBuildingLogNode(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var EditBuildingLogNode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "编辑项目施工日志节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑项目施工日志节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑项目施工日志节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑项目施工日志节点失败,project_id必须为整型"))
	}
	if c.Param("log_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑项目施工日志节点失败,请传入log_id!"))
	}
	log_id, err := strconv.ParseInt(c.Param("log_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑项目施工日志节点失败,log_id必须为整型"))
	}
	var node_id int64 = -1
	if strings.Trim(c.Param("node_id"), " ") != "" {
		node_id, err = strconv.ParseInt(c.Param("node_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑项目施工日志节点失败,node_id必须为整型"))
		}
	}
	var img_list []string
	if c.Param("img_list") != "" {
		img_list = strings.Split(c.Param("img_list"), ",")
	}
	if strings.Trim(c.Param("log_context"), " ") == "" && len(img_list) == 0 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑项目施工日志节点失败,不能发表空日志!"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目施工日志节点失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.EditProjectBuildingLog(user_id, project_id, log_id, node_id, bind_type, strings.Trim(c.Param("log_context"), " "), img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "编辑项目施工日志节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "编辑项目施工日志成功"))
})

var DeleteBuildingLogNode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "删除项目施工日志节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除项目施工日志节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除项目施工日志节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除项目施工日志节点失败,project_id必须为整型"))
	}
	if c.Param("log_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除项目施工日志节点失败,请传入log_id!"))
	}
	log_id, err := strconv.ParseInt(c.Param("log_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除项目施工日志节点失败,log_id必须为整型"))
	}
	b, err := l.DeleteBuildingLogNode(user_id, project_id, log_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除项目施工日志节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除项目施工日志成功"))
})

var GetProjectManager = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目的项目经理失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目的项目经理失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetProjectManager(project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

/**
//测试项目的建筑日志初始化
var InitProjectBuildingLog = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "初始化项目施工日志失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "初始化项目施工日志失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "初始化项目施工日志失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "初始化项目施工日志失败,project_id必须为整型"))
	}
	l.InitProjectBuildingLog(project_id)
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","user_id":%d}`, 200, "初始化项目施工日志成功", user_id))
})
*/

var ViewBillManagement = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理数据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理数据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理数据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理数据失败,project_id必须为整型"))
	}
	var begin_date, end_date time.Time
	if c.Param("begin_date") != "" {
		begin_date, err = time.Parse("2006-01-02", c.Param("begin_date"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理数据失败,begin_date需要遵循YYYY-MM-DD格式"))
		}
	} else {
		begin_date = time.Now().AddDate(-1, 0, 0)
	}
	if c.Param("end_date") != "" {
		end_date, err = time.Parse("2006-01-02", c.Param("end_date"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理数据失败,end_date需要遵循YYYY-MM-DD格式"))
		}
	} else {
		end_date = time.Now().AddDate(0, 0, 1)
	}
	if begin_date.Unix() > 0 {
		if end_date.Unix() < 0 {
			end_date = time.Now()
		}
		if end_date.Unix() < begin_date.Unix() {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理数据失败,end_date必须大于begin_date"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ViewBillManagement(user_id, project_id, bind_type, begin_date, end_date, strings.Trim(c.Param("keyword"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ViewBillDetail = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理明细失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理明细失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理明细失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理明细失败,project_id必须为整型"))
	}
	if strings.Trim(c.Param("company_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取票据管理明细失败,请传入company_name!"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}

	data, err := l.ViewBillDetail(user_id, project_id, bind_type, strings.Trim(c.Param("company_name"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var BillVerify = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,project_id必须为整型"))
	}

	if c.Param("purchasing_account_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,请传入purchasing_account_id!"))
	}
	purchasing_account_id, err := strconv.ParseInt(c.Param("purchasing_account_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,purchasing_account_id必须为整型"))
	}
	if c.Param("verify") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,请传入verify!"))
	}
	verify, err := strconv.ParseBool(strings.Trim(c.Param("verify"), " "))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,verify必须为布尔型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "核实票据失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.BillVerify(user_id, project_id, purchasing_account_id, bind_type, verify, strings.Trim(c.Param("memo"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "核实票据失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "核实票据成功"))
})

var ResubmitBill = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重新提交票据失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重新提交票据失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重新提交票据失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重新提交票据失败,project_id必须为整型"))
	}

	if c.Param("purchasing_account_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重新提交票据失败,请传入purchasing_account_id!"))
	}
	purchasing_account_id, err := strconv.ParseInt(c.Param("purchasing_account_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重新提交票据失败,purchasing_account_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重新提交票据失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	if c.Param("img_list") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重新提交票据失败,请传入img_list!"))
	}
	img_list := strings.Split(c.Param("img_list"), ",")
	b, err := l.ResubmitBill(user_id, project_id, purchasing_account_id, bind_type, img_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "重新提交票据失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "重新提交票据成功"))
})

var DeletePic = faygo.HandlerFunc(func(c *faygo.Context) error {
	if strings.Trim(c.Param("pic_url"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除图片失败,请传入pic_url!"))
	}
	b, err := l.DeletePic(strings.Trim(c.Param("pic_url"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除图片失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除图片成功"))
})

var GetConstructionPlan = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划失败,project_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetConstructionPlan(user_id, project_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var AddConstructionPlanNode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,project_id必须为整型"))
	}
	if c.Param("item_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,请传入item_id!"))
	}
	item_id, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,item_id必须为整型"))
	}
	if strings.Trim(c.Param("node_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,请传入node_name!"))
	}
	//if strings.Trim(c.Param("start_date")," ") == "" {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,请传入start_date!"))
	//}
	//if strings.Trim(c.Param("end_date")," ") == "" {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,请传入start_date!"))
	//}
	var current_status int
	if strings.Trim(c.Param("current_status"), " ") != "" {
		current_status, err = strconv.Atoi(strings.Trim(c.Param("current_status"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划过程子节点失败,current_status必须为整型"))
		}
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddConstructionPlanNode(user_id, project_id, item_id, current_status, bind_type, strings.Trim(c.Param("node_name"), " "),
		strings.Trim(c.Param("start_date"), " "), strings.Trim(c.Param("end_date"), " "),
		strings.Trim(c.Param("desc"), " "), strings.Trim(c.Param("photo_tips"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加施工计划过程子节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加施工计划过程子节点成功"))
})

var AddConstructionPlanItem = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,project_id必须为整型"))
	}
	if c.Param("parent_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,请传入parent_id!"))
	}
	parent_id, err := strconv.ParseInt(c.Param("parent_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,parent_id必须为整型"))
	}
	if strings.Trim(c.Param("item_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,请传入item_name!"))
	}
	if strings.Trim(c.Param("start_date"), " ") == "" && parent_id > 0 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,请传入start_date!"))
	}
	if strings.Trim(c.Param("end_date"), " ") == "" && parent_id > 0 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加施工计划子节点失败,请传入end_date!"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.AddConstructionPlanItem(user_id, project_id, parent_id, bind_type, strings.Trim(c.Param("item_name"), " "),
		strings.Trim(c.Param("start_date"), " "), strings.Trim(c.Param("end_date"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加施工计划子节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加施工计划子节点成功"))
})

var DelConstructionPlanItem = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,project_id必须为整型"))
	}
	if c.Param("item_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,请传入item_id!"))
	}
	item_id, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,item_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.DelConstructionPlanItem(user_id, project_id, item_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除施工计划子节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除施工计划子节点成功"))
})

var DelConstructionPlanNode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划过程子节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划过程子节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划过程子节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划过程子节点失败,project_id必须为整型"))
	}
	if c.Param("node_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划过程子节点失败,请传入node_id!"))
	}
	node_id, err := strconv.ParseInt(c.Param("node_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划过程子节点失败,node_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划过程子节点失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.DelConstructionPlanNode(user_id, project_id, node_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除施工计划过程子节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除施工计划过程子节点成功"))
})

var EditConstructionPlanItem = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,project_id必须为整型"))
	}
	if c.Param("parent_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,请传入parent_id!"))
	}
	parent_id, err := strconv.ParseInt(c.Param("parent_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,parent_id必须为整型"))
	}
	if c.Param("item_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,请传入item_id!"))
	}
	item_id, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除施工计划子节点失败,item_id必须为整型"))
	}
	if strings.Trim(c.Param("item_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,请传入start_date!"))
	}
	if strings.Trim(c.Param("start_date"), " ") == "" && parent_id > 0 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,请传入start_date!"))
	}
	if strings.Trim(c.Param("end_date"), " ") == "" && parent_id > 0 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划子节点失败,请传入start_date!"))
	}
	var seq int
	if strings.Trim(c.Param("seq"), " ") != "" {
		seq, err = strconv.Atoi(strings.Trim(c.Param("seq"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,seq必须为整型"))
		}
	}
	b, err := l.EditConstructionPlanItem(user_id, project_id, item_id, parent_id, seq, strings.Trim(c.Param("item_name"), " "),
		strings.Trim(c.Param("start_date"), " "), strings.Trim(c.Param("end_date"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "编辑施工计划子节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "编辑施工计划子节点成功"))
})

var EditConstructionPlanNode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,project_id必须为整型"))
	}
	if c.Param("item_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,请传入item_id!"))
	}
	item_id, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,item_id必须为整型"))
	}
	if c.Param("node_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,请传入node_id!"))
	}
	node_id, err := strconv.ParseInt(c.Param("node_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,node_id必须为整型"))
	}
	if strings.Trim(c.Param("node_name"), " ") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,请传入node_name!"))
	}
	//if strings.Trim(c.Param("start_date")," ") == "" {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,请传入start_date!"))
	//}
	//if strings.Trim(c.Param("end_date")," ") == "" {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,请传入start_date!"))
	//}
	var current_status, seq int
	if strings.Trim(c.Param("current_status"), " ") != "" {
		current_status, err = strconv.Atoi(strings.Trim(c.Param("current_status"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,current_status必须为整型"))
		}
	}
	if strings.Trim(c.Param("seq"), " ") != "" {
		seq, err = strconv.Atoi(strings.Trim(c.Param("seq"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,seq必须为整型"))
		}
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑施工计划过程子节点失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.EditConstructionPlanNode(user_id, project_id, item_id, node_id, current_status, seq, bind_type, strings.Trim(c.Param("node_name"), " "),
		strings.Trim(c.Param("start_date"), " "), strings.Trim(c.Param("end_date"), " "),
		strings.Trim(c.Param("desc"), " "), strings.Trim(c.Param("photo_tips"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "编辑施工计划过程子节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "编辑施工计划过程子节点成功"))
})

var GetConstructionPlanNode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,project_id必须为整型"))
	}
	if c.Param("node_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,请传入node_id!"))
	}
	node_id, err := strconv.ParseInt(c.Param("node_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,node_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetConstructionPlanNode(user_id, project_id, node_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetConstructionPlanItem = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,project_id必须为整型"))
	}
	if c.Param("item_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,请传入item_id!"))
	}
	item_id, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取施工计划过程子节点信息失败,item_id必须为整型"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.GetConstructionPlanItem(user_id, project_id, item_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var BatchAddPlanNode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "批量增加施工计划过程子节点失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "批量增加施工计划过程子节点失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "批量增加施工计划过程子节点失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "批量增加施工计划过程子节点失败,project_id必须为整型"))
	}
	if c.Param("item_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "批量增加施工计划过程子节点失败,请传入item_id!"))
	}
	item_id, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "批量增加施工计划过程子节点失败,item_id必须为整型"))
	}
	var node_list []string
	if c.Param("node_list") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "批量增加施工计划过程子节点失败,请传入node_list!"))
	}
	node_list = u.RemoveDuplicatesAndEmpty(strings.Split(c.Param("node_list"), ","))
	if len(node_list) < 1 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "批量增加施工计划过程子节点失败,请传入node_list!"))
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	b, err := l.BatchAddPlanNode(user_id, project_id, item_id, bind_type, node_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "批量增加施工计划过程子节点失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "批量增加施工计划过程子节点成功"))
})

var GetDevHouseList = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取绑定的开发商户型列表失败,请传入user_id!"))
		}
		user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取绑定的开发商户型列表失败,user_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取绑定的开发商户型列表失败,请传入project_id!"))
		}
		project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取绑定的开发商户型列表失败,project_id必须为整型"))
		}
		var bindType int
		if c.Param("bind_type") != "" {
			bindType, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取绑定的开发商户型列表失败,bind_type必须为整型!"))
			}
		} else {
			bindType = 1
		}

		var page_num, page_size int64
		if c.Param("page_num") == "" {
			page_num = 1
		} else {
			page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取绑定的开发商户型列表失败，page_num必须为整型"))
			}
		}
		if c.Param("page_size") == "" {
			page_size = 20
		} else {
			page_size, err = strconv.ParseInt(c.Param("page_size"), 10, 64)
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取绑定的开发商户型列表失败，page_size必须为整型"))
			}
		}
		data, pjname, err := l.GetDevHouseList(user_id, project_id, bindType, int(page_size), int(page_num-1)*int(page_size))
		b, err := json.Marshal(data)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","project_name":"%s","data":%s}`, code, msg, pjname, string(b)))
	}),
	"get_dev_house_list",
	"获取绑定的开发商户型列表",
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(6), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	faygo.ParamInfo{
		Name:     "project_id",
		In:       "query",
		Required: true,
		Model:    int(6), // API文档中显示的参数默认值
		Desc:     "项目编号",
	},
	faygo.ParamInfo{
		Name:     "bind_type",
		In:       "query",
		Required: false,
		Model:    int(2), // API文档中显示的参数默认值
		Desc:     "项目类型，1：建筑类;2:装修类",
	},
	faygo.ParamInfo{
		Name:     "page_num",
		In:       "query",
		Required: false,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "当前页码",
	},
	faygo.ParamInfo{
		Name:     "page_size",
		In:       "query",
		Required: false,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "返回数据大小",
	},
)

var GetChartGroupList = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取加入的群信息失败,请传入user_id!"))
		}
		user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取加入的群信息失败,user_id必须为整型"))
		}
		data, err := l.GetChartGroupList(user_id)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}), "get_chart_group_list", "获取加入的聊天群",
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(6), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
)

var CreateChartGroup = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建聊天群失败,请传入user_id!"))
		}
		user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建聊天群失败,user_id必须为整型"))
		}
		//if c.Param("project_id") == "" {
		//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建聊天群失败,请传入project_id!"))
		//}
		//project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		//if err != nil {
		//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建聊天群失败,project_id必须为整型"))
		//}
		var user_list []int64
		if c.Param("user_list") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建聊天群失败,请传入user_list!"))
		}
		user_list, err = u.StringToIntArray(c.Param("user_list"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建聊天群失败,user_list必须为整型数组"))
		}
		var app_name string
		if strings.Trim(c.Param("group_name"), " ") != "" {
			app_name = strings.Trim(c.Param("group_name"), " ")
		} else {
			app_name = "com.dingfang.dfjianzhuang"
		}
		b, groupId, err := l.CreateChartGroup(user_id, -1, strings.Trim(c.Param("group_name"), " "), app_name, user_list)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "创建聊天群失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","group_id":%d}`, code, "创建聊天群成功", groupId))
	}),
	"创建聊天群 create_chart_group",
	"服务器端返回的是否创建成功的信息",
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(6), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	faygo.ParamInfo{
		Name:     "user_list",
		In:       "query",
		Required: true,
		Model:    string("2,7"), // API文档中显示的参数默认值
		Desc:     "参加群聊的用户编号",
	},
	faygo.ParamInfo{
		Name:     "group_name",
		In:       "query",
		Required: false,
		Model:    string("xxx聊天群"), // API文档中显示的参数默认值
		Desc:     "参加群名称",
	},
	faygo.ParamInfo{
		Name:     "app_name",
		In:       "query",
		Required: false,
		Model:    string("com.dingfang.dfjianzhuang"), // API文档中显示的参数默认值
		Desc:     "当前应用的包名",
	},
)

var AddChartGroup = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "将用户拉入聊天群失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "将用户拉入聊天群失败,user_id必须为整型"))
	}
	if c.Param("object_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "将用户拉入聊天群失败,请传入object_id!"))
	}
	//object_id, err := strconv.ParseInt(c.Param("object_id"), 10, 64)
	//if err != nil {
	//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "将用户拉入聊天群失败,object_id必须为整型"))
	//}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "将用户拉入聊天群失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "将用户拉入聊天群失败,group_id必须为整型"))
	}
	object_arr := strings.Split(strings.Trim(c.Param("object_id"), " "), ",")
	if len(object_arr) < 1 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "将用户拉入聊天群失败,object_id值错误"))
	}
	b, err := l.AddChartGroup(user_id, group_id, object_arr)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "将用户拉入聊天群失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "将用户拉入聊天群成功"))
})

var QuitChartGroup = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "退出聊天群失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "退出聊天群失败,user_id必须为整型"))
	}

	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "退出聊天群失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "退出聊天群失败,group_id必须为整型"))
	}
	b, err := l.QuitChartGroup(user_id, group_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "退出聊天群失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "退出聊天群失败成功"))
})

var SetChartGroupInfo = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群信息失败,user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群信息失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群信息失败,group_id必须为整型"))
	}
	if c.Param("join_sign") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群信息失败,请传入join_sign!"))
	}
	join_sign, err := strconv.ParseBool(c.Param("join_sign"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群信息失败,join_sign必须为布尔型"))
	}
	if c.Param("view_his_chart") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群信息失败,请传入view_his_chart!"))
	}
	view_his_chart, err := strconv.ParseBool(c.Param("view_his_chart"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群信息失败,view_his_chart必须为布尔型"))
	}
	var group_photo string
	group_photo = strings.Trim(c.Param("group_photo"), " ")
	b, err := l.SetChartGroupInfo(user_id, group_id, join_sign, view_his_chart, strings.Trim(c.Param("group_name"), " "), group_photo)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "设置聊天群信息失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "设置聊天群信息成功"))
})

var SetChartGroupPersonalInfo = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群个人信息失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群个人信息失败,user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群个人信息失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群个人信息失败,group_id必须为整型"))
	}
	if c.Param("no_disturb") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群个人信息失败,请传入no_disturb!"))
	}
	no_disturb, err := strconv.ParseBool(c.Param("no_disturb"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群个人信息失败,no_disturb必须为布尔型"))
	}
	if c.Param("is_top") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群个人信息失败,请传入is_top!"))
	}
	is_top, err := strconv.ParseBool(c.Param("is_top"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群个人信息失败,is_top必须为布尔型"))
	}
	b, err := l.SetChartGroupPersonalInfo(user_id, group_id, strings.Trim(c.Param("user_name"), " "), no_disturb, is_top)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "设置聊天群个人信息失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "设置聊天群个人信息成功"))
})

var SetChartGroupManager = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群管理员失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群管理员失败,user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群管理员失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群管理员失败,group_id必须为整型"))
	}
	var user_list []int64
	if c.Param("user_list") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "设置聊天群管理员失败,请传入user_list!"))
	}
	user_list, err = u.StringToIntArray(c.Param("user_list"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "设置聊天群管理员失败,user_list必须为整型数组"))
	}
	b, err := l.SetChartGroupManager(user_id, group_id, user_list)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "设置聊天群管理员失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "设置聊天群管理员成功"))
})

var RequestJoinGroup = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "申请加入聊天群失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "申请加入聊天群失败,user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "申请加入聊天群失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "申请加入聊天群失败,group_id必须为整型"))
	}
	b, err := l.RequestJoinGroup(user_id, group_id, strings.Trim(c.Param("memo"), " "))
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "申请加入聊天群失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "申请加入聊天群成功"))
})

var AgreeJoinGroup = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "操作用户入群申请失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "操作用户入群申请失败,user_id必须为整型"))
	}
	if c.Param("request_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "操作用户入群申请失败,请传入request_id!"))
	}
	request_id, err := strconv.ParseInt(c.Param("request_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "操作用户入群申请失败,request_id必须为整型"))
	}
	if c.Param("action") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "操作用户入群申请失败,请传入action!"))
	}
	action, err := strconv.Atoi(c.Param("action"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "操作用户入群申请失败,action必须为整型"))
	}
	b, err := l.AgreeJoinGroup(user_id, request_id, action)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "操作用户入群申请失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "操作用户入群申请成功"))
})

var GetRequestJoinGroupList = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "申请加入聊天群失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "申请加入聊天群失败,user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "申请加入聊天群失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "申请加入聊天群失败,group_id必须为整型"))
	}
	data, err := l.GetRequestJoinGroupList(user_id, group_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetGroupQrcode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "生成聊天群二维码失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "生成聊天群二维码失败,user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "生成聊天群二维码失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "生成聊天群二维码失败,group_id必须为整型"))
	}
	data, err := l.GetGroupQrcode(user_id, group_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))

})

var ChartGroupQrcode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("share_user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "扫一扫群二维码获取群信息失败,请传入share_user_id!"))
	}
	share_user_id, err := strconv.ParseInt(c.Param("share_user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,share_user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,group_id必须为整型"))
	}
	if c.Param("group_type") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,请传入group_type!"))
	}
	group_type, err := strconv.ParseInt(c.Param("group_type"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,group_type必须为整型"))
	}
	if c.Param("join_sign") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,请传入join_sign!"))
	}
	join_sign, err := strconv.ParseBool(c.Param("join_sign"))
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,join_sign必须为布尔型"))
	}
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "扫一扫群二维码获取群信息失败,请传入share_user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,user_id必须为整型"))
	}
	data, err := l.ChartGroupQrcode(user_id, share_user_id, group_id, group_type, join_sign)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var GetChartGroupMsg = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取群离线消息失败,请传入user_id!"))
		}
		user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取群离线消息失败,user_id必须为整型"))
		}
		if c.Param("group_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取群离线消息失败,请传入group_id!"))
		}
		group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取群离线消息失败,group_id必须为整型"))
		}
		var last_time string
		if c.Param("last_time") == "" {
			last_time = time.Now().Format("2006-01-02 15:04:05")
		} else {
			last_time = c.Param("last_time")
		}

		var page_size int
		if c.Param("page_size") != "" {
			page_size, _ = strconv.Atoi(c.Param("page_size"))
		}
		if page_size <= 0 {
			page_size = 20
		}
		data, err := l.GetChartGroupMsg(user_id, group_id, last_time, page_size)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"获取群离线消息(get_chart_group_msg)",
	"服务器端返回的群离线消息",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(6), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	faygo.ParamInfo{
		Name:     "group_id",
		In:       "query",
		Required: true,
		Model:    int(106), // API文档中显示的参数默认值
		Desc:     "群编号",
	},
	faygo.ParamInfo{
		Name:     "last_time",
		In:       "query",
		Required: false,
		Model:    string("2017-09-01 00:00:00"), // API文档中显示的参数默认值
		Desc:     "上次最后时间",
	},
	faygo.ParamInfo{
		Name:     "page_size",
		In:       "query",
		Required: false,
		Model:    int(20), // API文档中显示的参数默认值
		Desc:     "消息条目数量",
	},
)

var SendGroupMsg = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" && c.Param("user_id") != "system" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发送群消息失败,请传入user_id!"))
	}
	if c.Param("group_id") == "" && c.Param("group_id") != "system" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发送群消息失败,请传入group_id!"))
	}
	var group_id, sender_id int64
	var err error
	sender_id, err = strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发送群消息失败,user_id必须为整型"))
	}
	group_id, err = strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发送群消息失败,group_id必须为整型"))
	}
	content := strings.Trim(c.Param("content"), " ")
	if len(content) == 0 {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "发送群消息失败,不能发送内容为空的消息"))
	}
	var send_type int
	if c.Param("send_type") == "" {
		send_type = 1
	} else {
		send_type, err = strconv.Atoi(c.Param("send_type"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, "401", "发送失败,传入send_type参数必须为整型"))
		}

	}
	msg_id, send_time, successed := l.SendGroupChartMsg(sender_id, group_id, send_type, content)
	if successed {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","msg_id":%d,"send_time":"%s"}`, 0, "success", msg_id, send_time.Format("2006-01-02 15:04:05")))
	} else {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "发送群消息失败"))
	}
})

var RemoveChartGroup = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "移除群成员失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "移除群成员失败,user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "移除群成员失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "移除群成员失败,group_id必须为整型"))
	}
	if c.Param("object_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "移除群成员失败,请传入object_id!"))
	}
	object_id, err := strconv.ParseInt(c.Param("object_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "移除群成员失败,object_id必须为整型"))
	}
	if user_id == object_id {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "移除群成员失败,不能移除自己"))
	}
	b, err := l.RemoveChartGroup(user_id, group_id, object_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	if !b {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "移除群成员失败"))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "移除群成员成功"))
})

var GetChartGroupMember = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取群成员列表失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取群成员列表失败,user_id必须为整型"))
	}
	if c.Param("group_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取群成员列表失败,请传入group_id!"))
	}
	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取群成员列表失败,group_id必须为整型"))
	}
	data, err := l.GetChartGroupMember(user_id, group_id)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ScanQrcode = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("scene") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫码失败,请传入scene!"))
	}
	scene := strings.ToLower(strings.Trim(c.Param("scene"), " "))
	var bind_type int
	var err error
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	switch scene {
	case "scan_build_qrcode":
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "扫一扫二维码获取信息失败,请传入user_id!"))
		}
		user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,user_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,请传入project_id!"))
		}
		project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,project_id必须为整型"))
		}
		data, err := l.ScanBuildQrcode(user_id, project_id, bind_type, c.Param("phone"), strings.Trim(c.Param("scene"), " "))
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if len(data) > 0 {
			return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, data))
		} else {
			return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, "功能尚未开放，请等待下一次更新"))
		}
	case "chart_group_qrcode":
		if c.Param("share_user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "扫一扫群二维码获取群信息失败,请传入share_user_id!"))
		}
		share_user_id, err := strconv.ParseInt(c.Param("share_user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,share_user_id必须为整型"))
		}
		if c.Param("group_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,请传入group_id!"))
		}
		group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,group_id必须为整型"))
		}
		if c.Param("group_type") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,请传入group_type!"))
		}
		group_type, err := strconv.ParseInt(c.Param("group_type"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,group_type必须为整型"))
		}
		if c.Param("join_sign") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,请传入join_sign!"))
		}
		join_sign, err := strconv.ParseBool(c.Param("join_sign"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,join_sign必须为布尔型"))
		}
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "扫一扫群二维码获取群信息失败,请传入user_id!"))
		}
		user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫群二维码获取群信息失败,user_id必须为整型"))
		}
		data, err := l.ChartGroupQrcode(user_id, share_user_id, group_id, group_type, join_sign)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	case "qrcode_login":
		var app_name string
		if c.Param("app_name") == "" {
			app_name = "com.dingfang.dffangkai"
		} else {
			app_name = strings.Trim(c.Param("app_name"), " ")
		}
		var project_id, company_id int64
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 401, "扫码登录失败,请传入user_id!"))
		}
		user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 401, "扫码登录失败,user_id必须为整型"))
		}
		switch app_name {
		case "com.dingfang.dffangkai":
			if c.Param("company_id") == "" {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 401, "扫码登录失败,请传入company_id!"))
			}
			company_id, err = strconv.ParseInt(c.Param("company_id"), 10, 64)
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 401, "扫码登录失败,company_id必须为整型"))
			}
		case "com.dingfang.dfjianzhuang":
			if c.Param("project_id") == "" {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 401, "扫码登录失败,请传入project_id!"))
			}
			project_id, err = strconv.ParseInt(c.Param("project_id"), 10, 64)
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 401, "扫码登录失败,project_id必须为整型"))
			}

		case "com.dingfang.dffuwushang":
			if c.Param("company_id") == "" {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 401, "扫码登录失败,请传入company_id!"))
			}
			company_id, err = strconv.ParseInt(c.Param("company_id"), 10, 64)
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 401, "扫码登录失败,company_id必须为整型"))
			}
		}
		isok, err := logic.QrcodeLogin(user_id, project_id, company_id, bind_type, strings.Trim(c.Param("uid"), " "), c.RealIP(), app_name)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 500, err))
		}
		if !isok {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"qrcode_login"}}`, 1000, "抱歉啊，系统出问题了，请联系管理员"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":true,"func":"qrcode_login"}}`, 0, "success"))
	case "invitation_join_project":
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "扫一扫二维码获取信息失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,user_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,请传入project_id!"))
		}
		projectId, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,project_id必须为整型"))
		}
		if c.Param("inviter_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "扫一扫二维码获取信息失败,请传入user_id!"))
		}
		inviterId, err := strconv.ParseInt(c.Param("inviter_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,inviter_id必须为整型"))
		}
		var bindType int
		if c.Param("bind_type") != "" {
			bindType, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "扫一扫二维码获取信息失败,bind_type必须为整型!"))
			}
		} else {
			bindType = 1
		}
		data, err := l.GetInvitationJoinProject(userId, inviterId, projectId, bindType)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}
	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":{"login_status":false,"func":"unkown"}}`, 0, "功能尚未开放，请等待下一次更新"))
})

var BindDevelopProject = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "绑定房企项目失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定房企项目失败,user_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "绑定房企项目失败,请传入project_id!"))
		}
		projectId, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定房企项目失败,project_id必须为整型"))
		}
		if c.Param("develop_project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "绑定房企项目失败,请传入develop_project_id!"))
		}
		developProjectId, err := strconv.ParseInt(c.Param("develop_project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定房企项目失败,develop_project_id必须为整型"))
		}
		var isAccept bool
		if strings.Trim(c.Param("is_accept"), " ") != "" {
			isAccept, err = strconv.ParseBool(c.Param("is_accept"))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定房企项目失败,is_accept必须为布尔型"))
			}
		} else {
			isAccept = true
		}
		var bindType int
		if c.Param("bind_type") != "" {
			bindType, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定房企项目失败,bind_type必须为整型!"))
			}
		} else {
			bindType = 1
		}
		b, err := l.BindDevelopProject(userId, projectId, developProjectId, isAccept, bindType)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "绑定房企项目失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "绑定房企项目成功"))
	}),
	"bind_develop_project",
	"绑定房企项目",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	faygo.ParamInfo{
		Name:     "project_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目编号",
	},
	faygo.ParamInfo{
		Name:     "develop_project_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "房企项目编号",
	},
	faygo.ParamInfo{
		Name:     "is_accept",
		In:       "query",
		Required: false,
		Model:    bool(true), // API文档中显示的参数默认值
		Desc:     "是否同意绑定房开企业",
	},
	faygo.ParamInfo{
		Name:     "bind_type",
		In:       "query",
		Required: false,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目类型， 1：建筑类;2：装修类",
	},
)

var GetBindDevelopProjectInfo = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取申请绑定项目请求的数据失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取申请绑定项目请求的数据失败,user_id必须为整型"))
		}
		if c.Param("develop_project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取申请绑定项目请求的数据失败,请传入develop_project_id!"))
		}
		developProjectId, err := strconv.ParseInt(c.Param("develop_project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取申请绑定项目请求的数据失败,develop_project_id必须为整型"))
		}
		data, err := l.GetBindDevelopProjectInfo(userId, developProjectId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_bind_develop_project_info",
	"获取申请绑定项目请求的数据",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	faygo.ParamInfo{
		Name:     "develop_project_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "房企项目编号",
	},
)

var GetBindingProject = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取可进行绑定的项目列表失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取可进行绑定的项目列表失败,user_id必须为整型"))
		}
		data, err := l.GetBindingProject(userId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_binding_project",
	"获取可进行绑定的项目列表",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
)

var GetMyShop = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取我的店铺的状态,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取我的店铺的状态,user_id必须为整型"))
		}
		data, err := l.GetMyShop(userId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_my_shop_status",
	"获取我的店铺的状态",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
)

var CreateMyShop = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的店铺失败,user_id必须为整型"))
		}
		if strings.Trim(c.Param("shop_name"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入shop_name!"))
		}
		shopName := strings.Trim(c.Param("shop_name"), " ")
		if strings.Trim(c.Param("shop_phone"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入shop_phone!"))
		}
		shopPhone := strings.Trim(c.Param("shop_phone"), " ")
		if c.Param("delivery") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入delivery!"))
		}
		delivery, err := strconv.ParseFloat(c.Param("delivery"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的店铺失败,delivery必须为浮点型"))
		}
		if c.Param("full_relief") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入full_relief!"))
		}
		fullRelief, err := strconv.ParseFloat(c.Param("full_relief"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的店铺失败,full_relief必须为浮点型"))
		}
		if c.Param("start_amount") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入start_amount!"))
		}
		startAmount, err := strconv.ParseFloat(c.Param("start_amount"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的店铺失败,start_amount必须为浮点型"))
		}
		if c.Param("lon") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入lon!"))
		}
		lon, err := strconv.ParseFloat(c.Param("lon"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的店铺失败,lon必须为浮点型"))
		}
		if c.Param("lat") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入lat!"))
		}
		lat, err := strconv.ParseFloat(c.Param("lat"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的店铺失败,lat必须为浮点型"))
		}
		b, shopId, err := l.CreateShop(userId, shopName, shopPhone, strings.Trim(c.Param("shop_pic"), " "), strings.Trim(c.Param("shop_addr"), " "),
			strings.Trim(c.Param("start_time"), " "), strings.Trim(c.Param("end_time"), " "), delivery, fullRelief, startAmount, lon, lat)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "创建我的店铺失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","shop_id":%d}`, code, "创建我的店铺成功", shopId))
	}),
	"create_shop",
	"创建我的店铺",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "店铺名称",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_pic",
		In:       "query",
		Required: false,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "店铺头像",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_phone",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "店铺联系电话",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "delivery",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "配送费用",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "full_relief",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "满额免配送费",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "start_amount",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "起送金额",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "start_time",
		In:       "query",
		Required: false,
		Model:    string("7:30"), // API文档中显示的参数默认值
		Desc:     "开始营业时间",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "end_time",
		In:       "query",
		Required: false,
		Model:    string("22:00"), // API文档中显示的参数默认值
		Desc:     "结束营业时间",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "lon",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "店铺经度",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "lat",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "店铺维度",
	},
)

var EditMyShop = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "编辑我的店铺失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑我的店铺失败,user_id必须为整型"))
		}
		if strings.Trim(c.Param("shop_name"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "编辑我的店铺失败,请传入shop_name!"))
		}
		shopName := strings.Trim(c.Param("shop_name"), " ")
		if strings.Trim(c.Param("shop_phone"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "编辑我的店铺失败,请传入shop_phone!"))
		}
		shopPhone := strings.Trim(c.Param("shop_phone"), " ")
		if c.Param("delivery") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "编辑我的店铺失败,请传入delivery!"))
		}
		delivery, err := strconv.ParseFloat(c.Param("delivery"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑我的店铺失败,delivery必须为浮点型"))
		}
		if c.Param("full_relief") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入full_relief!"))
		}
		fullRelief, err := strconv.ParseFloat(c.Param("full_relief"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的店铺失败,full_relief必须为浮点型"))
		}
		if c.Param("start_amount") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建我的店铺失败,请传入start_amount!"))
		}
		startAmount, err := strconv.ParseFloat(c.Param("start_amount"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的店铺失败,start_amount必须为浮点型"))
		}
		if c.Param("lon") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "编辑我的店铺失败,请传入lon!"))
		}
		lon, err := strconv.ParseFloat(c.Param("lon"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑我的店铺失败,lon必须为浮点型"))
		}
		if c.Param("lat") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "编辑我的店铺失败,请传入lat!"))
		}
		lat, err := strconv.ParseFloat(c.Param("lat"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "编辑我的店铺失败,lat必须为浮点型"))
		}
		b, err := l.EditMyShop(userId, shopName, shopPhone, strings.Trim(c.Param("shop_pic"), " "), strings.Trim(c.Param("shop_addr"), " "),
			strings.Trim(c.Param("start_time"), " "), strings.Trim(c.Param("end_time"), " "), delivery, fullRelief, startAmount, lon, lat)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "编辑我的店铺失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "编辑我的店铺成功"))
	}),
	"edit_shop",
	"编辑我的店铺",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "店铺名称",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_pic",
		In:       "query",
		Required: false,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "店铺头像",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_phone",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "店铺联系电话",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "delivery",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "配送费用",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "full_relief",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "满额免配送费",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "start_amount",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "起送金额",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "start_time",
		In:       "query",
		Required: false,
		Model:    string("7:30"), // API文档中显示的参数默认值
		Desc:     "开始营业时间",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "end_time",
		In:       "query",
		Required: false,
		Model:    string("22:00"), // API文档中显示的参数默认值
		Desc:     "结束营业时间",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "lon",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "店铺经度",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "lat",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "店铺维度",
	},
)

var AddShopGoodsType = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加我的店铺中商品分类失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加我的店铺中商品分类失败,user_id必须为整型"))
		}
		if strings.Trim(c.Param("type_name"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加我的店铺中商品分类失败,请传入type_name!"))
		}
		b, typrId, err := l.AddShopGoodsType(userId, strings.Trim(c.Param("type_name"), " "))
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加我的店铺中商品分类失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","type_id":%d}`, code, "增加我的店铺中商品分类成功", typrId))
	}),
	"add_shop_goods_type",
	"增加我的店铺中商品分类",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "type_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品类型名称",
	},
)

var EditShopGoodsType = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的店铺中商品分类失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的店铺中商品分类失败,user_id必须为整型"))
		}
		if c.Param("type_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的店铺中商品分类失败,请传入type_id!"))
		}
		typeId, err := strconv.ParseInt(c.Param("type_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的店铺中商品分类失败,type_id必须为整型"))
		}
		if strings.Trim(c.Param("type_name"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的店铺中商品分类失败,请传入type_name!"))
		}
		b, err := l.EditShopGoodsType(userId, typeId, strings.Trim(c.Param("type_name"), " "))
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改我的店铺中商品分类失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改我的店铺中商品分类成功"))
	}),
	"edit_shop_goods_type",
	"修改我的店铺中商品分类",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "type_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "商品类型编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "type_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品类型名称",
	},
)

var DeleteShopGoodsType = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "删除我的店铺中商品分类失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除我的店铺中商品分类失败,user_id必须为整型"))
		}
		if c.Param("type_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "删除我的店铺中商品分类失败,请传入type_id!"))
		}
		typeId, err := strconv.ParseInt(c.Param("type_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除我的店铺中商品分类失败,type_id必须为整型"))
		}
		b, err := l.DeleteShopGoodsType(userId, typeId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除我的店铺中商品分类失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除我的店铺中商品分类成功"))
	}),
	"delete_shop_goods_type",
	"删除我的店铺中商品分类",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "type_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "商品类型编号",
	},
)

var GetGoodsTypeList = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取我的店铺中商品分类失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取我的店铺中商品分类失败,user_id必须为整型"))
		}
		data, err := l.GetGoodsTypeList(userId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_my_shop_goods_type_list",
	"获取我的店铺中商品分类",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
)

var AddShopGoods = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "新增我的店铺中商品失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增我的店铺中商品失败,user_id必须为整型"))
		}
		if c.Param("type_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "新增我的店铺中商品失败,请传入type_id!"))
		}
		typeId, err := strconv.ParseInt(c.Param("type_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增我的店铺中商品失败,type_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "新增我的店铺中商品失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增我的店铺中商品失败,shop_id必须为整型"))
		}
		if strings.Trim(c.Param("goods_name"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增我的店铺中商品失败,goods_name必须为必填字段"))
		}
		if c.Param("price") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "新增我的店铺中商品失败,请传入price!"))
		}
		price, err := strconv.ParseFloat(c.Param("price"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "新增我的店铺中商品失败,price必须为浮点型"))
		}
		b, goodsId, err := l.AddShopGoods(userId, shopId, typeId, strings.Trim(c.Param("goods_name"), " "),
			strings.Trim(c.Param("goods_pic"), " "), price)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "新增我的店铺中商品失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","goods_id":%d}`, code, "新增我的店铺中商品成功", goodsId))
	}),
	"add_my_shop_goods",
	"增加店铺商品",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "type_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "商品类型编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品名称",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_pic",
		In:       "query",
		Required: false,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品图片",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "price",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "商品价格",
	},
)

var EditShopGoods = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的店铺中商品失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的店铺中商品失败,user_id必须为整型"))
		}
		if c.Param("goods_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的店铺中商品失败,请传入goods_id!"))
		}
		goodsId, err := strconv.ParseInt(c.Param("goods_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的店铺中商品失败,goods_id必须为整型"))
		}
		if c.Param("type_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的店铺中商品失败,请传入type_id!"))
		}
		typeId, err := strconv.ParseInt(c.Param("type_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的店铺中商品失败,type_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的店铺中商品失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的店铺中商品失败,shop_id必须为整型"))
		}
		if strings.Trim(c.Param("goods_name"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的店铺中商品失败,goods_name必须为必填字段"))
		}
		if c.Param("price") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的店铺中商品失败,请传入price!"))
		}
		price, err := strconv.ParseFloat(c.Param("price"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的店铺中商品失败,price必须为浮点型"))
		}
		b, err := l.EditShopGoods(userId, shopId, typeId, goodsId, strings.Trim(c.Param("goods_name"), " "),
			strings.Trim(c.Param("goods_pic"), " "), price)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改我的店铺中商品失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改我的店铺中商品成功"))
	}),
	"edit_my_shop_goods",
	"修改我的店铺中商品",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "商品编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "type_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "商品类型编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品名称",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_pic",
		In:       "query",
		Required: false,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品图片",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "price",
		In:       "query",
		Required: true,
		Model:    float64(0), // API文档中显示的参数默认值
		Desc:     "商品价格",
	},
)

var DeleteShopGoods = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "删除我的店铺中的商品失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除我的店铺中的商品失败,user_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "删除我的店铺中的商品失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除我的店铺中的商品失败,shop_id必须为整型"))
		}
		goodsArr := strings.Split(c.Param("goods_id"), ",")
		if len(goodsArr) < 0 || goodsArr[0] == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除我的店铺中的商品失败,goods_id至少有一个以上的商品编号"))
		}
		b, err := l.DeleteShopGoods(userId, shopId, goodsArr)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除我的店铺中的商品失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除我的店铺中的商品成功"))
	}),
	"del_my_shop_goods",
	"删除我的店铺中的商品",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_id",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品编号，多个商品用英文逗号进行分割",
	},
)

var GetShopDetail = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取店铺详情失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取店铺详情失败,user_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取店铺详情失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取店铺详情失败,shop_id必须为整型"))
		}
		data, err := l.GetShopDetail(userId, shopId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_shop_detail",
	"获取店铺详情",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
)

var GetShopGoodsDetail = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取店铺中的商品详情失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取店铺中的商品详情失败,user_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取店铺中的商品详情失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取店铺中的商品详情失败,shop_id必须为整型"))
		}
		if c.Param("goods_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取店铺中的商品详情失败,请传入goods_id!"))
		}
		goodsId, err := strconv.ParseInt(c.Param("goods_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取店铺中的商品详情失败,goods_id必须为整型"))
		}
		data, err := l.GetShopGoodsDetail(userId, shopId, goodsId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_shop_goods_detail",
	"获取店铺中的商品详情",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "商品编号",
	},
)

var AddMyCart = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加到我的购物车失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加到我的购物车失败,user_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加到我的购物车失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加到我的购物车失败,shop_id必须为整型"))
		}
		if c.Param("goods_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加到我的购物车失败,请传入goods_id!"))
		}
		goodsId, err := strconv.ParseInt(c.Param("goods_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加到我的购物车失败,goods_id必须为整型"))
		}
		if c.Param("count") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "增加到我的购物车失败,请传入count!"))
		}
		count, err := strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加到我的购物车失败,count必须为整型"))
		}
		b, err := l.AddMyCart(userId, shopId, goodsId, count)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加到我的购物车失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加到我的购物车成功"))
	}),
	"add_my_cart",
	"增加到我的购物车",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "商品编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "count",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "商品数量",
	},
)

var DelMyCart = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "从我的购物车中移除失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "从我的购物车中移除失败,user_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "从我的购物车中移除失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "从我的购物车中移除失败,shop_id必须为整型"))
		}
		if c.Param("goods_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "从我的购物车中移除失败,请传入goods_id!"))
		}
		goodsId, err := strconv.ParseInt(c.Param("goods_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "从我的购物车中移除失败,goods_id必须为整型"))
		}
		b, err := l.DelMyCart(userId, shopId, goodsId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "从我的购物车中移除失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "从我的购物车中移除失败成功"))
	}),
	"del_my_cart",
	"从我的购物车中移除",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "商品编号",
	},
)

var ModifyCartGoodsCount = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的购物车中商品数量失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的购物车中商品数量失败,user_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的购物车中商品数量失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的购物车中商品数量失败,shop_id必须为整型"))
		}
		if c.Param("goods_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的购物车中商品数量失败,请传入goods_id!"))
		}
		goodsId, err := strconv.ParseInt(c.Param("goods_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的购物车中商品数量失败,goods_id必须为整型"))
		}
		if c.Param("count") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的购物车中商品数量失败,请传入count!"))
		}
		count, err := strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的购物车中商品数量失败,count必须为整型"))
		}
		b, err := l.ModifyCartGoodsCount(userId, shopId, goodsId, count)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改我的购物车中商品数量失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改我的购物车中商品数量成功"))
	}),
	"modify_cart_goods_count",
	"修改我的购物车中商品数量失败",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "goods_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "商品编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "count",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "商品数量",
	},
)

var GetMyCart = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的购物车中商品数量失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的购物车中商品数量失败,user_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改我的购物车中商品数量失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改我的购物车中商品数量失败,shop_id必须为整型"))
		}
		data, err := l.GetMyCart(userId, shopId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_my_cart",
	"获取我的购物车中商品",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
)

var CreateShopOrder = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建店铺订单失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建店铺订单失败,user_id必须为整型"))
		}
		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "创建店铺订单失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建店铺订单失败,shop_id必须为整型"))
		}
		detail := strings.Trim(c.Param("detail"), " ")
		if len(detail) < 1 {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建店铺订单失败,detail为必填字段"))
		}
		b, orderId, err := l.CreateShopOrder(userId, shopId, detail)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "创建店铺订单失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","order_id":%d}`, code, "创建店铺订单成功", orderId))
	}),
	"create_shop_order",
	"创建店铺订单",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "detail",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品明细",
	},
)

var GetShopOrderDetail = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取店铺订单详情失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取店铺订单详情失败,user_id必须为整型"))
		}
		if c.Param("order_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取店铺订单详情失败,请传入order_id!"))
		}
		orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取店铺订单详情失败,order_id必须为整型"))
		}
		data, err := l.GetShopOrderDetail(userId, orderId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_shop_order_detail",
	"获取店铺订单详情",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "order_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "订单编号",
	},
)

var EditShopOrderAddr = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改订单送货地址失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单送货地址失败,user_id必须为整型"))
		}
		if c.Param("order_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改订单送货地址失败,请传入order_id!"))
		}
		orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单送货地址失败,order_id必须为整型"))
		}
		if c.Param("addr_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改订单送货地址失败,请传入addr_id!"))
		}
		addrId, err := strconv.ParseInt(c.Param("addr_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单送货地址失败,addr_id必须为整型"))
		}
		b, err := l.EditShopOrderAddr(userId, orderId, addrId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改订单送货地址失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","order_id":%d}`, code, "修改订单送货地址成功", orderId))
	}),
	"edit_shop_order_addr",
	"修改订单送货地址",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "order_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "订单编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "addr_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户送货地址编号",
	},
)

var EditShopOrderMemo = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改订单备注信息失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单备注信息失败,user_id必须为整型"))
		}
		if c.Param("order_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改订单备注信息失败,请传入order_id!"))
		}
		orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单备注信息失败,order_id必须为整型"))
		}

		b, err := l.EditShopOrderMemo(userId, orderId, strings.Trim(c.Param("memo"), " "))
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改订单备注信息失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","order_id":%d}`, code, "修改订单备注信息成功", orderId))
	}),
	"edit_shop_order_addr",
	"修改订单备注信息",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "order_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "订单编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "memo",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "订单备注",
	},
)

var GetShopOrderList = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取我的店铺订单列表失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取我的店铺订单列表失败,user_id必须为整型"))
		}
		var pageNum, count int
		if c.Param("page_num") == "" {
			pageNum = 1
		} else {
			pageNum, err = strconv.Atoi(c.Param("page_num"))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，page_num必须为整型"))
			}
		}
		if c.Param("count") == "" {
			count = 20
		} else {
			count, err = strconv.Atoi(c.Param("count"))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，count必须为整型"))
			}
		}
		var status string
		if strings.Trim(c.Param("status"), " ") != "" {
			status = strings.Trim(c.Param("status"), " ")

		} else {
			status = "-1"
		}
		data, err := l.GetShopOrderList(userId, pageNum, count, status)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_shop_order_list",
	"获取我的店铺订单列表",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "page_num",
		In:       "query",
		Required: false,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "页码，默认1",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "count",
		In:       "query",
		Required: false,
		Model:    int(20), // API文档中显示的参数默认值
		Desc:     "返回的数量，默认20",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "status",
		In:       "query",
		Required: false,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "订单状态，默认-1获取全部状态，0:等待支付;1:支付完成，等待商家确认;2:商家已确认，等待确认收货;3:已确认收货，订单完成;如果要获取多个状态，用英文逗号分隔",
	},
)

var GetBuyerShopOrderList = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取我的店铺订单列表失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取我的店铺订单列表失败,user_id必须为整型"))
		}
		var page_num, count int
		if c.Param("page_num") == "" {
			page_num = 1
		} else {
			page_num, err = strconv.Atoi(c.Param("page_num"))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，page_num必须为整型"))
			}
		}
		if c.Param("count") == "" {
			count = 20
		} else {
			count, err = strconv.Atoi(c.Param("count"))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，count必须为整型"))
			}
		}
		var status string
		if strings.Trim(c.Param("status"), " ") != "" {
			status = strings.Trim(c.Param("status"), " ")

		} else {
			status = "-1"
		}
		data, err := l.GetBuyerShopOrderList(userId, page_num, count, status)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_buyer_shop_order_list",
	"获取我的店铺购买订单列表",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "page_num",
		In:       "query",
		Required: false,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "页码，默认1",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "count",
		In:       "query",
		Required: false,
		Model:    int(20), // API文档中显示的参数默认值
		Desc:     "返回的数量，默认20",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "status",
		In:       "query",
		Required: false,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "订单状态，默认-1获取全部状态，0:等待支付;1:支付完成，等待商家确认;2:商家已确认，等待确认收货;3:已确认收货，订单完成;如果要获取多个状态，用英文逗号分隔",
	},
)

var PayShopOrde = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "结算我的店铺购买订单失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "结算我的店铺购买订单失败,user_id必须为整型"))
		}
		if c.Param("order_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "结算我的店铺购买订单失败,请传入order_id!"))
		}
		orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "结算我的店铺购买订单失败,order_id必须为整型"))
		}
		if c.Param("pay_amount") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "结算我的店铺购买订单失败,请传入order_id!"))
		}
		payAmount, err := strconv.ParseFloat(c.Param("pay_amount"), 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "结算我的店铺购买订单失败,pay_amount必须为数字型"))
		}
		paySignPass := strings.Trim(c.Param("pay_sign_pass"), " ")
		if len(paySignPass) < 1 {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "结算我的店铺购买订单失败,请传入pay_sign_pass!"))
		}
		b, err := l.PayShopOrde(userId, orderId, payAmount, paySignPass)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "结算我的店铺购买订单失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","order_id":%d}`, code, "结算我的店铺购买订单成功", orderId))
	}),
	"pay_shop_order",
	"结算我的店铺购买订单",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "order_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "订单编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "pay_amount",
		In:       "query",
		Required: true,
		Model:    float64(5), // API文档中显示的参数默认值
		Desc:     "支付金额",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "pay_sign_pass",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "验签后的支付密码",
	},
)

var ConfirmShopOrder = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "卖家我的店铺确认订单失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "卖家我的店铺确认订单失败,user_id必须为整型"))
		}
		if c.Param("order_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "卖家我的店铺确认订单失败,请传入order_id!"))
		}
		orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "卖家我的店铺确认订单失败,order_id必须为整型"))
		}
		b, err := l.ConfirmShopOrde(userId, orderId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "卖家我的店铺确认订单失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","order_id":%d}`, code, "卖家我的店铺确认订单成功", orderId))
	}),
	"confirm_shop_order",
	"卖家我的店铺确认订单",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "order_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "订单编号",
	},
)

var ShopOrderConfirmReceipt = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "买家确认订单收货失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "买家确认订单收货失败,user_id必须为整型"))
		}
		if c.Param("order_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "买家确认订单收货失败,请传入order_id!"))
		}
		orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "买家确认订单收货失败,order_id必须为整型"))
		}
		b, err := l.ShopOrderConfirmReceipt(userId, orderId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "买家确认订单收货失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","order_id":%d}`, code, "买家确认订单收货成功", orderId))
	}),
	"shop_order_confirm_receipt",
	"买家确认订单收货",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "order_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "订单编号",
	},
)

var EditShopOrderDetail = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改订单中商品的数量失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单中商品的数量失败,user_id必须为整型"))
		}
		if c.Param("order_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "修改订单中商品的数量失败,请传入order_id!"))
		}
		orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单中商品的数量失败,order_id必须为整型"))
		}
		detail := strings.Trim(c.Param("detail"), " ")
		if len(detail) < 1 {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "修改订单中商品的数量失败,detail为必填字段"))
		}
		b, err := l.EditShopOrderDetail(userId, orderId, detail)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "修改订单中商品的数量失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","order_id":%d}`, code, "修改订单中商品的数量成功", orderId))
	}),
	"edit_shop_order_detail_count",
	"修改订单中商品的数量",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "order_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "订单编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "detail",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "商品明细",
	},
)

var EmptyingCart = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "清空购物车失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "清空购物车失败,user_id必须为整型"))
		}

		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "清空购物车失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "清空购物车失败,shop_id必须为整型"))
		}
		b, err := l.EmptyingCart(userId, shopId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "清空购物车失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "清空购物车成功"))
	}),
	"emptying_cart",
	"清空购物车",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
)

var OpenShopRequest = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "开店申请失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "开店申请失败,user_id必须为整型"))
		}

		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "开店申请失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "开店申请失败,shop_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "开店申请失败,请传入project_id!"))
		}
		projectId, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "开店申请失败,project_id必须为整型"))
		}
		if c.Param("bind_type") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "开店申请失败,请传入bind_type!"))
		}
		bindType, err := strconv.Atoi(c.Param("bind_type"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "开店申请失败,bind_type必须为整型"))
		}
		b, err := l.OpenShopRequest(userId, shopId, projectId, bindType)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "开店申请失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "开店申请成功"))
	}),
	"open_shop_request",
	"开店申请",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "project_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "项目编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "bind_type",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目类型，1:建筑类;2:装修类",
	},
)

var AuditOpenShopRequest = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "是否同意在项目中的开店申请操作失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "是否同意在项目中的开店申请操作失败,user_id必须为整型"))
		}

		if c.Param("shop_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "是否同意在项目中的开店申请操作失败,请传入shop_id!"))
		}
		shopId, err := strconv.ParseInt(c.Param("shop_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "是否同意在项目中的开店申请操作失败,shop_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "是否同意在项目中的开店申请操作失败,请传入project_id!"))
		}
		projectId, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "是否同意在项目中的开店申请操作失败,project_id必须为整型"))
		}
		if c.Param("bind_type") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "是否同意在项目中的开店申请操作失败,请传入bind_type!"))
		}
		bindType, err := strconv.Atoi(c.Param("bind_type"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "是否同意在项目中的开店申请操作失败,bind_type必须为整型"))
		}
		if c.Param("is_agree") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "是否同意在项目中的开店申请操作失败,请传入is_agree!"))
		}
		isAgree, err := strconv.ParseBool(c.Param("is_agree"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "是否同意在项目中的开店申请操作失败,is_agree必须为布尔型"))
		}
		b, err := l.AuditOpenShopRequest(userId, shopId, projectId, bindType, isAgree)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "是否同意在项目中的开店申请操作失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "是否同意在项目中的开店申请操作成功"))
	}),
	"audit_open_shop_request",
	"是否同意在项目中的开店申请",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "shop_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "店铺编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "project_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "项目编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "bind_type",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目类型，1:建筑类;2:装修类",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "is_agree",
		In:       "query",
		Required: true,
		Model:    bool(true), // API文档中显示的参数默认值
		Desc:     "是否同意店铺开通",
	},
)

var GetRequestProjectByShop = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取已申请的项目失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取已申请的项目失败,user_id必须为整型"))
		}
		var key = ""
		if strings.Trim(c.Param("keyword"), " ") != "" {
			key = strings.Trim(c.Param("keyword"), " ")
		}
		data, err := l.GetRequestProjectByShop(userId, key)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_request_project_by_shop",
	"获取已申请的项目以及周边1公里的项目",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "keyword",
		In:       "query",
		Required: false,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "查询关键字",
	},
)

var GetOpenShop = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取已开通的项目失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取已开通的项目失败,user_id必须为整型"))
		}
		data, err := l.GetOpenShop(userId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_open_shop",
	"获取已开通的项目",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
)

var GetOpenShopByProject = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取已开通的店铺失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取已开通的店铺失败,user_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取已开通的店铺失败,请传入project_id!"))
		}
		projectId, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取已开通的店铺失败,project_id必须为整型"))
		}
		data, err := l.GetOpenShopByProject(userId, projectId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_open_shop_by_project",
	"获取已开通的店铺",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "project_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目编号",
	},
)

var GetDocumentTypeList = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档类型失败,请传入user_id!"))
		}
		user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档类型失败,user_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档类型失败,请传入project_id!"))
		}
		project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取项目文档类型失败,project_id必须为整型"))
		}

		var bind_type int
		if c.Param("bind_type") != "" {
			bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
			}
		} else {
			bind_type = 1
		}
		data, err := l.GetDocumentTypeList(user_id, project_id, bind_type)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_doc_type_list",
	"获取项目文档类型",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "project_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "项目编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "bind_type",
		In:       "query",
		Required: false,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目绑定类型，1：建筑类;2:装修类",
	},
)

var GetMyDocDirList = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取我的私有项目文档目录失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取我的私有项目文档目录失败,user_id必须为整型"))
		}
		var parent = 0
		if strings.Trim(c.Param("parent"), " ") != "" {
			parent, err = strconv.Atoi(c.Param("parent"))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取我的私有项目文档目录失败,parent必须为整型!"))
			}
		}
		data, err := l.GetMyDocDirList(userId, parent)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_my_doc_dir_list",
	"获取我的私有项目文档目录",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
)

var GetDocumentList = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,user_id必须为整型"))
		}
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,请传入project_id!"))
		}
		projectId, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,project_id必须为整型"))
		}

		var bindType int
		if c.Param("bind_type") != "" {
			bindType, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,bind_type必须为整型!"))
			}
		} else {
			bindType = 1
		}
		if c.Param("type_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,请传入type_id!"))
		}
		typeId, err := strconv.ParseInt(c.Param("type_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,type_id必须为整型"))
		}
		//if c.Param("dir_id") == "" {
		//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,请传入dir_id!"))
		//}
		//dirId, err := strconv.ParseInt(c.Param("dir_id"), 10, 64)
		//if err != nil {
		//	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,dir_id必须为整型"))
		//}
		if c.Param("doc_type") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,请传入doc_type!"))
		}
		docType, err := strconv.Atoi(c.Param("doc_type"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取分类下的项目文档失败,doc_type必须为整型"))
		}
		data, err := l.GetProjectDocList(userId, projectId, typeId, bindType, docType, strings.Trim(c.Param("keyword"), " "))
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"get_project_doc_list",
	"获取分类下的项目文档",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "project_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "bind_type",
		In:       "query",
		Required: false,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目绑定类型，1:建筑类;2:装修类",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "type_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目文档类型编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "keyword",
		In:       "query",
		Required: false,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "搜索关键字",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "doc_type",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "项目文档目录类型，1:项目文档类型;2:共享文档类型;3:个人目录类型",
	},
)

var ShareDocument = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "分享文档失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "分享文档失败,user_id必须为整型"))
		}
		if c.Param("doc_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "分享文档失败,请传入doc_id!"))
		}
		docId, err := strconv.ParseInt(c.Param("doc_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "分享文档失败,doc_id必须为整型"))
		}
		if c.Param("target") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "分享文档失败,请传入target!"))
		}
		target, err := strconv.Atoi(c.Param("target"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "分享文档失败,target必须为整型"))
		}
		targetIdArr, err := u.StringToIntArray(strings.Trim(c.Param("target_id"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "分享文档失败,target_id必须为整型，多个用英文逗号进行分割"))
		}
		b, err := l.ShareDocument(userId, docId, target, targetIdArr)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "分享文档失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "分享文档成功"))
	}),
	"share_document",
	"分享文档",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "doc_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "文档编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "target",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "分享对象类型，1：用户;2:班组;3:项目",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "target_id",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "分享对象编号，多个对象用英文逗号进行分割",
	},
)

var RenameDocument = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重命名文档失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重命名文档失败,user_id必须为整型"))
		}
		if c.Param("doc_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重命名文档失败,请传入doc_id!"))
		}
		docId, err := strconv.ParseInt(c.Param("doc_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重命名文档失败,doc_id必须为整型"))
		}
		docName := strings.Trim(c.Param("doc_name"), " ")
		if len(docName) < 1 {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "重命名文档失败,请传入doc_name"))
		}
		b, err := l.RenameDocument(userId, docId, docName)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "重命名文档失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "重命名文档成功"))
	}),
	"rename_document",
	"重命名文档",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "doc_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "文档编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "doc_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "新的文档名称",
	},
)

var CancelShareDocument = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "取消分享文档失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "取消分享文档失败,user_id必须为整型"))
		}
		if c.Param("doc_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "取消分享文档失败,请传入doc_id!"))
		}
		docId, err := strconv.ParseInt(c.Param("doc_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "取消分享文档失败,doc_id必须为整型"))
		}
		if c.Param("target") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "取消分享文档失败,请传入target!"))
		}
		target, err := strconv.Atoi(c.Param("target"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "取消分享文档失败,target必须为整型"))
		}
		targetIdArr, err := u.StringToIntArray(strings.Trim(c.Param("target_id"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "取消分享文档失败,target_id必须为整型，多个用英文逗号进行分割"))
		}
		b, err := l.CancelShareDocument(userId, docId, target, targetIdArr)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "取消分享文档失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "取消分享文档成功"))
	}),
	"cancel_share_document",
	"取消分享文档",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "doc_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "文档编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "target",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "分享对象类型，1：用户;2:班组;3:项目",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "target_id",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "分享对象编号，多个对象用英文逗号进行分割",
	},
)

var DeleteDocument = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档失败,user_id必须为整型"))
		}
		if c.Param("doc_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档失败,请传入doc_id!"))
		}
		docId, err := strconv.ParseInt(c.Param("doc_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "删除文档失败,doc_id必须为整型"))
		}
		b, err := l.DeleteDocument(userId, docId)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除文档失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "删除文档成功"))
	}),
	"delete_document",
	"删除文档",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "doc_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "文档编号",
	},
)

var NewUploadDocument = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传文档失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传文档失败,user_id必须为整型"))
		}
		if c.Param("dir_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传文档失败,请传入dir_id!"))
		}
		dirId, err := strconv.ParseInt(c.Param("dir_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传文档失败,dir_id必须为整型"))
		}
		docName := strings.Trim(c.Param("doc_name"), " ")
		if len(docName) < 1 {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "上传文档失败,请传入doc_name"))
		}
		var imgList []string
		if c.Param("img_list") != "" {
			imgList = strings.Split(c.Param("img_list"), ",")
		}
		b, docId, err := l.NewUploadDocument(userId, dirId, docName, imgList)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "上传文档失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","doc_id":%d}`, code, "上传文档成功", docId))
	}),
	"new_upload_document",
	"上传文档",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "doc_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "文档名称",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "dir_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "文档存放目录编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "img_list",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "图片地址，多张图片用英文逗号进行分割",
	},
)

var NewDocDir = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的文档目录失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的文档目录失败,user_id必须为整型"))
		}

		dirName := strings.Trim(c.Param("dir_name"), " ")
		if len(dirName) < 1 {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "创建我的文档目录失败,请传入dir_name"))
		}
		b, dirId, err := l.NewDocDir(userId, dirName)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "创建我的文档目录失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","dir_id":%d}`, code, "创建我的文档目录成功", dirId))
	}),
	"new_doc_dir",
	"创建我的文档目录",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "dir_name",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "文档目录名称",
	},
)

var AddOrderExtendInfo = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("user_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加订单状态的扩展信息失败,请传入user_id!"))
		}
		userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加订单状态的扩展信息失败,user_id必须为整型"))
		}
		if c.Param("order_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加订单状态的扩展信息失败,请传入order_id!"))
		}
		orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加订单状态的扩展信息失败,order_id必须为整型"))
		}
		if strings.Trim(c.Param("param"), " ") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加订单状态的扩展信息失败,请传入param!"))
		}
		param := strings.Trim(c.Param("param"), " ")
		if c.Param("extend_type") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加订单状态的扩展信息失败,请传入extend_type!"))
		}
		extendType, err := strconv.ParseInt(c.Param("extend_type"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "增加订单状态的扩展信息失败,extend_type必须为整型"))
		}
		b, err := l.AddOrderExtendInfo(userId, orderId, extendType, param)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		if !b {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加订单状态的扩展信息失败"))
		}
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, "增加订单状态的扩展信息成功"))
	}),
	"add_order_extend_info",
	"增加订单状态的扩展信息",
	// 定义参数
	faygo.ParamInfo{
		Name:     "user_id",
		In:       "query",
		Required: true,
		Model:    int(5), // API文档中显示的参数默认值
		Desc:     "用户编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "order_id",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "订单编号",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "param",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "订单状态扩展信息",
	},
	// 定义参数
	faygo.ParamInfo{
		Name:     "extend_type",
		In:       "query",
		Required: true,
		Model:    int(1), // API文档中显示的参数默认值
		Desc:     "订单状态扩展信息类型， 1：发货信息；2：制作进度",
	},
)

var GetSecurityCheckList2 = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取安全检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取安全检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败,project_id必须为整型"))
	}
	var page_num, count int64
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，count必须为整型"))
		}
	}
	var keyword string
	if c.Param("keyword") != "" {
		keyword = strings.Trim(c.Param("keyword"), " ")
	}
	var status int
	if c.Param("status") != "" {
		status, err = strconv.Atoi(c.Param("status"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，status必须为整型"))
		}
	} else {
		status = -1
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}

	sc, count, err := l.GetSecurityCheck2(user_id, project_id, int((page_num-1)*count), int(count), status, bind_type, keyword)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(sc)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s,"total_count":%d}`, code, msg, string(b), count))
})

var GetQualityTestingList2 = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取质量检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取质量检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取质量检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取质量检查失败,project_id必须为整型"))
	}
	var page_num, count int64
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取质量检查失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取质量检查失败，count必须为整型"))
		}
	}
	var keyword string
	if c.Param("keyword") != "" {
		keyword = strings.Trim(c.Param("keyword"), " ")
	}
	var status int
	if c.Param("status") != "" {
		status, err = strconv.Atoi(c.Param("status"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，status必须为整型"))
		}
	} else {
		status = -1
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	qt, count, err := l.GetQualityTesting2(user_id, project_id, int((page_num-1)*count), int(count), status, bind_type, keyword)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(qt)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s,"total_count":%d}`, code, msg, string(b), count))
})

var GetFieldInspectionList2 = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取协作巡查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取协作巡查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "获取协作巡查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取协作巡查失败,project_id必须为整型"))
	}
	var page_num, count int64
	if c.Param("page_num") == "" {
		page_num = 1
	} else {
		page_num, err = strconv.ParseInt(c.Param("page_num"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取协作巡查失败，page_num必须为整型"))
		}
	}
	if c.Param("count") == "" {
		count = 20
	} else {
		count, err = strconv.ParseInt(c.Param("count"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取协作巡查失败，count必须为整型"))
		}
	}
	var keyword string
	if c.Param("keyword") != "" {
		keyword = strings.Trim(c.Param("keyword"), " ")
	}
	var status int
	if c.Param("status") != "" {
		status, err = strconv.Atoi(c.Param("status"))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "获取安全检查失败，status必须为整型"))
		}
	} else {
		status = -1
	}

	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "绑定项目资料失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	fi, count, err := l.GetFieldInspection2(user_id, project_id, int((page_num-1)*count), int(count), status, bind_type, keyword)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(fi)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s,"total_count":%d}`, code, msg, string(b), count))
})

var ViewSecurityCheck2 = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "查看安全检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,project_id必须为整型"))
	}
	if c.Param("sc_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "查看安全检查失败,请传入sc_id!"))
	}
	sc_id, err := strconv.ParseInt(c.Param("sc_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,sc_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看安全检查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ViewSecurityCheck2(user_id, project_id, sc_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ViewQualityTesting2 = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "查看质量检查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,project_id必须为整型"))
	}
	if c.Param("qt_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "查看质量检查失败,请传入qt_id!"))
	}
	qt_id, err := strconv.ParseInt(c.Param("qt_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,qt_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看质量检查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ViewQualityTesting2(user_id, project_id, qt_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var ViewFieldInspection2 = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("user_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "查看协作巡查失败,请传入user_id!"))
	}
	user_id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,user_id必须为整型"))
	}
	if c.Param("project_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,请传入project_id!"))
	}
	project_id, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,project_id必须为整型"))
	}
	if c.Param("fi_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 200, "查看协作巡查失败,请传入fi_id!"))
	}
	fi_id, err := strconv.ParseInt(c.Param("fi_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,fi_id必须为整型"))
	}
	var bind_type int
	if c.Param("bind_type") != "" {
		bind_type, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "查看协作巡查失败,bind_type必须为整型!"))
		}
	} else {
		bind_type = 1
	}
	data, err := l.ViewFieldInspection2(user_id, project_id, fi_id, bind_type)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)
	return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
})

var PostElementProperties = faygo.WrapDoc(
	faygo.HandlerFunc(func(c *faygo.Context) error {
		if c.Param("project_id") == "" {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 400, "提交Bim构件属性失败,请传入project_id!"))
		}
		projectId, err := strconv.ParseInt(c.Param("project_id"), 10, 64)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "提交Bim构件属性失败,project_id必须为整型"))
		}
		var bindType int
		if c.Param("bind_type") != "" {
			bindType, err = strconv.Atoi(strings.Trim(c.Param("bind_type"), " "))
			if err != nil {
				return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "提交Bim构件属性失败,bind_type必须为整型!"))
			}
		} else {
			bindType = 1
		}
		context := strings.Trim(c.Param("context"), " ")
		data, err := l.PostElementProperties(projectId, bindType, context)
		code, msg := u.AdapterError(err)
		if err != nil {
			return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
		}
		b, _ := json.Marshal(data)
		return c.String(200, fmt.Sprintf(`{"code":%d, "msg":"%s","data":%s}`, code, msg, string(b)))
	}),
	"post_element_properties",
	"提交Bim构件属性",
	// 定义参数
	faygo.ParamInfo{
		Name:     "project_id",
		In:       "query",
		Required: true,
		Model:    int(3), // API文档中显示的参数默认值
		Desc:     "项目编号",
	},
	faygo.ParamInfo{
		Name:     "bind_type",
		In:       "query",
		Required: true,
		Model:    int(3), // API文档中显示的参数默认值
		Desc:     "项目类型",
	},
	faygo.ParamInfo{
		Name:     "context",
		In:       "query",
		Required: true,
		Model:    string(""), // API文档中显示的参数默认值
		Desc:     "属性内容",
	},
)

var WeatherInfo = faygo.HandlerFunc(func(c *faygo.Context) error {
	if c.Param("city_id") == "" {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "city_id不能为空"))
	}
	cityId, err := strconv.ParseInt(c.Param("city_id"), 10, 64)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, 401, "city_id必须整型"))
	}

	data, err := l.GetWeatherInfo(cityId)
	code, msg := u.AdapterError(err)
	if err != nil {
		return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	}
	b, _ := json.Marshal(data)

	return c.String(200, fmt.Sprintf(`{"code":%d,"msg":"%s","data":%s}`, code, msg, string(b)))
})
