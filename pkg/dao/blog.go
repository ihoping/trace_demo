package dao

import "gorm.io/gorm"

type Article struct {
	ID    string `gorm:"id"`
	Title string `gorm:"title"`
}

func (Article) TableName() string {
	return "t_article"
}

func (dao *Dao) GetArticleDetail(id string) (*Article, error) {
	article := &Article{}
	err := dao.db.Where("id = ?", id).First(article).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return article, nil
}
