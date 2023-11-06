package admin

import (
	"akBlog/app/config"
	"akBlog/app/mirrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func verifyIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.ClientIP() == config.Get("serverIP") || ctx.ClientIP() == "::1" {
			ctx.Next()
		} else {
			ctx.String(200, "sb no allow le")
			ctx.Abort()
		}
	}
}

// 管理员后台服务器
func AdminServer() http.Handler {
	// 使用默认路由
	adminR := gin.New()

	// 所有页面都要验证IP
	adminR.Use(verifyIP())

	// 开启镜像站文件上传功能
	mirrors.UploadMirrorFile(adminR)
	return adminR
}
