package app

import (
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// InitIris 初始化Iris相关配置
func InitIris() {
	// 创建一个 iris 的 Application，负责管理应用的状态，提供高效的web处理能力
	app := iris.New()
	// 设置应用的日志级别：开发时为debug,生产应视情况调整为info或warn级别
	app.Logger().SetLevel("debug")
	// 使用 recover 中间件，它保证应用从各种panic恢复过来并记录相关日志,日志显示级别:warn
	app.Use(recover.New())
	// 使用 logger 中间件，不是与框架的 Logger 混淆，该中间件仅记录 http 请求的日志
	app.Use(logger.New())
	// 对于指定的 http 请求方法，允许重新注册到 iris 的路由组中，如果已注册过则忽略
	app.AllowMethods(iris.MethodOptions)

	// 对于任意的 http 协议的"/"请求,都用同一个处理函数
	app.Any("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to v-mall-go!</h1>")
	})
	// start record.
	// app.Use(func(ctx iris.Context) {
	// 	ctx.Record()
	// 	ctx.Next()
	// })

	// // collect and "log".
	// app.Done(func(ctx iris.Context) {
	// 	body := ctx.Recorder().Body()
	// 	// Should print success.
	// 	app.Logger().Infof("sent: %s", string(body))
	// })

	// It applies per Party and its children,
	// therefore, you can create a routes := app.Party("/path")
	// and set middlewares, their rules and the routes there as well.
	// app.SetExecutionRules(iris.ExecutionRules{
	// 	Done: iris.ExecutionOptions{Force: true},
	// })

	app.RegisterView(iris.HTML("./web/views", ".html"))

	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.ViewData("message", "Hello World!")
	// 	ctx.View("index.html")
	// })

	// app.Get("/save", func(ctx iris.Context) {
	// 	ctx.WriteString("success")
	// })

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"messgae": "pong", "time": time.Now()})
	})
	// app.Get("/user/id/{id:uint64}", func(ctx iris.Context) {
	// 	userID, _ := ctx.Params().GetUint64("id")
	// 	ctx.Writef("User Id %d", userID)
	// })
	// app.Post("path", func(ctx iris.Context) {
	// 	ctx.WriteString("Post response")
	// })

	// app.Get("/exit", func(ctx iris.Context) {
	// 	ctx.HTML(`
	// 		<h1>Hi</h1>
	// 		<h2>测试退出功能</h2>
	// 	`)
	// })

	// 程序退出之后执行相应功能
	// app.ConfigureHost(func(s *iris.Supervisor) {
	// 	s.RegisterOnShutdown(func() {
	// 		app.Logger().Info("server terminated!")
	// 	})
	// })

	// 优雅退出，注册中断响应函数
	// iris.RegisterOnInterrupt(func() {
	// 	timeout := time.Minute * 2
	// 	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// 	defer cancel()
	// 	app.Logger().Info("过2分钟再执行Shutdown函数")
	// 	app.Shutdown(ctx)
	// })
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"), iris.WithTimeFormat("2006-01-02 13:04:05"), iris.WithoutInterruptHandler)

	// app.NewHost(&http.Server{Addr: "9090"}).ListenAndServe()

}

func myLog(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
