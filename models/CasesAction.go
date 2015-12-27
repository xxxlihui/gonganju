package models
import (
	"gz.com/master/random"
	"time"
)


/*

*/
/*
案件执行信息
*/
type CasesAction struct {
	CacId string `orm:"pk;size(32)"`
	CacInfo string `orm:"type(text)"`
	CacTime time.Time
	CacAddTime time.Time
}
func (this *CasesAction)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *CasesAction)Insert()(int64,error){
	if this.CacId==""{
		this.CacId=random.NewId()
	}
	return o.Insert(this)
}
func (this *CasesAction)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *CasesAction)Delete()(int64,error){
	return o.Delete(this)
}
