package router

import (
	"go-cleanarch/internal/controller"
	"go-cleanarch/internal/repository"
	"go-cleanarch/internal/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func addTodoGroup(router *gin.Engine) {
	todoRepo, err := repository.NewPostgresTodoRepository()
	if err != nil {
		slog.Error("[Router] NewPostgresTodoRepository", "err", err)
		panic(err)
	}

	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("", todoController.GetAll)
		todoGroup.POST("", todoController.PostOne)
		todoGroup.GET("/:id", todoController.GetOne)
		todoGroup.PATCH("/:id", todoController.UpdateOne)
		todoGroup.DELETE("/:id", todoController.DeleteOne)
	}
}
