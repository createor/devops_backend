package views

import (
	"net/http"

	"server/db"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type NewArticle struct {
	Title   string `json:"title" bind:"required"`   // 文章标题
	Content string `json:"content" bind:"required"` // 文章内容
}

// 添加文章接口
func AddArticle(c *gin.Context) {
	userID, ok := c.Get("user_id") // 获取用户id
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
	articleID := utils.NewUuid()
	contentID := utils.GetFileNameByRandom()
	isCreate, _ := utils.SaveArticle(utils.FilterArticle(a.Content), contentID)
	err = db.CreateArticle(db.Article{
		Uuid:    articleID,
		Title:   a.Title,
		Author:  author,
		Content: contentID,
	})
	if err != nil || !isCreate {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "文章创建失败",
		})
		return
	}
	_ = db.CreateArticlePermission(db.ArticlePermission{
		UserID:    author,
		ArticleID: articleID,
		Read:      "1",
		Write:     "1",
	})
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "文章创建成功",
	})
}

// 获取用户能阅读的所有文章
func GetAllArticleByUser(c *gin.Context) {
	userID, ok := c.Get("user_id") // 获取用户id
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "认证失败",
		})
		return
	}
	articleList, err := db.QueryArticleByUser(userID.(string))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   articleList,
		"msg":    "查询成功",
	})
}

// 用户浏览文章
func BrowseArticle(c *gin.Context) {
	// 获取用户id
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "认证失败",
		})
		return
	}
	// 获取文章id
	articleID := c.Param("articleID")
	// 查询权限
	isAllow := db.QueryArticleByContent(userID.(string), articleID)
	if !isAllow {
		c.JSON(http.StatusForbidden, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "无权限阅读",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   utils.ShowArtcle(articleID),
		"msg":    "查询成功",
	})
}
