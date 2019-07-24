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
	beego.Router("/recordList", &controllers.RecordController{}, "get:GetList")
	beego.Router("/showOne", &controllers.RecordController{}, "get:ShowOne")
	beego.Router("/reward", &controllers.RewardController{}, "post:Save")
	beego.Router("/rewardList", &controllers.RewardController{}, "get:ShowList")

}
