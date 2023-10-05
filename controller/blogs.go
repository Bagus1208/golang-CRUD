package controller

import (
	"fmt"
	"go_bagus-ario-yudanto/18_Middleware/praktikum/helper"
	"go_bagus-ario-yudanto/18_Middleware/praktikum/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	model model.BlogModel
}

func (blogController *BlogController) InitBlogController(blogModel model.BlogModel) {
	blogController.model = blogModel
}

func (blogController *BlogController) GetBlogs(c echo.Context) error {
	var result = blogController.model.GetAllBlogs()

	return c.JSON(http.StatusOK, helper.SetResponse("success get all blogs", result))
}

func (blogController *BlogController) GetBlog(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	var result = blogController.model.GetBlogById(id)

	return c.JSON(http.StatusOK, helper.SetResponse("success get blog", result))
}

func (blogController *BlogController) Create(c echo.Context) error {
	var blog model.Blog
	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	var result = blogController.model.CreateBlog(blog)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, helper.SetResponse("cannot process data, something happend", nil))
	}

	return c.JSON(http.StatusCreated, helper.SetResponse("success create blog", result))
}

func (blogController *BlogController) Update(c echo.Context) error {
	var update model.Blog
	if err := c.Bind(&update); err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(fmt.Sprint("error when parshing data -", err.Error()), nil))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	update.ID = uint(id)
	var result = blogController.model.UpdateBlog(update)

	return c.JSON(http.StatusOK, helper.SetResponse("success update", result))
}

func (blogController *BlogController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.SetResponse(err.Error(), nil))
	}

	blogController.model.DeleteBlog(id)

	return c.JSON(http.StatusOK, []any{})
}
