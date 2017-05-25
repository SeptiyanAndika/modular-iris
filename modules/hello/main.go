package hello

import "gopkg.in/kataras/iris.v6"

//Install i
func Install(app *iris.Framework) {

	logic := Logic{}

	app.Get("/hello", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, logic.helloFunction("John"))
	})

}
