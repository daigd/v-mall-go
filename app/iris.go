package app

import (
	"github.com/kataras/iris"
)

// InitIris 初始化Iris相关配置
func InitIris() {
	app := iris.New()
	app.Use(myLog)
	// app.Use(logger.New())
	app.RegisterView(iris.HTML("./web/views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello World!")
		ctx.View("index.html")
	})
	app.Get("ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"messgae": "pong"})
	})
	app.Get("hi", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"messgae": "hi"})
	})

	app.Run(iris.Addr(":8088"), iris.WithCharset("UTF-8"))

}

func myLog(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
