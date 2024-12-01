package main

/*
单体服务：Sync
*/
import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yanweidong/stat/node"
)

func main() {
	port := os.Getenv("STAT_PORT")
	if strings.TrimSpace(port) == "" {
		port = "9030"
	}
	log.Println("Starting stat service on port(ENV:STAT_PORT/default:9030) " + port)

	gin.SetMode(gin.ReleaseMode)

	// 禁用日志输出
	gin.DefaultWriter = io.Discard

	r := gin.Default()

	r.Use(gin.Recovery())
	// 配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                                                                            // 允许的源
		AllowMethods:     []string{"GET", "OPTIONS"},                                                                                               // 允许的方法
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Grpc-Response-Code", "Request-Id", ",Work-Space"}, // 允许的头部
		ExposeHeaders:    []string{"Content-Length"},                                                                                               // 暴露的头部
		AllowCredentials: true,                                                                                                                     // 是否允许发送 Cookie
		MaxAge:           12 * time.Hour,                                                                                                           // 预检请求的有效期
	}))

	//公共 不需要验证token
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, node.Stat())
	})

	// 默认启动的是 8080端口，也可以自己定义启动端口
	r.Run(":" + port)
}
