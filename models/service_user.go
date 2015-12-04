package models
import (
	"strings"
	"gonganju/service"
	"github.com/astaxie/beego/orm"
	"gonganju/consts/errs"
)
type Polices []*Police
type Approves []*Approve
type CaseActions []*CasesAction
type CaseSuspectes []*CaseSuspects
/*获取案件处理的警官信息*/
func(this *Cases)GetPolices()(police Polices){

	return
}
/*根据案件获取案件得处理流程*/
func (this *Cases)GetApproves()(approves Approves){
	return
}
/*获取案件得执行信息*/
func (this *Cases)GetAction()(casesActions CaseActions){
	return
}
func GetPoliceByAccPwd(acc,pwd string)(police *Police,rsp *service.Rsp){
	rsp=&service.Rsp{}
	acc=strings.Trim(acc," ")
	pwd=strings.Trim(pwd," ")
	if acc==""||pwd==""{
		rsp.ResultID=1
		rsp.ResultMsg="账号密码不能为空"
		return
	}
	police=&Police{PoliceAcc:acc,PolicePwd:pwd}
	switch err:=police.Read("PoliceAcc","PolicePwd");err {
	default:
		rsp.ResultID =errs.ErrSqlExecption
		rsp.ResultMsg ="数据库访问报错"
		police=nil
	case nil:
	case orm.ErrNoRows:
		rsp.ResultID =errs.ErrNoRow
		rsp.ResultMsg ="账号密码错误"
		police=nil
	}
	return
}
func GetPoliceById(id string)(police *Police,rsp *service.Rsp){
	police=&Police{}
	police.PoliceId=id
	switch err := police.Read();err {
	default:
		rsp.ResultID =errs.ErrSqlExecption
		rsp.ResultMsg ="数据库访问报错"
		police=nil
	case nil:
	case orm.ErrNoRows:
		rsp.ResultID =errs.ErrNoRow
		rsp.ResultMsg ="用户已经删除"
		police=nil
	}
	return
}
