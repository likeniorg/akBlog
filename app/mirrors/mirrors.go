package mirrors

import (
	filehashchecking "akBlog/pkg/fileHashChecking"

	"github.com/gin-gonic/gin"
)

// 开启显示页面和文件上传功能
func EnableALLFeatures(r *gin.Engine) {
	DisplayHTMLPage(r)

	UploadMirrorFile(r)
}

// 显示镜像站页面
func DisplayHTMLPage(r *gin.Engine) {
	// 加载镜像站文件
	r.LoadHTMLFiles("web/protect/mirrors.html")

	r.GET("/mirrors", func(ctx *gin.Context) {

		// 验证文件哈希值
		success, fail := filehashchecking.CheckingHash()

		// 返回网页验证记录
		ctx.HTML(200, "mirrors.html", gin.H{"success": success, "fail": fail})
	})

}

// 上传文件到镜像站
func UploadMirrorFile(r *gin.Engine) {

	// 文件上传
	r.POST("/saveFile", func(ctx *gin.Context) {
		// 单独文件上传
		file, _ := ctx.FormFile("file")

		// 写入HASH
		filehashchecking.AddHash(file.Filename)

		//保存文件
		if err := ctx.SaveUploadedFile(file, filehashchecking.ScanDirPath+file.Filename); err != nil {
			ctx.Redirect(200, "/mirrors")
		} else {
			ctx.String(500, "上传出错："+err.Error())
		}

	})

}
