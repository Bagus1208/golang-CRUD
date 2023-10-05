package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	UserID uint   `json:"user_id" form:"user_id"`
	Judul  string `json:"judul" form:"judul"`
	Konten string `json:"konten" form:"konten"`
}

type BlogModel struct {
	DB *gorm.DB
}

func (blogModel *BlogModel) Init(DB *gorm.DB) {
	blogModel.DB = DB
}

func (blogModel *BlogModel) GetAllBlogs() []Blog {
	var listBlog []Blog
	if err := blogModel.DB.Find(&listBlog).Error; err != nil {
		logrus.Error("Model : Get all blog data error, ", err.Error())
		return nil
	}

	return listBlog
}

func (blogModel *BlogModel) GetBlogById(id int) Blog {
	var blog Blog
	blog.ID = uint(id)
	if err := blogModel.DB.First(&blog).Error; err != nil {
		logrus.Error("Model : Get Blog data error, ", err.Error())
		return Blog{}
	}

	return blog
}

func (blogModel *BlogModel) CreateBlog(newBlog Blog) *Blog {
	if err := blogModel.DB.Create(&newBlog).Error; err != nil {
		logrus.Error("Model : Create Blog data error, ", err.Error())
		return nil
	}

	return &newBlog
}

func (blogModel *BlogModel) UpdateBlog(updateBlog Blog) *Blog {
	var query = blogModel.DB.Updates(updateBlog)
	if err := query.Error; err != nil {
		logrus.Error("Model : Update Blog data error, ", err.Error())
		return nil
	}

	if dataCount := query.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data affected")
		return &Blog{}
	}

	return &updateBlog
}

func (blogModel *BlogModel) DeleteBlog(id int) {
	var deleteBlog Blog
	deleteBlog.ID = uint(id)
	if err := blogModel.DB.Delete(&deleteBlog).Error; err != nil {
		logrus.Error("Model : Delete Blog data error, ", err.Error())
	}
}
