package models

/*
权限说明
*/
type Permit struct {
	PermitId int64 `orm:"pk;auto"`
	PermitCode int64
	PermitInfo string `orm:"size(200)"`
}
func (this *Permit)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Permit)Insert() (int64,error){
	return o.Insert(this)
}
func (this *Permit)Update(field...string) (int64,error){
	return o.Update(this,field...)
}
func (this *Permit)Delete() (int64,error){
	return o.Delete(this)
}

