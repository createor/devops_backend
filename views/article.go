package views

import (
	"net/http"

	"server/db"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type NewArticle struct {
	Title   string `json:"title" bind:"required"`
	Content string `json:"content" bind:"required"`
}

func CreateArticle(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "认证失败",
		})
		return
	}
	var a NewArticle
	err := c.BindJSON(&a)
	if err != nil || !utils.Check(a) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "标题或内容不能为空",
		})
		return
	}
	author := userID.(string)
	err = db.CreateArticle(db.Article{
		Uuid:   utils.NewUuid(),
		Title:  a.Title,
		Author: author,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "文章创建失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "文章创建成功",
	})
}
