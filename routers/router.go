package routers

import (
	"github.com/astaxie/beego"
	"github.com/datdao0712/bunteam-colorize-ai/controllers"
)

func init() {
	beego.Router("/*", &controllers.MainController{})
}
