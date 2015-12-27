package models
import (
	"gz.com/master/random"
	"time"
)

/*
嫌疑人
*/
type Suspects struct {
	SuspectsId string `orm:"pk;size(32)"`
	SuspectsName string `orm:"size(100)"`
	SuspectsIdCardNumber string `orm:"size(100)"`
	SuspectsNative string `orm:"size(100)"`
	SuspectsAddress string `orm:"size(100)"`
	SuspectsPhotoId string `orm:"size(100)"`
	SuspectsCode string `orm:"size(32)"` //嫌疑人编号
	AddTime time.Time `orm:"auto_now_add"`
}
func (this *Suspects)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Suspects)Insert()(int64,error){
	if this.SuspectsId==""{
		this.SuspectsId=random.NewId()
	}
	return o.Insert(this)
}
func (this *Suspects)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Suspects)Delete()(int64,error){
	return o.Delete(this)
}
