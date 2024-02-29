package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/ginGormCrudAPI/controller"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	// Code for the router
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	baseRouter := router.Group("/api")

	tagRouter := baseRouter.Group("/tags")

	tagRouter.GET("", tagsController.GetAllTags)
	tagRouter.GET("/:tagId", tagsController.GetById)
	tagRouter.POST("", tagsController.CreateTag)
	tagRouter.PUT("/:tagId", tagsController.UpdateTag)
	tagRouter.DELETE("/:tagId", tagsController.DeleteTag)

	return router

}
