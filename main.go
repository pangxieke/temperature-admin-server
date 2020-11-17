package main

import (
	"github.com/astaxie/beego"
	_ "temp-admin/routers"
)

func main() {
	beego.Run(":19980")
}
