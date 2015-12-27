package models
import "gz.com/master/random"

/*
部门信息
*/
type Dept struct {
	DeptId string `orm:"pk;size(32)"`
	ParentId string `orm:"size(32)"`
	DeptName string  `orm:"size(32)"`
	DeptAddress string  `orm:"size(200)"`
	DeptTelephoneNumber string  `orm:"size(32)"`
}
func (this *Dept)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Dept)Insert()(int64,error){
	if this.DeptId==""{
		this.DeptId=random.NewId()
	}
	return o.Insert(this)
}
func (this *Dept)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Dept)Delete()(int64,error){
	return o.Delete(this)
}
