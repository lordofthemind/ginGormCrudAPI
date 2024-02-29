// service/tagServiceImplementation.go

package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/lordofthemind/ginGormCrudAPI/data/request"
	"github.com/lordofthemind/ginGormCrudAPI/data/response"
	"github.com/lordofthemind/ginGormCrudAPI/helper"
	"github.com/lordofthemind/ginGormCrudAPI/model"
	"github.com/lordofthemind/ginGormCrudAPI/repository"
)

// TagServiceImplementation is the implementation of TagsService.
type TagServiceImplementation struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

// NewTagServiceImplementation creates a new instance of TagServiceImplementation.
func NewTagServiceImplementation(tagRepository repository.TagsRepository, validate *validator.Validate) *TagServiceImplementation {
	return &TagServiceImplementation{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

func (t *TagServiceImplementation) Create(tags request.CreateTagRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	TagsModel := model.TagsModel{
		Name: tags.Name,
	}
	t.TagsRepository.SaveTag(TagsModel)
}

func (t *TagServiceImplementation) Delete(tagsId int) {
	// Construct a TagsModel with the provided ID
	tagToDelete := model.TagsModel{ID: tagsId}

	// Delegates the deletion to the repository
	t.TagsRepository.DeleteTag(tagToDelete)
}

func (t *TagServiceImplementation) GetAll() []response.TagsResponse {
	result := t.TagsRepository.GetAllTags()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.ID,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}

// FindById implements TagsService
// func (t *TagServiceImplementation) GetById(tagsId int) response.TagsResponse {
// 	tagData, err := t.TagsRepository.GetTagByID(tagsId)

// 	tagResponse := response.TagsResponse{
// 		Id:   tagData.ID,
// 		Name: tagData.Name,
// 	}
// 	return tagResponse
// }

func (t *TagServiceImplementation) GetById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.GetTagByID(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.ID,
		Name: tagData.Name,
	}
	return tagResponse
}

func (t *TagServiceImplementation) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.GetTagByID(tags.Id)
	helper.ErrorPanic(err)

	tagData.Name = tags.Name
	t.TagsRepository.UpdateTag(tagData)
}
