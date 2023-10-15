package main

import (
	"fmt"
	"time"

	"syntax/internal/repository"
	"syntax/internal/repository/dao"
	"syntax/internal/service"
	"syntax/internal/web"
	"syntax/utils/middleware"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	db := initDatabase()
	server := initWebServer()
	db.Debug()
	initUser(db, server)

	if err := server.Run(":8882"); err != nil {
		fmt.Printf("服务器启动失败.请检查端口是否占用并检查配置文件!error: %v\n", err)
	}
}

func initUser(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDAO(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	sv := web.NewUserHandle(us)
	sv.RegisterRouter(server)
}

func initWebServer() *gin.Engine {
	sv := gin.Default()
	sv.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true, // 是否允许携带cookie
		MaxAge:           12 * time.Hour,
	}))
	// 初始化session
	store := cookie.NewStore([]byte("sse"))
	sv.Use(sessions.Sessions("ssid", store))
	// 使用中间件 过滤 session
	sessionMiddle := middleware.SessionMiddleware{}
	sv.Use(sessionMiddle.CheckLogin())

	return sv
}

func initDatabase() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/webook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("数据库连接失败!error: %v\n", err)
		return nil
	}
	err = dao.InitTables(db)
	if err != nil {
		fmt.Printf("数据库初始化表失败!error: %v\n", err)
		return nil
	}
	return db
}
