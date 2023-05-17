package views

import "github.com/gin-gonic/gin"

/*
路由
*/

func InitRouter() *gin.Engine {
	// 设置为release模式
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// 使用中间件
	writer, err := logMiddleware()
	if err == nil {
		gin.DefaultWriter = writer
	}
	r.Use(gin.Logger()) // 认证中间件
	r.Use(gin.Recovery())
	r.Use(authMiddleware()) // 鉴权中间件
	v1 := r.Group("/api")
	{
		// 验证码路由
		r_code := v1.Group("captcha")
		{
			r_code.GET("/code", GenerateCaptchaHandler) // 获取验证码图片
		}
		// 用户路由
		r_user := v1.Group("/user")
		{
			r_user.GET("/getkey", ObtainKey)             // 获取公钥
			r_user.POST("/login", Login)                 // 用户登陆接口
			r_user.POST("/changepasswd", ChangePassword) // 用户修改密码接口
			r_user.GET("/logout", LogOut)                // 用户登出接口
			r_user.POST("/add", AddUser)                 // 添加用户接口
		}
		// 角色路由
		r_role := v1.Group("/role")
		{
			r_role.GET("/user", GetUserRole) // 获取所有
			r_role.POST("/add", AddRole)     // 添加角色接口
			r_role.POST("/bind", BindRole)   // 绑定用户角色接口
		}
		// 部门路由
		r_department := v1.Group("/department")
		{
			r_department.GET("/user", GetAllUserFromDepartment)
		}
		// 文章路由
		r_article := v1.Group("/article")
		{
			r_article.POST("/add", AddArticle)               // 创建文章接口
			r_article.GET("/page/:articleID", BrowseArticle) // 查看文章接口
		}
		// 主机路由
		r_host := v1.Group("/host")
		{
			r_host.GET("/all", GetAllHostBySelf)
		}
		// ws路由
		r_ws := v1.Group("/ws")
		{
			r_ws.GET("/ssh", NewWebsocketHandler(PreloadConnect, ConnectSSH, true))
		}
	}

	return r
}
