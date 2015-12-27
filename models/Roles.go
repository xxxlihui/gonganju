package models
import (
	"gz.com/master/random"
	"time"
)

/*
角色信息
*/
type Roles struct {
	RolesId string `orm:"pk;size(32)"`
	RolesName string `orm:"size(32)"`
	RolesAddTime time.Time	`orm:"auto_now_add"`
	RolesPermitCode int64
}
func (this *Roles)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Roles)Insert()(int64,error){
	if this.RolesId==""{
		this.RolesId=random.NewId()
	}
	return o.Insert(this)
}
func (this *Roles)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Roles)Delete()(int64,error){
	return o.Delete(this)
}
