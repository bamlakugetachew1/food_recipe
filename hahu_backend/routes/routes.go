package routes

import (
	"hahu_jobs_projects/controller"

	"github.com/gin-gonic/gin"
)

func UseRoute(router *gin.Engine) {
	publicRoutes := router.Group("/route")
	publicRoutes.POST("/register", controller.Registfromhasura)
	publicRoutes.POST("/login", controller.Logfromhasura)
	publicRoutes.POST("/file", controller.Uploadfile)
	publicRoutes.POST("/handleuser", controller.Handleuser)

}
