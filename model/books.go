package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penulis  string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}

type BookModel struct {
	DB *gorm.DB
}

func (bookModel *BookModel) Init(DB *gorm.DB) {
	bookModel.DB = DB
}

func (bookModel *BookModel) GetAllBooks() []Book {
	var listBook []Book
	if err := bookModel.DB.Find(&listBook).Error; err != nil {
		logrus.Error("Model : Get all user data error, ", err.Error())
		return nil
	}

	return listBook
}

func (bookModel *BookModel) GetBookById(id int) Book {
	var book Book
	book.ID = uint(id)
	if err := bookModel.DB.First(&book).Error; err != nil {
		logrus.Error("Model : Get user data error, ", err.Error())
		return Book{}
	}

	return book
}

func (bookModel *BookModel) InsertBook(newBook Book) *Book {
	if err := bookModel.DB.Create(&newBook).Error; err != nil {
		logrus.Error("Model : Create user data error, ", err.Error())
		return nil
	}

	return &newBook
}

func (bookModel *BookModel) UpdateBook(updateBook Book) *Book {
	var query = bookModel.DB.Updates(updateBook)
	if err := query.Error; err != nil {
		logrus.Error("Model : Update book data error, ", err.Error())
		return nil
	}

	if dataCount := query.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data affected")
		return &Book{}
	}

	return &updateBook
}

func (bookModel *BookModel) DeleteBook(id int) {
	var deleteBook Book
	deleteBook.ID = uint(id)
	if err := bookModel.DB.Delete(&deleteBook).Error; err != nil {
		logrus.Error("Model : Delete user data error, ", err.Error())
	}
}
