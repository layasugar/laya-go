package routes

import (
	"github.com/layasugar/laya"
	"github.com/layasugar/laya-template/controllers"
	"github.com/layasugar/laya-template/controllers/file"
	"net/http"
)

func Register(r *laya.WebServer) {
	r.GET("/", controllers.Ctrl.Version)   // version
	r.POST("/test", controllers.Ctrl.Test) // 测试接口

	// 文件服务器
	r.POST("/app/files/upload", file.Ctrl.Upload)
	r.StaticFS("/app/files", http.Dir("files"))

	RegisterApp(r)
	RegisterAdmin(r)
	RegisterTest(r)
}
