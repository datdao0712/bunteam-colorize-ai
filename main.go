package main

import (
	"github.com/astaxie/beego"
	_ "github.com/datdao0712/bunteam-colorize-ai/routers"
)

func main() {
	beego.BConfig.WebConfig.ViewsPath = "frontend/bunteam-colorize-ai/dist/"
	beego.Run()
}
