package models
import (
	"time"
	"gz.com/master/random"
)
//案件走的流程记录
type Approve struct {
	ApproveId string `orm:"pk;size(32)"`
	ApproveCaseId string `orm:"size(32)"`
	ApproveStatue int
	ApproveInfo string `orm:size(4000)`
	ApproveAddTime time.Time
}
func (this *Approve)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Approve)Insert()(int64,error){
	if this.ApproveId==""{
		this.ApproveId=random.NewId()
	}
	return o.Insert(this)
}
func (this *Approve)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Approve)Delete()(int64,error){
	return o.Delete(this)
}
