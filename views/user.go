package views

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/db"
	auth "server/utils"
)

/*
用户相关视图
*/

type UserToken struct {
	AccessToken string `json:"access_token"` // 用户token
	ExpireTime  int64  `json:"expire_time"`  // 过期时间
}

// 用户登陆视图
// 请求方法:POST
func Login(c *gin.Context) {
	// 获取post请求的表单中的用户名和密码
	username := c.PostForm("username")
	passwaord := c.PostForm("password")
	code := c.PostForm("code")
	if username == "" || passwaord == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "用户名或密码不能为空",
		})
		return
	}
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "验证码不能为空",
		})
		return
	}
	// 先验证验证码
	codeID, err := c.Cookie("CODEID")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "错误的请求",
		})
		return
	}
	if !VerifyCaptchaHandler(codeID, code) {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "验证码错误",
		})
		return
	}
	// 再校验用户
	// 用户、密码解密
	decode_name, _ := auth.DecryptData(username)
	decode_passwd, _ := auth.DecryptData(passwaord)
	// 校验用户
	current_user, isExist, err := db.CheckUser(decode_name, decode_passwd)
	if err != nil || !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "用户名或密码错误",
		})
		return
	}
	var t UserToken
	t.AccessToken, _ = createToken(current_user.Uuid, current_user.Name)
	t.ExpireTime = 24 * 60 * 60
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   t,
		"msg":    "认证成功",
	})
}

// 创建用户
type NewUser struct {
	UserName   string `json:"username" bind:"required"` // 用户名
	Department string `json:"department"`               // 部门
	Job        string `json:"job"`                      // 工作职位
}

func AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "创建用户成功",
	})
}

// 用户获取公钥视图
// 请求方法:GET
func ObtainKey(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   auth.GetPublicKey(),
		"msg":    "请求成功",
	})
}

// 用户退出视图
// 请求方法:GET
func LogOut(c *gin.Context) {
	// 清除缓存中的token
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "请求成功",
	})
}

// 用户修改密码视图
// 请求方法:POST
type NewPassword struct {
	OldPasswd     string `json:"old_password" bind:"required"`
	NewPasswd     string `json:"new_password" bind:"required"`
	ConfirmPasswd string `json:"confirm_password" bind:"required"`
}

func ChangePassword(c *gin.Context) {
	// 获取用户信息
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "认证失败",
		})
		return
	}
	var n NewPassword
	err := c.BindJSON(&n)
	if err != nil || !auth.Check(n) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "错误的请求",
		})
		return
	}
	if n.OldPasswd == n.NewPasswd {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "新旧密码不能一样",
		})
		return
	}
	if n.NewPasswd != n.ConfirmPasswd {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "新密码与确认密码不一致",
		})
		return
	}
	// 修改密码
	updateUserId := userID.(string)
	isMatched, _ := db.MatchPassword(updateUserId, n.OldPasswd)
	if !isMatched {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "旧密码不正确",
		})
		return
	}
	new_user := db.User{
		Password: n.NewPasswd,
	}
	if err := db.UpdateUserByUuid(updateUserId, new_user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "密码修改失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "密码修改成功",
	})
}

// 用户修改资料
// 请求方法:POST
type UserInfo struct {
	RealName  string `json:"real_name"`  // 真实姓名
	Sex       int8   `json:"sex"`        // 性别
	Email     string `json:"email"`      // 邮箱地址
	AvatarUrl string `json:"avatar_url"` // 头像地址
}

func ModifyInfo(c *gin.Context) {
	var u UserInfo
	err := c.BindJSON(&u)
	if err != nil || !auth.Check(u) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"data":   "",
			"msg":    "错误的请求",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   "",
		"msg":    "修改成功",
	})
}
