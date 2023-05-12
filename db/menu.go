package db

type Menu struct {
	Uuid       string `gorm:"type:varchar(32);index:'menu_index';primary_key"`
	MenuID     string `gorm:"type:varchar(20);unique"` // 菜单id
	MenuName   string `gorm:"type:varchar(20)"`        // 菜单名称
	MenuPath   string `gorm:"type:varchar(20)"`        // 菜单路径
	MenuIcon   string `gorm:"type:varchar(20)"`        // 菜单图标
	ParentMenu string `gorm:"type:varchar(32)"`        // 上一级菜单
	Level      string `gorm:"type:varchar(2)"`
	Createor   string `gorm:"type:varchar(32)"`
	CreateTime string `gorm:"type:varchar(20)"`
	Updateor   string `gorm:"type:varchar(32)"`
	UpdateTime string `gorm:"type:varchar(20)"`
}

type MenuWithRole struct {
	MenuID     string `gorm:"type:varchar(32)"`
	RoleID     string `gorm:"type:varchar(32)"`
	Updateor   string `gorm:"type:varchar(32)"`
	UpdateTime string `gorm:"type:varchar(20)"`
}
