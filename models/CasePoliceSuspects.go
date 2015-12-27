package models
import "gz.com/master/random"


/*
案件中警官抓捕嫌疑人的个人情况
*/
type CasePoliceSuspects struct {
	CpsId string `orm:"pk;size(32)"`
	CpsDeptId string `orm:"size(32)"`
	CpsCaseId string `orm:"size(32)"`
	CpsPoliceId string `orm:"size(32)"`
	CpsSuspectsId string `orm:"size(32)"`
	CpsInfo string `orm:"type(text)"`
}
func (this *CasePoliceSuspects)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *CasePoliceSuspects)Insert()(int64,error){
	if this.CpsCaseId==""{
		this.CpsCaseId=random.NewId()
	}
	return o.Insert(this)
}
func (this *CasePoliceSuspects)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *CasePoliceSuspects)Delete()(int64,error){
	return o.Delete(this)
}

