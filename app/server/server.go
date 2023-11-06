package server

import (
	"akBlog/app/admin"
	"akBlog/app/config"
	"akBlog/app/public"
	"akBlog/app/util"
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 公开入口使用HTTP协议
func EntranceHTTP(port string) {
	// 不指定端口将使用默认端口
	if port == "" {
		port = config.Get("port")
	}
	server := http.Server{
		Addr:    port,
		Handler: public.Server(),
	}
	err := server.ListenAndServe()
	util.ErrprDisplay(err)
}

// 公开入口使用HTTPS协议
func EntranceHTTPS() {
	server := http.Server{
		Addr:    config.Get("port"),
		Handler: public.Server(),
	}

	// 证书名字规范 域名 + ".crt" 域名 + ".key"
	err := server.ListenAndServeTLS("config/cert/Ca/"+config.Get("domain")+".crt", "config/cert/Ca/"+config.Get("domain")+".key")
	util.ErrprDisplay(err)
}

// 管理员使用HTTP协议
func AdminHTTP() {
	server := http.Server{
		Addr:    config.Get("adminPort"),
		Handler: admin.AdminServer(),
	}
	err := server.ListenAndServe()
	util.ErrprDisplay(err)
}

// 管理员端口使用HTTPS协议
func AdminHttps() {
	gin.SetMode(gin.ReleaseMode)

	server := http.Server{
		Addr:    config.Get("adminPort"),
		Handler: admin.AdminServer(),
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	// 证书名字规范 域名 + ".crt" 域名 + ".key"
	err := server.ListenAndServeTLS("./config/cert/adminCa/"+config.Get("domain")+".crt", "./config/cert/adminCa/"+config.Get("domain")+".key")

	// 可能会阻塞输出错误
	util.ErrprDisplay(err)
}
