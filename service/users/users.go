package users

import (
	"errors"
	"go-easy-admin/common/global"
	"go-easy-admin/dao"
	"go-easy-admin/models"
	"strconv"
)

// 用户注册

func (u *userInfo) Register(user *models.User) error {
	err := dao.NewUserInterface().Register(user)
	if err != nil {
		global.TPLogger.Error("用户注册失败：", err)
		return errors.New("用户注册失败")
	}
	return err
}

// 用户详情

func (u *userInfo) UserInfo(id string) (*models.User, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		global.TPLogger.Error("用户详情查看失败：", err)
		return nil, errors.New("用户详情查看失败")
	}
	data, err := dao.NewUserInterface().UserInfo(uint(idInt))
	if err != nil {
		global.TPLogger.Error("用户详情查看失败：", err)
		return nil, errors.New("用户详情查看失败")
	}
	return data, nil
}

// 用户列表

func (u *userInfo) UserList(username string, limit, page int) (*models.UserList, error) {
	data, err := dao.NewUserInterface().UserList(username, limit, page)
	if err != nil {
		global.TPLogger.Error("用户列表查询失败：", err)
		return nil, errors.New("用户列表查询失败")
	}
	return data, nil
}