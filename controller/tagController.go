package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/ginGormCrudAPI/data/request"
	"github.com/lordofthemind/ginGormCrudAPI/data/response"
	"github.com/lordofthemind/ginGormCrudAPI/helper"
	"github.com/lordofthemind/ginGormCrudAPI/service"
)

type TagsController struct {
	TagService service.TagsService
}

func NewTagsController(tagService service.TagsService) *TagsController {
	return &TagsController{
		TagService: tagService,
	}
}

// CreateTag is a controller method to create a new tag.
func (controller *TagsController) CreateTag(ctx *gin.Context) {
	// implementation for creating a tag
	createTagRequest := request.CreateTagRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.TagService.Create(createTagRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTag is a controller method to update a new tag.
func (controller *TagsController) UpdateTag(ctx *gin.Context) {
	// implementation for updating a tag
	updateTagRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagRequest.Id = id

	controller.TagService.Update(updateTagRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTag is a controller method to delete a new tag.
func (controller *TagsController) DeleteTag(ctx *gin.Context) {
	// implementation for deleting a tag
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.TagService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(http.StatusOK, webResponse)
}

// CreateTag is a controller method to create a new tag.
func (controller *TagsController) GetById(ctx *gin.Context) {
	// implementation for creating a tag
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.TagService.GetById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(http.StatusOK, webResponse)
}

// GetAllTags is a controller method to get all tags.
func (controller *TagsController) GetAllTags(ctx *gin.Context) {
	// implementation for getting all tags
	tagsResponse := controller.TagService.GetAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   tagsResponse,
	}
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(http.StatusOK, webResponse)
}
