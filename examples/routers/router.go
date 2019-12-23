package routers

import (
	"github.com/freeFishGo/examples/fishgo"
	mvc "github.com/freeFishGo/middlewares/mvc/router"
)

func init() {
	// 注册主路由ControllerActionFuncName字段不用设置        设置了也不会生效
	fishgo.Mvc.AddMainRouter(&mvc.ControllerActionInfo{RouterPattern: "/{ Controller}/{Action}"})
}
