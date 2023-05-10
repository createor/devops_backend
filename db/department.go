package db

/*
部门表
*/

type Department struct {
	Uuid           string `gorm:"varchar(32);index:'department_index';primary_key"`
	DepartmentName string `gorm:"varchar(20)"` // 部门名称
}
