package controller

import (
	"fmt"
	"go_bagus-ario-yudanto/18_Middleware/praktikum/helper"
	"go_bagus-ario-yudanto/18_Middleware/praktikum/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	model model.BookModel
}

func (bookController *BookController) InitBookController(bookModel model.BookModel) {
	bookController.model = bookModel
}

func (bookController *BookController) GetBooks(c echo.Context) error {
	var result = bookController.model.GetAllBooks()

	return c.JSON(http.StatusOK, helper.SetResponse("success get all books", result))
}

func (bookController *BookController) GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	var result = bookController.model.GetBookById(id)

	return c.JSON(http.StatusOK, helper.SetResponse("success get book", result))
}

func (bookController *BookController) Insert(c echo.Context) error {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	var result = bookController.model.InsertBook(book)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
	}

	return c.JSON(http.StatusCreated, helper.SetResponse("success insert book", result))
}

func (bookController *BookController) Update(c echo.Context) error {
	var update model.Book
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	update.ID = uint(id)
	var result = bookController.model.UpdateBook(update)

	return c.JSON(http.StatusOK, helper.SetResponse("success update", result))
}

func (bookController *BookController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	bookController.model.DeleteBook(id)

	return c.JSON(http.StatusOK, []any{})
}
