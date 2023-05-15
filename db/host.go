package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Host struct {
	Uuid       string `gorm:"type:varchar(32);index:'host_index';primary_key" json:"id"`
	Ip         string `gorm:"type:varchar(20);unique" json:"ip"`
	Port       string `gorm:"type:varchar(10)" json:"port"`
	HostName   string `gorm:"type:varchar(20);column:hostname" json:"hostname"`
	Createor   string `gorm:"type:varchar(32)" json:"-"`
	CreateTime string `gorm:"type:varchar(20)" json:"-"`
	Updateor   string `gorm:"type:varchar(32)" json:"-"`
	UpdateTime string `gorm:"type:varchar(20)" json:"-"`
	UserList   []HostManyUser
}

type HostManyUser struct {
	HostID   string `gorm:"type:varchar(32)"`
	UserName string `gorm:"type:varchar(20);column:username"`
	Password string `gorm:"type:varchar(20)"`
}

type HostWithUser struct {
	HostID     string `gorm:"type:varchar(32)"`
	UserName   string `gorm:"type:varchar(20);column:username"`
	UserID     string `gorm:"type:varchar(32)"`
	Updateor   string `gorm:"type:varchar(32)"`
	UpdateTime string `gorm:"type:varchar(20)"`
}

type SingleHost struct {
	Ip       string
	Port     string
	UserName string
	Password string
}

type ManyHost struct {
	Uuid     string   `json:"id,omitempty"`
	Ip       string   `json:"ip,omitempty"`
	HostName string   `json:"hostname,omitempty"`
	UserName string   `gorm:"column:username" json:"-"`
	UserList []string `json:"user,omitempty"`
}

func MergeHost(m []ManyHost) []ManyHost {
	fmt.Println(m)
	nm := make([]ManyHost, 0)
	n := make(map[string]any)
	for i := 0; i < len(m); i++ {
		m[i].UserList = append(m[i].UserList, m[i].UserName)
		for j := i + 1; j < len(m); j++ {
			_, ok := n[m[i].Uuid]
			if !ok {
				if m[i].Uuid == m[j].Uuid {
					m[i].UserList = append(m[i].UserList, m[j].UserName)
				}
			}
		}
		if _, ok := n[m[i].Uuid]; !ok {
			n[m[i].Uuid] = struct{}{}
			nm = append(nm, m[i])
		}
	}
	return nm
}

// 插入主机数据
func CreateHost() error {
	return nil
}

// 查询用户拥有的主机列表
func QueryAllHostByUser(userID string) ([]ManyHost, error) {
	h := make([]ManyHost, 0)
	err := DB.Table("host a").Select("a.uuid,a.ip,a.hostname,b.username").Joins("INNER JOIN host_with_user b ON a.uuid = b.host_id AND b.user_id = ?", userID).Find(&h).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return h, nil
}

// 根据主机ip、用户名查找用户是否拥有权限
func QueryHostByName(host, name, userID string) (bool, error) {
	var c int
	if err := DB.Table("host a").Joins("INNER JOIN host_with_user b ON a.uuid = b.host_id AND a.ip = ? AND b.username = ? AND b.user_id = ?", host, name, userID).Count(&c).Error; err != nil {
		return false, err
	}
	if c == 0 {
		return false, nil
	}
	return true, nil
}

// 查找主机
func QueryHost(where ...any) (*SingleHost, error) {
	var s SingleHost
	err := DB.Table("host a").Select("a.ip,a.port,b.username,b.password").Joins("INNER JOIN host_many_user b ON a.uuid = b.host_id").First(&s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &s, nil
}
