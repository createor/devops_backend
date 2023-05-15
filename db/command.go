package db

/*
远程主机后记录命令并过滤敏感命令，例如su等命令申请执行
*/

// 用户命令记录
type CommandRecord struct {
	ID         string `gorm:"pramry_key"`
	IP         string `gorm:"type:varchar(20)"`
	UserName   string `gorm:"type:varchar(20);column:username"`
	Command    string `gorm:"type:varchar(255)"` // 命令
	Createor   string `gorm:"type:varchar(32)"`
	CreateTime string `gorm:"type:varchar(20)"`
}

type ProhibitCommand struct {
	ID         string `gorm:"pramry_key"`
	Command    string `gorm:"type:varchar(255)"` // 禁止用户执行的命令，需要申请执行
	Createor   string `gorm:"type:varchar(32)"`
	CreateTime string `gorm:"type:varchar(20)"`
}

func CreateCommandRecord(c *CommandRecord) error {
	tx := DB.Table("command_record").Begin()
	if err := tx.Create(&c).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func CreateProhibitCommand(p *ProhibitCommand) error {
	tx := DB.Table("prohibit_command").Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
