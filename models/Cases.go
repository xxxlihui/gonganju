package models
import (
	"time"
	"gz.com/master/random"
)

/*
案件信息
*/
type Cases struct {
	CasesId string `orm:"pk;size(32)"`
	CasesName string `orm:"size(200)"`
	CasesCode string `orm:"size(32)"`//案件编号
	CasesInfo string `orm:"type(text)"`
	CaseAllInfo string `orm:"type(text)"`//用于全文索引,把案件得全部内容豆放入去
	CasesTime time.Time
	CasesAddTime time.Time `orm:"auto_now_add"`
	CasesGrade float64 `orm:"type(numeric(3, 2))"`//案件得分数
	CasesStatus int //案件状态
	CasesStartTime time.Time
	CasesEndTime time.Time
	CasesEnd bool
}
func (this *Cases)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Cases)Insert()(int64,error){
	if this.CasesId==""{
		this.CasesId=random.NewId()
	}
	return o.Insert(this)
}
func (this *Cases)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Cases)Delete()(int64,error){
	return o.Delete(this)
}
