package routers

import (
	"gonganju/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.Index{})
	beego.Router("/login", &controllers.Login{})
	beego.Router("/logout", &controllers.Logout{})
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/prettify", "static/prettify")
	userNs := beego.NewNamespace(
		"/UserNS",
		beego.NSAutoRouter(&controllers.Cases{}),
		beego.NSAutoRouter(&controllers.CasesAction{}),
		beego.NSAutoRouter(&controllers.Police{}),
		beego.NSAutoRouter(&controllers.PoliceInfo{}),
	)
	beego.AddNamespace(userNs)
}
