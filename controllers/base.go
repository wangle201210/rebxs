package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

var (
	resp = Response{}
)

type BaseController struct {
	beego.Controller
}

// 预留个基础控制器
func init()  {
	beego.BConfig.WebConfig.TemplateLeft="<<<"
	beego.BConfig.WebConfig.TemplateRight=">>>"
	return
}

//Response 结构体
type Response struct {
	Code int64         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type status struct {
	code int64
	text string
}
var (
	addSuccess		= status{http.StatusCreated,"添加成功"}
	addError		= status{http.StatusNoContent,"添加失败"}
	readSuccess		= status{http.StatusOK,"查询成功"}
	readError		= status{http.StatusNoContent,"查询失败"}
	deleteSuccess	= status{http.StatusOK,"删除成功"}
	deleteError		= status{http.StatusNoContent,"查询失败"}
	updateSuccess	= status{http.StatusResetContent,"更新成功"}
	updateError		= status{http.StatusNoContent,"更新失败"}
	paraError       = status{http.StatusAccepted,"传入参数不对"}
)