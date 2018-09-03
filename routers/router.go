package routers

import (
	"github.com/astaxie/beego"
	"github.com/datdao0712/bunteam-colorize-ai/controllers"
)

func init() {
	beego.SetStaticPath("/frontend", "frontend/bunteam-colorize-ai/dist")
	beego.Router("/", &controllers.MainController{})
}
