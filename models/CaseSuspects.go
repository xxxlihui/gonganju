package models
import "gz.com/master/random"


/*
案件关联的嫌疑人,嫌疑人在这个案件中的信息
 */
type CaseSuspects struct {
	CsId string `orm:"pk;size(32)"`
	CsCaseId string `orm:"size(32)"`
	CsSuspectsId string `orm:"size(32)"`
	CsInfo string `orm:"type(text)"`
}
func (this *CaseSuspects)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *CaseSuspects)Insert()(int64,error){
	if this.CsCaseId==""{
		this.CsCaseId=random.NewId()
	}
	return o.Insert(this)
}
func (this *CaseSuspects)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *CaseSuspects)Delete()(int64,error){
	return o.Delete(this)
}
