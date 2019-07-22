package routers

import (
	"github.com/astaxie/beego"
	"github.com/wangle201210/rebxs/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.UserController{}, "get:All")
	beego.Router("/projects", &controllers.ProjectController{}, "get:All")
	beego.Router("/record", &controllers.RecordController{}, "post:Save")
}
