package main

import (
	"github.com/astaxie/beego"
	_ "github.com/wangle201210/rebxs/routers"
	"github.com/wangle201210/rebxs/models"

)

func main() {
	beego.Run()
}

func init() {
	models.Init()
}
