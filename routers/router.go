package routers

import (
	"github.com/astaxie/beego"
	"temp-admin/config"
	"temp-admin/controllers"
	"temp-admin/logs"
	"temp-admin/models"
)

func init() {
	startUp()
	nsUser := beego.NewNamespace("/api/back/",
		beego.NSRouter("login", &controllers.AccountController{}, "post:Login"),
		beego.NSRouter("register", &controllers.AccountController{}, "put:Register"),
		beego.NSRouter("store", &controllers.StoreController{}, "post:Update"),
		beego.NSRouter("store", &controllers.StoreController{}, "get:Lists"),
		beego.NSRouter("store/info", &controllers.StoreController{}, "get:Info"),

		beego.NSRouter("device", &controllers.DeviceControllers{}, "get:Lists"),
		beego.NSRouter("device", &controllers.DeviceControllers{}, "post:Bind"),
		beego.NSRouter("device/:sn([A-Z0-9]+)", &controllers.DeviceControllers{}, "delete:UnBind"),

		beego.NSRouter("user", &controllers.UserControllers{}, "get:Lists"),
		beego.NSRouter("user", &controllers.UserControllers{}, "post:Update"),
		beego.NSRouter("user/:id([0-9]+)", &controllers.UserControllers{}, "delete:Del"),

		beego.NSRouter("temp", &controllers.TempControllers{}, "get:Lists"),
		//beego.NSRouter("logout", &controllers.BaseController{}, "post:LogOut"),
	)
	beego.AddNamespace(nsUser)
}

//项目所有相关的初始化
func startUp() {
	beego.BConfig.CopyRequestBody = true
	//beego.BConfig.AppName = "temp-admin"
	//日志 初始化
	if err := config.Init(); err != nil {
		panic(err)
	}
	if err := logs.InitLogs(); err != nil {
		panic(err)
	}
	if err := models.Init(); err != nil {
		panic(err)
	}
}
