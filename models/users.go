package models
import (
	"time"
	"nong/random"
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
/*
部门信息
*/
type Dept struct {
	DeptId string `orm:"pk;size(32)"`
	ParentId string `orm:"size(32)"`
	DeptName string  `orm:"size(32)"`
	DeptAddress string  `orm:"size(200)"`
	DeptTelephoneNumber string  `orm:"size(32)"`
}
func (this *Dept)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Dept)Insert()(int64,error){
	if this.DeptId==""{
		this.DeptId=random.NewId()
	}
	return o.Insert(this)
}
func (this *Dept)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Dept)Delete()(int64,error){
	return o.Delete(this)
}
/*
角色信息
*/
type Roles struct {
	RolesId string `orm:"pk;size(32)"`
	RolesName string `orm:"size(32)"`
	RolesAddTime time.Time	`orm:"auto_now_add"`
	RolesPermitCode int64
}
func (this *Roles)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *Roles)Insert()(int64,error){
	if this.RolesId==""{
		this.RolesId=random.NewId()
	}
	return o.Insert(this)
}
func (this *Roles)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *Roles)Delete()(int64,error){
	return o.Delete(this)
}
/*
角色对应的权限
*/
type RolesPermit struct {
	RolesPermitId string `orm:"pk;size(32)"`
	RolesId string `orm:"size(32)"`
	PermitId string `orm:"size(32)"`
}
func (this *RolesPermit)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *RolesPermit)Insert()(int64,error){
	if this.RolesPermitId==""{
		this.RolesPermitId=random.NewId()
	}
	return o.Insert(this)
}
func (this *RolesPermit)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *RolesPermit)Delete()(int64,error){
	return o.Delete(this)
}
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
/*
案件信息
*/
type Cases struct {
	CasesId string `orm:"pk;size(32)"`
	CasesName string `orm:"size(200)"`
	CasesCode string `orm:"size(32)"`//案件编号
	CasesInfo string `orm:"type(text)"`
	CasesTime time.Time
	CasesAddTime time.Time `orm:"auto_now_add"`
	CasesGrade float64 `orm:"type(numeric(3, 2))"`
	CasesStatus int //案件状态
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
/*
案件安排的人员
*/
type CasesPolice struct {
	CpolId string `orm:"pk;size(32)"`
	CpoCaseId string `orm:"size(32)"`
	CpoPoliceId string `orm:"size(32)"`
	CpoRole string `orm:"size(32)"`
	CpsGrade float64 `orm:"type(numeric(3, 2))"` //案件中每个警官应该获得的分数
	CpsDispatchId string `orm:"size(32)"`
	CpsAddTime time.Time `orm:"auto_now_add"`
}
func (this *CasesPolice)Read(field ...string) error {
	return o.Read(this, field...)
}
func (this *CasesPolice)Insert()(int64,error){
	if this.CpoCaseId==""{
		this.CpoCaseId=random.NewId()
	}
	return o.Insert(this)
}
func (this *CasesPolice)Update(field...string)(int64,error){
	return o.Update(this,field...)
}
func (this *CasesPolice)Delete()(int64,error){
	return o.Delete(this)
}
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
