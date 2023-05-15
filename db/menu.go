package db

type Menu struct {
	Uuid       string `gorm:"type:varchar(32);index:'menu_index';primary_key" json:"id"`
	MenuID     string `gorm:"type:varchar(20);unique" json:"menu_id"` // 菜单id
	MenuName   string `gorm:"type:varchar(20)" json:"menu_name"`      // 菜单名称
	MenuPath   string `gorm:"type:varchar(20)" json:"menu_path"`      // 菜单路径
	MenuIcon   string `gorm:"type:varchar(20)" json:"menu_icon"`      // 菜单图标
	ParentMenu string `gorm:"type:varchar(32)"`                       // 上一级菜单
	Level      string `gorm:"type:varchar(2)"`
	Createor   string `gorm:"type:varchar(32)" json:"-"`
	CreateTime string `gorm:"type:varchar(20)" json:"-"`
	Updateor   string `gorm:"type:varchar(32)" json:"-"`
	UpdateTime string `gorm:"type:varchar(20)" json:"-"`
}

type MenuWithRole struct {
	MenuID     string `gorm:"type:varchar(32)"`
	RoleID     string `gorm:"type:varchar(32)"`
	Updateor   string `gorm:"type:varchar(32)"`
	UpdateTime string `gorm:"type:varchar(20)"`
}

// 根据角色查询所有菜单
func QueryAllMenuByRole(roleID string) {

}
