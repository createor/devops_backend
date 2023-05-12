package db

/*
部门表
*/

type Department struct {
	Uuid             string `gorm:"type:varchar(32);index:'department_index';primary_key"`
	DepartmentID     string `gorm:"type:varchar(20);unique"`
	DepartmentName   string `gorm:"type:varchar(20)"` // 部门名称
	ParentDepartment string `gorm:"type:varchar(32)"` // 上一级部门的uuid
	Level            string `gorm:"type:varchar(10)"` // 等级
	Createor         string `gorm:"type:varchar(32)"` // 创建者
	CreateTime       string `gorm:"type:varchar(20)"` // 创建时间
	Updateor         string `gorm:"type:varchar(32)"`
	UpdateTime       string `gorm:"type:varchar(20)"`
}

type DepartmentWithUser struct {
	DepartmentID string `gorm:"type:varchar(32)"`
	UserID       string `gorm:"type:varchar(32);unique"`
	Updateor     string `gorm:"type:varchar(32)"`
	UpdateTime   string `gorm:"type:varchar(20)"`
}

type DepartmentManage struct {
	DepartmentID string `gorm:"type:varchar(32)"`
	UserID       string `gorm:"type:varchar(32);unique"`
	Updateor     string `gorm:"type:varchar(32)"`
	UpdateTime   string `gorm:"type:varchar(20)"`
}

// 创建部门
func CreateDepartment(department Department) error {
	tx := DB.Table("department").Begin()
	if err := tx.Create(&department).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 创建部门与用户的关联
func CreateDepartmentWithUser(releation DepartmentWithUser) error {
	tx := DB.Table("department_with_user").Begin()
	if err := tx.Create(&releation).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 查询用户所属的部门
// func QueryDepartmentByUser(userID string) *Department {
// 	var d Department
// 	DB.Table("")
// 	return &d
// }

// 查询部门下的所有用户
func QueryAllUserByDepartment(departmentID string) ([]User, error) {
	result := make([]User, 0)
	if err := DB.Table("department a").Select("b.uuid, b.name").Joins("INNER JOIN (select t1.uuid,t1.name,t2.department_id from user t1 LEFT JOIN department_with_user t2 ON t1.uuid = t2.user_id) b ON a.uuid = b.department_id AND b.department_id = ?", departmentID).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// 绑定部门管理员，部门管理员只能是此部门的成员，可以管理本级及下级单位

// 修改用户部门，需先确定是否绑定部门管理员，绑定则需先解除绑定才能修改
