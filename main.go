package main

import (
	"fmt"
	"simple-block-api/config"
	"simple-block-api/handler"
	"simple-block-api/middleware"
	"simple-block-api/models"
	"simple-block-api/utils"

	"github.com/gin-gonic/gin"
)

func initTables() {
	db := utils.GetDB()
	db.AutoMigrate(models.User{}, models.Post{}, models.Comment{}, models.ErrorLog{})
}

func main() {
	config.InitViper()
	fmt.Println("==================== 配置文件加载成功 ====================")
	fmt.Printf("Host: %s\n", config.Cfg.Server.Host)
	fmt.Printf("Port: %d\n", config.Cfg.Server.Port)
	fmt.Printf("User: %s\n", config.Cfg.Database.User)
	fmt.Printf("Name: %s\n", config.Cfg.Database.Name)

	// initTables()
	startServer()
}

func startServer() {
	// 创建 Gin 引擎实例 engine
	r := gin.New()
	r.Use(middleware.MyLogger, middleware.MyRecovery(), middleware.MyJWT)

	v1 := r.Group("/api/v1")
	v1.GET("/health", func(ctx *gin.Context) {
		utils.RespWithMessage(ctx, "I'm alive.")
	})

	v1.POST("/login", handler.UserLogin)
	v1.POST("/register", handler.UserRegister)

	post := v1.Group("/post")
	post.POST("/list", handler.PostList)
	post.POST("/create", handler.PostCreate)
	post.GET("/detail/:id", handler.PostDetail)
	post.PUT("/update/:id", handler.PostUpdate)
	post.DELETE("/delete/:id", handler.PostDelete)

	comment := v1.Group("/comment")
	comment.POST("/list", handler.CommentList)
	comment.POST("/create", handler.CommentCreate)

	addr := config.Cfg.Server.Host + ":" + fmt.Sprint(config.Cfg.Server.Port)
	fmt.Printf("服务器正在运行，监听地址: %s\n", addr)
	r.Run(addr)
}
