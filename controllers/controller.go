package controllers
import (
	"github.com/astaxie/beego"
	"gonganju/models"
	"gonganju/consts/errs"
)
type Controller struct {
	beego.Controller
	Police *models.Police
	QueryType int
}
func (c *Controller)Prepare(){
	c.TplExt="html"
	policeId,ok:= c.GetSession("policeId").(string)
	if !ok||policeId==""{
		c.Redirect("/",302)
		c.StopRun()
	}
	police,rsp:=models.GetPoliceById(policeId)
	if rsp.ResultID!=0{
		rsp.ResultID=errs.ErrNoUser
		rsp.ResultMsg="用户已退出登陆"
		switch {
		case c.Ctx.Input.AcceptsJson():
			c.Data["json"]=rsp
			c.ServeJsonp()
			c.StopRun()
			return
		case c.Ctx.Input.AcceptsXml():
			c.Data["xml"]=rsp
			c.ServeXml()
			c.StopRun()
			return
		case c.Ctx.Input.AcceptsHtml():
			fallthrough
		default:
			c.Redirect("/",302)
			c.StopRun()
		}
	}
	c.Police=police
}