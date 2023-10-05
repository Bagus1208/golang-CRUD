package routes

import (
	"go_bagus-ario-yudanto/18_Middleware/praktikum/configs"
	"go_bagus-ario-yudanto/18_Middleware/praktikum/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, usercontroller controller.UserController, config configs.Config) {
	e.POST("/users", usercontroller.Create)
	e.POST("/users/login", usercontroller.Login)

	users := e.Group("/users")
	users.Use(echojwt.JWT([]byte(config.Secret)))
	users.GET("", usercontroller.GetUsers)
	users.GET("/:id", usercontroller.GetUser)
	users.GET("/blogs", usercontroller.GetBlogs)
	users.PUT("/:id", usercontroller.Update)
	users.DELETE("/:id", usercontroller.Delete)
}

func RouteBook(e *echo.Echo, bookController controller.BookController, config configs.Config) {
	books := e.Group("/books")
	books.Use(echojwt.JWT([]byte(config.Secret)))
	books.GET("", bookController.GetBooks)
	books.GET("/:id", bookController.GetBook)
	books.POST("", bookController.Insert)
	books.PUT("/:id", bookController.Update)
	books.DELETE("/:id", bookController.Delete)
}

func RouteBlog(e *echo.Echo, blogController controller.BlogController) {
	blogs := e.Group("/blogs")
	blogs.GET("", blogController.GetBlogs)
	blogs.GET("/:id", blogController.GetBlog)
	blogs.POST("", blogController.Create)
	blogs.PUT("/:id", blogController.Update)
	blogs.DELETE("/:id", blogController.Delete)
}
