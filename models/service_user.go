package models
import (
	"strings"
	"gonganju/service"
	"github.com/astaxie/beego/orm"
	"gonganju/consts/errs"
	"time"
	"bytes"
	"github.com/astaxie/beego"
	"nong/pagination"
	"fmt"
)
type Polices []*Police
type Approves []*Approve
type CaseActions []*CasesAction
type CaseSuspectes []*CaseSuspects
/*获取案件处理的警官信息*/
func (this *Cases)GetPolices() (police Polices) {

	return
}
/*根据案件获取案件得处理流程*/
func (this *Cases)GetApproves() (approves Approves) {
	return
}
/*获取案件得执行信息*/
func (this *Cases)GetAction() (casesActions CaseActions) {
	return
}
func GetPoliceByAccPwd(acc, pwd string) (police *Police, rsp *service.Rsp) {
	rsp = &service.Rsp{}
	acc = strings.Trim(acc, " ")
	pwd = strings.Trim(pwd, " ")
	if acc == "" || pwd == "" {
		rsp.ResultID = 1
		rsp.ResultMsg = "账号密码不能为空"
		return
	}
	police = &Police{PoliceAcc:acc, PolicePwd:pwd}
	switch err := police.Read("PoliceAcc", "PolicePwd"); err {
	default:
		rsp.ResultID = errs.ErrSqlExecption
		rsp.ResultMsg = "数据库访问报错"
		police = nil
	case nil:
	case orm.ErrNoRows:
		rsp.ResultID = errs.ErrNoRow
		rsp.ResultMsg = "账号密码错误"
		police = nil
	}
	return
}
func GetPoliceById(id string) (police *Police, rsp *service.Rsp) {
	police = &Police{}
	rsp=new(service.Rsp)
	police.PoliceId = id
	switch err := police.Read(); err {
	default:
		rsp.ResultID = errs.ErrSqlExecption
		rsp.ResultMsg = "数据库访问报错"
		police = nil
	case nil:
	case orm.ErrNoRows:
		rsp.ResultID = errs.ErrNoRow
		rsp.ResultMsg = "用户已经删除"
		police = nil
	}
	return
}
/*获取个人的待处理案件*/
func SearchCases(page, numPerPage int, keyWord []string, startTime, endTime *time.Time, policeId string,
caseStatues int, suspectsId string, order[]string) (cases []*Cases, paginator *pagination.Paginator,err error) {
	where := &bytes.Buffer{}
	//where.WriteString(" 1=1")
	param := make([]interface{},0,10)
	//whereb:=false
	if len(keyWord)>0 {
		s2 := &bytes.Buffer{}
		for _, s := range keyWord {
			s2.WriteString(" or cases_all_info like ?")
			param=append(param, "%" + s + "%")
		}
		where.WriteString(" and (")
		where.Write(s2.Bytes()[4:])
		where.WriteString(")")

	}
	if startTime!=nil {
		where.WriteString(" and case_start_time>=?")
		param=append(param, startTime)
	}
	if endTime !=nil {
		where.WriteString(" and case_end_time<=?")
		param=append(param, endTime)
	}
	if caseStatues > 0 {
		where.WriteString(" and caseStatus=?")
		param=append(param,caseStatues)
	}
	innerJoin := ""
	if policeId != "" {
		innerJoin += " inner join case_police on case_police.cps_case_id=cases.cases_id and cps_police_id=?"
		param=append(param, policeId)
	}
	if suspectsId != "" {
		innerJoin += " innerjoin case_suspects on cas_cases_id=cases.case_id and cas_susecpects_id and cas_susecpects_id=?"
		param=append(param, suspectsId)
	}
	sql1 := ""
	sql2 := ""
	if where.Len() > 0 {
		sql1 = "select cases.* from cases" + string(where.Bytes()[5:]) + innerJoin
		sql2 = "select count(1) as totalRows from cases" + string(where.Bytes()[5:]) + innerJoin

	}else {
		sql1 = "select cases.* from cases" + innerJoin
		sql2 = "select count(1) from cases" + innerJoin
	}
	var maps []orm.Params
	_, err= o.Raw(sql2, param...).Values(&maps)
	if err != nil {
		beego.Error("查询案件总数时发生错误,sql:" ,sql2 , ",param:" ,param , "err:", err)
		return
	}
	totalRow, ok := maps[0]["totalRows"].(int)
	if !ok{
		totalRow=0
	}
	paginator=pagination.NewPaginator(numPerPage,totalRow,page)
	if totalRow==0{
		return
	}
	if len(order)>0{
		sql1+=" order by"
		for _,s:=range order{
			sql1+=" "+s
		}
	}
	switch  o.Driver().Type() {
	default:
		beego.Error("暂时支持mysql和postgres")
	case orm.DR_MySQL:
		sql1+=fmt.Sprintf(" limit %d ,%d",paginator.Offset(),paginator.PerPageNums)//	" limit "+paginator.Offset()+","+paginator.PerPageNums
	case orm.DR_Postgres:
		sql1+=fmt.Sprintf(" limit %d offset %d",paginator.PerPageNums,paginator.Offset())//" limit "+paginator.PerPageNums+" offset "+paginator.Offset()
	}
	_,err=o.Raw(sql1,param...).QueryRows(&cases)
	if err!=nil{
		beego.Error("查询案件失败，err:",err)
		return
	}
	return
}