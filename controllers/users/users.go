package users

import (
	"github.com/gin-gonic/gin"
	"go-easy-admin/common/global"
	"go-easy-admin/models"
	"go-easy-admin/service/users"
)

// 用户注册

func Register(ctx *gin.Context) {
	params := new(models.User)
	if err := ctx.ShouldBind(&params); err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	err := users.NewUserInfo().Register(params)
	if err != nil {
		global.TPLogger.Error(err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "用户注册成功！！！")
	return
}

// 用户详情

func GetUserInfo(ctx *gin.Context) {
	idStr, _ := ctx.GetQuery("id")
	data, err := users.NewUserInfo().UserInfo(idStr)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}

// 用户列表

func ListUser(ctx *gin.Context) {
	params := new(struct {
		Name  string `form:"name"`
		Limit int    `form:"limit"`
		Page  int    `form:"page"`
	})
	if err := ctx.ShouldBindQuery(&params); err != nil {
		global.TPLogger.Error("用户查询数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	data, err := users.NewUserInfo().UserList(params.Name, params.Limit, params.Page)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", data)

}
