package service

import (
	"github.com/lordofthemind/ginGormCrudAPI/data/request"
	"github.com/lordofthemind/ginGormCrudAPI/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	GetById(tagsId int) response.TagsResponse
	GetAll() []response.TagsResponse
}
