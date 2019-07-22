package main

import (
	"github.com/astaxie/beego"
	"github.com/wangle201210/rebxs/models"
	_ "github.com/wangle201210/rebxs/routers"
)

func main() {
	beego.Run()
}

func init() {
	models.Init()
	beego.SetStaticPath("/static","static")
}
