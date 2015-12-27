package models
import "gz.com/master/random"

/*
图片信息
*/
type Photos struct {
	PhotosId string `orm:"pk;size(32)"`
	PhotosCasesId string `orm:"size(32)"`
	PhotosSuspectsId string `orm:"size(32)"`
	PhotosCacId string `orm:"size(32)"`
	PhotosPoliceId string `orm:"size(32)"`
	PhotosInfo string `orm:"size(400)"`
	PhotosExt string `orm:"size(20)"`
}
func (this *Photos)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Photos)Insert()(int64,error){
	if this.PhotosId==""{
		this.PhotosId=random.NewId()
	}
	return o.Insert(this)
}
func (this *Photos)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Photos)Delete()(int64,error){
	return o.Delete(this)
}
