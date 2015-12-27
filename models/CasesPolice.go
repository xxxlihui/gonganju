package models
import (
	"gz.com/master/random"
	"time"
)

/*
案件安排的人员
*/
type CasesPolice struct {
	CsolId        string `orm:"pk;size(32)"`
	CpsCaseId     string `orm:"size(32)"`
	CpsPoliceId   string `orm:"size(32)"`
	CpsRole       string `orm:"size(32)"`
	CpsGrade      float64 `orm:"type(numeric(3, 2))"` //案件中每个警官应该获得的分数
	CpsDispatchId string `orm:"size(32)"`
	CpsAddTime    time.Time `orm:"auto_now_add"`
}
func (this *CasesPolice)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *CasesPolice)Insert()(int64,error){
	if this.CpsCaseId ==""{
		this.CpsCaseId =random.NewId()
	}
	return o.Insert(this)
}
func (this *CasesPolice)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *CasesPolice)Delete()(int64,error){
	return o.Delete(this)
}
