package db

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Role struct {
	Uuid       string `gorm:"type:varchar(32);index:'role_index';primary_key" json:"id"` // uuid,角色的唯一标识
	RoleID     string `gorm:"type:varchar(20);column:role_id;unique" json:"role_id"`     // 角色id
	RoleName   string `gorm:"type:varchar(20)" json:"role_name"`                         // 角色名称
	Createor   string `gorm:"type:varchar(32)" json:"-"`                                 // 创建人,用户的uuid
	CreateTime string `gorm:"type:varchar(20)" json:"-"`                                 // 创建时间
	Updateor   string `gorm:"type:varchar(32)" json:"-"`                                 // 更新人,用户的Uuid
	UpdateTime string `gorm:"type:varchar(20)" json:"-"`                                 // 更新时间
}

type RoleWithUser struct {
	RoleID     string `gorm:"type:varchar(32);column:role_id"` // 角色的uuid
	UserID     string `gorm:"type:varchar(32);column:user_id"` // 用户的uuid
	Updateor   string `gorm:"type:varchar(32)"`                // 更新人,用户的uuid
	UpdateTime string `gorm:"type:varchar(20)"`                // 更新时间
}

// 根据关键字查询角色数据
func FilterRole(where ...any) (*Role, error) {
	var d Role
	err := DB.Table("role").First(&d, where...).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &d, nil
}

// 创建角色
func CreateRole(userID string, roleInfo Role) (bool, error) {
	isExist, err := FilterRole(Role{RoleID: roleInfo.RoleID})
	if err != nil {
		return false, err
	}
	if isExist != nil {
		return true, nil
	}
	tx := DB.Table("role").Begin()
	roleInfo.Createor = userID
	roleInfo.CreateTime = NewTableTime()
	if err = tx.Create(&roleInfo).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return false, nil
}

func FilterRoleWithUser(info RoleWithUser) (bool, error) {
	var c int
	if err := DB.Table("role_with_user").Where(info).Count(&c).Error; err != nil {
		return false, err
	}
	if c != 0 {
		return true, nil
	}
	return false, nil
}

// 添加用户与角色的关联
func CreateRoleWithUser(userID string, info RoleWithUser) (bool, error) {
	isExist, err := FilterRoleWithUser(info)
	if err != nil {
		return false, err
	}
	if isExist {
		return true, nil
	}
	tx := DB.Table("role_with_user").Begin()
	info.Updateor = userID
	info.UpdateTime = NewTableTime()
	if err = tx.Create(&info).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return false, nil
}

// 查找用户绑定的角色
func QueryRoleByUser(userID string) ([]Role, bool, error) {
	result := make([]Role, 0)
	err := DB.Table("role").Select("role.uuid, role.role_id, role.role_name").Joins("INNER JOIN role_with_user ON role.uuid = role_with_user.role_id and role_with_user.user_id = ?", userID).Find(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return result, false, nil
	} else if err != nil {
		return result, false, err
	}
	return result, true, nil
}
