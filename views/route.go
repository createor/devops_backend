package views

import "github.com/gin-gonic/gin"

// 路由

func InitRouter() *gin.Engine {
	// 设置为release模式
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// 使用中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(authMiddleware())
	v1 := r.Group("/api")
	{
		// 用户路由
		r_user := v1.Group("/user")
		{
			r_user.GET("/getkey", ObtainKey)
			r_user.POST("/login", Login)
			r_user.POST("/changepasswd", ChangePassword)
		}
		// 角色路由
		r_role := v1.Group("/role")
		{
			r_role.GET("/user", GetUserRole)
			r_role.POST("/add", AddRole)
			r_role.POST("/bind", BindRole)
		}
		r_department := v1.Group("/department")
		{
			r_department.GET("/user", GetAllUserFromDepartment)
		}
	}

	return r
}
