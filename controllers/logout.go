package controllers
import "github.com/astaxie/beego"
type Logout struct {
	beego.Controller
}
func (c *Logout)Get(){
	c.DelSession("policeId")
	c.Redirect("/",302)
}