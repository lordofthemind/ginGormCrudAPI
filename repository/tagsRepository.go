package repository

import "github.com/lordofthemind/ginGormCrudAPI/model"

type TagsRepository interface {
	SaveTag(tags model.TagsModel)
	UpdateTag(tags model.TagsModel)
	DeleteTag(tags model.TagsModel)
	GetAllTags() []model.TagsModel
	GetTagByID(id int) (tags model.TagsModel, err error)
}
