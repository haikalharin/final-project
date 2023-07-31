package router

import (
	"final-project/controller"
	"final-project/middleware"
	"final-project/repository"
	"final-project/service"

	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/app"
)

func CommentRouter(router *gin.Engine) {
	db := app.NewDB()

	repoPhoto := repository.NewPhotoRepository(db)
	srvPhoto := service.PhotoService(repoPhoto)

	repoComment := repository.NewCommentRepository(db)
	srvComment := service.NewCommentService(repoComment)

	ctrl := controller.NewCommentController(srvComment, srvPhoto)

	commentRouter := router.Group("/comments", middleware.AuthMiddleware())

	{
		commentRouter.POST("/", ctrl.Create)
		commentRouter.GET("/", ctrl.GetAll)
		{
			commentRouter.PUT("/:commentId", middleware.CommentMiddleware(srvComment), ctrl.Update)
			commentRouter.DELETE("/:commentId", middleware.CommentMiddleware(srvComment), ctrl.Delete)
		}
	}
}
