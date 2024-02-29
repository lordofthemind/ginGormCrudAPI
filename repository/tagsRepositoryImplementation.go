// repository/tagsRepositoryImplementation.go

package repository

import (
	"github.com/lordofthemind/ginGormCrudAPI/data/request"
	"github.com/lordofthemind/ginGormCrudAPI/helper"
	"github.com/lordofthemind/ginGormCrudAPI/model"
	"gorm.io/gorm"
)

// TagsRepositoryImplementation is the implementation of TagsRepository.
type TagsRepositoryImplementation struct {
	DB *gorm.DB
}

func NewTagsRepositoryImplementation(db *gorm.DB) *TagsRepositoryImplementation {
	return &TagsRepositoryImplementation{DB: db}
}

func (t *TagsRepositoryImplementation) SaveTag(tags model.TagsModel) {
	// implementation for saving a tag
	result := t.DB.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImplementation) UpdateTag(tags model.TagsModel) {
	// implementation for updating a tag
	var updateTag = request.UpdateTagsRequest{
		Id:   tags.ID,
		Name: tags.Name,
	}
	result := t.DB.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImplementation) DeleteTag(tags model.TagsModel) {
	// implementation for deleting a tag
	var tag model.TagsModel
	result := t.DB.Where("id = ?", tags.ID).Delete(&tag)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImplementation) GetAllTags() []model.TagsModel {
	// implementation for getting all tags
	var tags []model.TagsModel
	result := t.DB.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

func (t *TagsRepositoryImplementation) GetTagByID(tagId int) (tags model.TagsModel, err error) {
	// implementation for getting a tag by ID
	var tag model.TagsModel
	result := t.DB.Find(&tag, tagId)
	if result.Error != nil {
		return model.TagsModel{}, result.Error
	}

	return tag, nil
}
