package db

type Article struct {
	Uuid      string `gorm:"type:varchar(32);index:'article_index';primary_key" json:"id"` // 文章的唯一标识
	Title     string `gorm:"type:varchar(255)" json:"title"`                               // 文章标题
	Author    string `gorm:"type:varchar(32)" json:"author"`                               // 文章作者，用户的uuid
	Content   string `gorm:"type:varchar(32)"`                                             // 文章的内容，内容的uuid
	IsAllRead string `gorm:"type:varchar(2);default:'1'"`                                  // 是否所有人可读,1-是,0-否,默认所有人可读
	// 是否所有人可编辑,1-是,0-否,默认所有人不可编辑
	// 如果所有人可编辑,那么必须所有人可读也就是把is_all_read更新为1
	IsAllWrite string `gorm:"type:varchar(2);default:'0'"`
	CreateTime string `gorm:"type:varchar(20)"` // 创建时间
	Updateor   string `gorm:"type:varchar(32)"` // 最后一次更新的人的uuid
	UpdateTime string `gorm:"type:varchar(20)"` // 更新时间
}

type ArticlePermission struct {
	UserID       string `gorm:"type:varchar(32)"` // 用户的uuid
	ArticleID    string `gorm:"type:varchar(32)"` // 文章的uuid
	DepartmentID string `gorm:"type:varchar(32)"` // 部门的uuid
	Read         string `gorm:"type:varchar(2)"`  // 读取权限,1-可读,0-不可读
	Write        string `gorm:"type:varchar(2)"`  // 编辑权限,1-可编辑,0-不可编辑
	Updateor     string `gorm:"type:varchar(32)"` // 更新的用户的uuid
	UpdateTime   string `gorm:"type:varchar(20)"` // 更新时间
}

// CreateArticle 创建文章
func CreateArticle(articleInfo Article) error {
	tx := DB.Table("article").Begin()
	articleInfo.CreateTime = NewTableTime()
	if err := tx.Create(&articleInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	// 创建此文章的人拥有可读可写权限，写入permission表中
	return nil
}

// CreateArticlePermission 创建文章权限
func CreateArticlePermission(permission ArticlePermission) error {
	tx := DB.Table("article_permission").Begin()
	permission.UpdateTime = NewTableTime()
	if err := tx.Create(&permission).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func UpdateArticlePermission(permission ArticlePermission) error {
	tx := DB.Table("article_permission").Begin()
	if permission.Write == "1" { // 如果用户拥有可写权限那么也应该拥有可读权限
		permission.Read = "1"
	}
	if err := tx.Create(&permission).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 根据用户id查询此用户能看到的文章列表
func QueryArticleByUser(userID string) ([]Article, error) {
	artcileList := make([]Article, 0)
	if err := DB.Table("article").Select("uuid,title").Where("is_all_read = '1'").Find(&artcileList).Error; err != nil {
		return nil, err
	}
	// if err := DB.Table("").Where().Find().Error; err != nil {

	// }
	// 合并列表并去重
	return artcileList, nil
}

// 查询用户是否有编辑文章的权限
func QueryArticle(userID, articleID, departmentID string) (bool, error) {
	var isEdit int
	// 先查询此文章是否所有人可编辑
	if err := DB.Table("article").Where("uuid = ? AND is_all_write = '1'", articleID).Count(&isEdit).Error; err != nil {
		return false, err
	}
	if isEdit == 0 {
		// 查询文章权限表
		if userID != "" {
			if err := DB.Table("artcile_permission").Where("article_id = ? AND user_id = ? AND write = '1'").Count(&isEdit).Error; err != nil {
				return false, err
			}
			if isEdit != 0 {
				return true, nil
			}
		}
		if departmentID != "" {
			if err := DB.Table("artcile_permission").Where("article_id = ? AND department_id = ? AND write = '1'").Count(&isEdit).Error; err != nil {
				return false, err
			}
			if isEdit != 0 {
				return true, nil
			}
		}
	} else {
		return true, nil
	}
	return false, nil
}


func QueryArticleByContent(userID,contentID string) bool {
	return true
}