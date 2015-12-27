package controllers
import (
	"github.com/astaxie/beego"
	"gonganju/models"
)
type Login struct {
	beego.Controller
}
func (c *Login)Post(){
	acc:=c.GetString("acc")
	pwd:=c.GetString("pwd")
	police,rsp:=models.GetPoliceByAccPwd(acc,pwd)
	if police!=nil{
		c.SetSession("policeId",police.PoliceId)
	}
	c.Data["json"]=rsp
	c.ServeJson()
}