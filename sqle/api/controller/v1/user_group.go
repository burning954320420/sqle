package v1

import (
	"github.com/actiontech/sqle/sqle/api/controller"

	"github.com/labstack/echo/v4"
)

type CreateUserGroupReqV1 struct {
	Name       string   `json:"user_group_name" form:"user_group_name" example:"test" valid:"required,name"`
	Desc       string   `json:"user_group_desc" form:"user_group_desc" example:"this is a group"`
	Users      []string `json:"user_name_list" form:"user_name_list"`
	IsDisabled bool     `json:"is_disabled" form:"is_disabled"`
}

// @Summary 创建用户组
// @Description create user group
// @Id CreateUserGroupV1
// @Tags user_group
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param instance body v1.CreateUserGroupReqV1 true "create user group"
// @Success 200 {object} controller.BaseRes
// @router /v1/user_groups [post]
func CreateUserGroup(c echo.Context) (err error) {
	return JSONNewNotImplementedErr(c)
}

type GetUserGroupsResV1 struct {
	controller.BaseRes
	Data      []*UserGroupListItemResV1 `json:"data"`
	TotalNums uint64                    `json:"total_nums"`
}

type UserGroupListItemResV1 struct {
	ID         uint   `json:"id"`
	Name       string `json:"group_name"`
	Desc       string `json:"group_desc"`
	IsDisabled bool   `json:"is_disabled"`
}

// @Summary 获取用户组列表
// @Description get user group info list
// @Id GetUserGroupsV1
// @Tags user_group
// @Id getUserGroupListV1
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} v1.GetUserGroupsResV1
// @router /v1/user_groups [get]
func GetUserGroups(c echo.Context) (err error) {
	return JSONNewNotImplementedErr(c)
}
