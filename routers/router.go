package routers

import (
	"github.com/wangle201210/rebxs/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
