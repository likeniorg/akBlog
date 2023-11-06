package public

import (
	filehashchecking "akBlog/pkg/fileHashChecking"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 公开的web服务
func Server() http.Handler {
	// 使用默认中间件(Logger和Recovery)
	r := gin.Default()

	// 主页头像
	r.StaticFile("/header.jpeg", "web/header.jpeg")

	//引入网站所需静态资源
	r.Static("/assets", "web/assets")

	// 显示mirrors目录
	r.StaticFS("/static/files/", http.Dir("web/files"))

	// 显示静态页面
	WebPage(r)

	filehashchecking.CreateHash()
	return r
}
