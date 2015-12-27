package models
import (
	"time"
	"nong/random"
)
/*
警官信息
*/
type Police struct {
	PoliceId string  `orm:"pk;size(32)"`
	PoliceName string `orm:"size(32)"` //
	PoliceAcc string  `orm:"size(32)"`
	PolicePwd string `orm:"size(32)"`
	PoliceIcon string `orm:"size(100)"` //头像
	PoliceTelephoneNumber string `orm:"size(32)"`
	PoliceDeptId string `orm:"size(32)"`
	PoliceCode string `orm:"size(32)"` //警号
	PoliceAddTime time.Time `orm:"auto_now_add"`
	/*简单化权限控制*/
	PoliceRole string `orm:"size(32)"`
	HideHeader bool
	ThemeName string `orm:"size(20"`
}
func (this *Police)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Police)Insert()(int64,error){
	if this.PoliceId==""{
		this.PoliceId=random.NewId()
	}
	if this.PoliceRole==""{
		this.PoliceRole="user"//普通用户
	}
	return o.Insert(this)
}
func (this *Police)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Police)Delete()(int64,error){
	return o.Delete(this)
}
