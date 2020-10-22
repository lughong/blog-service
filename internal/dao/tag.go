package dao

import (
	"github.com/lughong/blog-service/internal/model"
	"github.com/lughong/blog-service/pkg/app"
)

func (d *Dao) GetTag(tagID uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{
		State: state,
		Model: &model.Model{
			ID: tagID,
		},
	}

	return tag.Get(d.engine)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createBy,
		},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	values := map[string]interface{}{
		"Name":       name,
		"State":      state,
		"ModifiedBy": modifiedBy,
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}

	return tag.Delete(d.engine)
}
