package models

import "time"

type BuildRole struct {
	Id                 int64     `xorm:"not null pk autoincr INT(11)" json:"id"`           //编号
	Name               string    `xorm:"VARCHAR(30)" json:"name"`                          //角色名称
	ParentId           int64     `xorm:"INT(11) index" json:"parent_id"`                   //上级编号
	IsConstructionTeam bool      `xorm:"default 0 TINYINT(1)" json:"is_construction_team"` //是否施工班组
	Seq                int       `xorm:"default 0 INT(11)" json:"seq"`                     //排序
	BindType           int       `xorm:"default 1 INT(4)" json:"bind_type"`                //绑定类型
	DeleteTime         time.Time `xorm:"deleted" json:"delete_time"`                       //删除时间
}

type DefaultBuildRole struct {
	Id         int64     `xorm:"not null pk autoincr INT(11)" json:"id"` //编号
	RoleId     int64     `xorm:"INT(11) index" json:"role_id"`           //角色编号
	ModelId    int64     `xorm:"INT(11) index" json:"model_id"`          //模块编号
	RightValue string    `xorm:"VARCHAR(10)" json:"right_value"`         //权限值，二进制字符串保存，顺序为可读、可新建、可修改、可删除、可打印、可分享，例子：001100
	CreateTime time.Time `xorm:"created" json:"create_time"`             //创建时间
	BindType   int       `xorm:"default 1 INT(4)" json:"bind_type"`      //绑定类型
}

type ProjectRoleByBuild struct {
	Id                 int64  `xorm:"not null pk autoincr INT(11)" json:"id"`           //编号
	Name               string `xorm:"VARCHAR(30)" json:"name"`                          //角色名称
	ParentId           int64  `xorm:"INT(11) index" json:"parent_id"`                   //上级编号
	ProjectId          int64  `xorm:"INT(11) index" json:"project_id"`                  //项目编号
	IsConstructionTeam bool   `xorm:"default 0 TINYINT(1)" json:"is_construction_team"` //是否施工班组
	DefaultBuildRole   int64  `xorm:"default -1 INT(11)" json:"default_build_role"`     //默认项目角色编号
	Seq                int    `xorm:"default 0 INT(11)" json:"seq"`                     //排序
	BindType           int    `xorm:"default 1 INT(4)" json:"bind_type"`                //绑定类型
}
type BuildProjectRole struct {
	Id         int64     `xorm:"not null pk autoincr INT(11)" json:"id"` //编号
	RoleId     int64     `xorm:"INT(11) index" json:"role_id"`           //角色编号
	ModelId    int64     `xorm:"INT(11) index" json:"model_id"`          //模块编号
	ProjectId  int64     `xorm:"INT(11) index" json:"project_id"`        //项目编号
	RightValue string    `xorm:"VARCHAR(10)" json:"right_value"`         //权限值，二进制字符串保存，顺序为可读、可新建、可修改、可删除、可打印、可分享，例子：001100
	CreateTime time.Time `xorm:"created" json:"create_time"`             //创建时间
	BindType   int       `xorm:"default 1 INT(4)" json:"bind_type"`      //绑定类型
}

type EmployeeRole struct {
	Id               int64     `xorm:"not null pk autoincr INT(11)" json:"id"`       //编号
	UserId           int64     `xorm:"INT(11) index" json:"user_id"`                 //用户编号
	ProjectId        int64     `xorm:"INT(11) index" json:"project_id"`              //项目编号
	BuildRoleId      int64     `xorm:"INT(11) index" json:"build_role_id"`           //角色编号
	IsManager        bool      `xorm:"default false BOOL" json:"is_manager"`         //是否管理人员
	DefaultBuildRole int64     `xorm:"default -1 INT(11)" json:"default_build_role"` //默认项目角色编号
	BindType         int       `xorm:"default 1 INT(4)" json:"bind_type"`            //绑定类型
	DeleteTime       time.Time `xorm:"deleted" json:"delete_time"`                   //删除时间
}

func GetProjectRoleById(id int64) string {
	p := new(ProjectRoleByBuild)
	h, err := DB.Id(id).Get(p)
	if err != nil {
		return ""
	}
	if !h {
		return ""
	}
	return p.Name
}

func GetUserRoleByProjectId(user_id, project_id int64, bind_type int) string {
	er := new(EmployeeRole)
	h, err := DB.Where("project_id = ? and bind_type = ? and user_id = ?", project_id, bind_type, user_id).Get(er)
	if err != nil {
		return ""
	}
	if !h {
		return ""
	}
	return GetProjectRoleById(er.BuildRoleId)
}

func GetUserRoleIdByProjectId(user_id, project_id int64, bind_type int) int64 {
	er := new(EmployeeRole)
	h, err := DB.Where("project_id = ? and bind_type = ? and user_id = ?", project_id, bind_type, user_id).Get(er)
	if err != nil {
		return -1
	}
	if !h {
		return -1
	}
	return er.BuildRoleId
}

func GetUserParentRoleIdByRoleId(role_id int64) int64 {
	p := new(ProjectRoleByBuild)
	h, err := DB.Id(role_id).Get(p)
	if err != nil {
		return -1
	}
	if !h {
		return -1
	}
	return p.ParentId
}

func GetDefaultRoleParentId(user_id, project_id int64, bind_type int) int64 {
	role_id := GetUserRoleIdByProjectId(user_id, project_id, bind_type)
	parent_id := GetUserParentRoleIdByRoleId(role_id)
	p := new(ProjectRoleByBuild)
	h, err := DB.Id(parent_id).Get(p)
	if err != nil {
		return -1
	}
	if !h {
		return -1
	}
	return p.DefaultBuildRole
}
