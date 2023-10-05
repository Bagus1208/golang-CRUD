package main

import (
	"CRUD_bagus/configs"
	"CRUD_bagus/controller"
	"CRUD_bagus/helper"
	"CRUD_bagus/model"
	"CRUD_bagus/routes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	var e = echo.New()
	var config = configs.InitConfig()

	var DB = model.InitModel(*config)
	model.Migrate(DB)

	userModel := model.UserModel{}
	userModel.Init(DB)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel, *config)

	bookModel := model.BookModel{}
	bookModel.Init(DB)

	bookControll := controller.BookController{}
	bookControll.InitBookController(bookModel)

	blogModel := model.BlogModel{}
	blogModel.Init(DB)

	blogControll := controller.BlogController{}
	blogControll.InitBlogController(blogModel)

	helper.LogMiddlewares(e)

	routes.RouteUser(e, userControll, *config)
	routes.RouteBook(e, bookControll, *config)
	routes.RouteBlog(e, blogControll)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Server_Port)).Error())
}
