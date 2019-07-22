package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/wangle201210/rebxs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Get("/test",func(ctx *context.Context){
		ctx.Output.Body([]byte("hello world"))
	})
}
