package db

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// 用户表
type User struct {
	Uuid         string `gorm:"type:varchar(32);index:'user_index';primary_key"` // uuid，用户唯一标识
	Name         string `gorm:"type:varchar(20);unique"`                         // 用户名,未加密
	RealName     string `gorm:"type:varchar(20);unique;column:realname"`         // 用户真实姓名
	UserName     string `gorm:"type:varchar(32);column:username"`                // 用户名,MD5加密后
	Password     string `gorm:"type:varchar(32)"`                                // 密码,MD5加密后
	Sex          string `gorm:"type:varchar(2);default:'1'"`                     // 性别,1-男,2-女
	Status       string `gorm:"type:varchar(2);default:'0'"`                     // 状态,1-启用，0-禁用
	AreaProvince string `gorm:"varchar(10)"`                                     // 地区省份
	AreaCity     string `gorm:"varchar(10)"`                                     // 地区城市
	CreateTime   string `gorm:"varchar(20)"`                                     // 创建时间
	UpdateTime   string `gorm:"varchar(20)"`                                     // 更新时间
	AvatarImg    string // 用户头像地址
}

func InitUser(newUser any) {
	// 插入User用户表

	// 插入Department部门表
}

// CreateUser 创建用户
//
// 参数:
//
//	userInfo: User,用户信息
//
// 返回:
//
//	error: 错误信息
func CreateUser(userInfo User) error {
	tx := DB.Table("user").Begin() // 开启事务
	userInfo.CreateTime = NewTableTime()
	if err := tx.Create(&userInfo).Error; err != nil {
		tx.Rollback() // 回滚
		return err
	}
	tx.Commit() // 提交事务
	return nil
}

// CheckUser 根据用户名密码校验用户
//
// 参数:
//
//	username: string, 用户名
//	password: string, 密码
//
// 返回:
//
//	*User: 用户信息
//	bool: 是否为空
//	error: 错误信息
func CheckUser(username, password string) (*User, bool, error) {
	var u User
	err := DB.Table("user").Where("username = ? AND password = ? AND status = '1'", username, password).Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // 查询为空时
		return nil, false, nil
	} else if err != nil { // 查询错误时
		return nil, false, err
	}
	return &u, true, nil
}

// UpdateUserByUuid 根据Uuid更新用户信息
//
// 参数:
//
//  uuid: string, 用户的唯一id
//  userInfo, User, 用户信息
//
// 返回:
//
//  error: 错误信息
func UpdateUserByUuid(uuid string, userInfo User) error {
	tx := DB.Table("user").Begin()
	userInfo.UpdateTime = NewTableTime()
	if err := tx.Where("uuid = ?", uuid).Updates(userInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// MatchPassword 根据uuid和密码匹配用户
//
// 参数:
//
//  uuid: string, 用户的唯一标识
//  password: string, 用户的密码
//
// 返回:
//
//  bool: 是否匹配
//  error: 错误信息
func MatchPassword(uuid, password string) (bool, error) {
	var count int
	result := DB.Table("user").Where("uuid = ? AND password = ?", uuid, password).Count(&count)
	if err := result.Error; err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
