package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"vincentmegia.com/todoweb/pkg/web/components"
)

func main() {
	//
	fmt.Println("hello todo app-------------")
	/* 	checkbox := components.TodoItem("1", "This is a simple description", "this is a simple value") */
	html := components.Html()
	server := echo.New()

	// assign paths
	// default package
	server.GET("/", func(e echo.Context) error {
		return html.Render(e.Request().Context(), e.Response().Writer)
	})
	server.Logger.Fatal(server.Start(":3000"))
	/* server.POST("/todo", func(e echo.Context) error {
		fmt.Println("creating a todo item in db")
		return checkbox.Render(e.Request().Context(), e.Response().Writer)
	}) */

}
