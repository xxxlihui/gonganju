package models
import "gz.com/master/random"

/*
角色对应的权限
*/
type RolesPermit struct {
	RolesPermitId string `orm:"pk;size(32)"`
	RolesId string `orm:"size(32)"`
	PermitId string `orm:"size(32)"`
}
func (this *RolesPermit)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *RolesPermit)Insert()(int64,error){
	if this.RolesPermitId==""{
		this.RolesPermitId=random.NewId()
	}
	return o.Insert(this)
}
func (this *RolesPermit)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *RolesPermit)Delete()(int64,error){
	return o.Delete(this)
}
