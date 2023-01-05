package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "net/http/pprof"
	"web_app/controller"
	"web_app/docs"
	"web_app/logger"
	"web_app/middleware"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		// 如果是发布模式 控制台不输出调试信息
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)
	// JWT认证中间件
	v1.Use(middleware.JWTAuthMiddleware())

	{
		// 所有社区
		v1.GET("/community", controller.CommunityHandler)
		// 社区详情
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		// 创建贴子
		v1.POST("/post", controller.CreatePostHandler)
		// 获取贴子详情
		v1.GET("/post/:id", controller.PostDetailHandler)
		// 获取帖子列表
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数或社区获取帖子获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)

		// 投票功能
		v1.POST("/vote", controller.PostVoteController)

	}

	pprof.Register(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
