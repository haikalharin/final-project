package router

import (
	"final-project/controller"
	"final-project/middleware"
	"final-project/repository"
	"final-project/service"

	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/app"
)

func PhotoRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewPhotoRepository(db)
	srv := service.NewPhotoService(repo)
	ctrl := controller.NewPhotoController(srv)

	photoRouter := router.Group("/photos", middleware.AuthMiddleware())

	{
		photoRouter.POST("/", ctrl.Create)
		photoRouter.GET("/", ctrl.GetAll)
		{
			photoRouter.PUT("/:id", middleware.PhotoMiddleware(srv), ctrl.Update)
			photoRouter.DELETE("/:id", middleware.PhotoMiddleware(srv), ctrl.Delete)
		}
	}
}
