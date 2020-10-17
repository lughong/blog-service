package dao

import "github.com/lughong/blog-service/internal/model"

func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByAID(d.engine)
}

func (d *Dao) getArticleTagListByTID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagID: tagID}
	return articleTag.ListByTID(d.engine)
}

func (d *Dao) GetArticleTagListByAIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIDs(d.engine, articleIDs)
}

func (d *Dao) CreateArticleTag(articleID, tagID uint32, CreatedBy string) error {
	articleTag := model.ArticleTag{
		Model: &model.Model{
			CreatedBy: CreatedBy,
		},
		ArticleID: articleID,
		TagID: tagID,
	}

	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleID, tagID uint32, ModifiedBy string) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	values := map[string]interface{}{
		"article_id": articleID,
		"tag_id": tagID,
		"modified_by": ModifiedBy,
	}

	return articleTag.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.DeleteOne(d.engine)
}