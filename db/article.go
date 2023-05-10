package db

type Article struct {
	Uuid       string `gorm:"varchar(32);index:'article_index';primary_key"`
	Title      string
	Author     string `gorm:"varchar(32)"`
	Content    string `gorm:"varchar(32)"`
	IsAllRead  string `gorm:"varchar(2);default:'1'"`
	IsAllWrite string `gorm:"varchar(2);default:'0'"`
	CreateTime string `gorm:"varchar(20)"`
	UpdateTime string `gorm:"varchar(20)"`
}

type ArticlePermission struct {
	UserID     string `gorm:"varchar(32)"`
	ArticleID  string `gorm:"varchar(32)"`
	Read       string `gorm:"varchar(2)"`
	Write      string `gorm:"varchar(2)"`
	UpdateTime string `gorm:"varchar(20)"`
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
