package public

import (
	"akBlog/app/util"
	filehashchecking "akBlog/pkg/fileHashChecking"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// 前端页面生成
func WebPage(r *gin.Engine) {
	// 加载mirrors文件
	r.LoadHTMLFiles("web/protect/mirrors.html")
	r.StaticFile("/scss/style.min.abbd69b2908fdfcd5179898beaafd374514a86538d81639ddd2c58c06ae54e40.css", "web/scss/style.min.abbd69b2908fdfcd5179898beaafd374514a86538d81639ddd2c58c06ae54e40.css")
	r.StaticFile("/ts/main.js", "web/ts/main.js")

	r.GET("/", func(ctx *gin.Context) {
		ctx.File("./web/index.html")
	})

	recursionMakeLink(r, "about")
	recursionMakeLink(r, "archives")
	recursionMakeLink(r, "search")
	recursionMakeLink(r, "links")
	recursionMakeLink(r, "categories")
	recursionMakeLink(r, "tags")
	recursionMakeLink(r, "p")
	recursionMakeLink(r, "page")
	recursionMakeLink(r, "post")

	r.GET("/404.html", func(ctx *gin.Context) {
		ctx.File("./web/404.html")
	})

	r.GET("/mirrors", func(ctx *gin.Context) {
		success, fail := filehashchecking.CheckingHash()
		ctx.HTML(200, "mirrors.html", gin.H{"sucess": success, "fail": fail})
	})

}

// 递归制造HTML链接，递归不需要传递完整目录
func recursionMakeLink(r *gin.Engine, Dirpath string) {
	// 当前目录
	currentPath := filepath.Join("web", Dirpath)

	// 读取当前文件夹
	dir, err := os.ReadDir(currentPath)
	util.ErrprDisplay(err)

	for _, v := range dir {
		if v.IsDir() {
			recursionMakeLink(r, filepath.Join(Dirpath, v.Name()))
		} else {
			if v.Name() == "index.html" {
				r.GET(filepath.Join("/", Dirpath, "/"), func(ctx *gin.Context) {
					ctx.File(filepath.Join("./", currentPath, "index.html"))
				})
			}
		}
	}

}
