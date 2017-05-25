package main

import (
	"modular-iris/modules/hello"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// Adapt the "httprouter", faster,
	// but it has limits on named path parameters' validation,
	// you can adapt "gorillamux" if you need regexp path validation!
	app.Adapt(httprouter.New())

	app.Get("/", func(ctx *iris.Context) {
		ctx.Writef("hello world\n")
	})
	hello.Install(app)
	app.Listen(":8080")
}
