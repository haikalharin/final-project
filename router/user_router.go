package router

import (
	"final-project/controller"
	"final-project/middleware"
	"final-project/repository"
	"final-project/service"

	"github.com/gin-gonic/gin"
	"github.com/masred/scalable-web-service-with-golang/session-12/final-project/app"
)

func UserRouter(router *gin.Engine) {
	db := app.NewDB()

	repo := repository.NewUserRepository(db)
	srv := service.NewUserService(repo)
	ctrl := controller.NewUserController(srv)

	userRouter := router.Group("/users")

	{
		userRouter.POST("/register", ctrl.Register)
		userRouter.POST("/login", ctrl.Login)

		{
			userRouter.PUT("/", middleware.AuthMiddleware(), ctrl.Update)
			userRouter.DELETE("/", middleware.AuthMiddleware(), ctrl.Delete)
		}
	}
}
