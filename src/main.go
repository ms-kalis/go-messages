package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	app.Run(iris.Addr("8060"))
}
func notFoundHandler(ctx iris.Context) {
	ctx.HTML("Custom route for 404 not found http code, here you can render a view, html, json <b>any valid response</b>.")
}
