package blog_dao

type Blog struct {
	ID    string `gorm:"id"`
	Title string `gorm:"title"`
}

func (Blog) TableName() string {
	return "t_blog"
}
