package models
import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"nong/random"
	"os"
	"io"
)
var o orm.Ormer
func Init(){
	RegisterModels()
	beego.Debug("models.inint 初始化数据库")
	orm.RegisterDriver("postgres", orm.DR_Postgres)
	maxIdle := beego.AppConfig.DefaultInt("db_maxIdle", 30)
	maxConn := beego.AppConfig.DefaultInt("db_maxConn", 50)
	//postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	conStr := beego.AppConfig.String("db_conStr")
	beego.Info("数据库链接：",conStr,",maxIdle:",maxIdle,",maxConn:",maxConn)
	err := orm.RegisterDataBase("default", "postgres", conStr, maxIdle, maxConn)
	if err!=nil {
		beego.Error("初始化数据库失败\r\n", err)
		return
	}
	orm.RegisterModel(

	)
	beego.Info("初始化数据库成功")
	name := "default"

	// drop table 后再建表
	force := false

	// 打印执行过程
	verbose := true

	err= orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error("创建表失败", err)
	}
	o=orm.NewOrm()
}
func RegisterModels(){
	orm.RegisterModel(
		new(Police),new(CasesAction),new(Cases),new(Permit),
		new(Photos),new(Approve),new(Suspects),new(RolesPermit),
		new(Roles),new(Dept),new(CasePoliceSuspects),new(CaseSuspects),
		new(CasesPolice),
	)

}
func AddAdmin(){
	adminRoleId:="67C96D63288C4842A75B9EBC1E6243D7"
	adminPermitId:=1
	/*添加默认管理员*/
	police:=new(Police)
	police.PoliceAcc="admin"
	police.PoliceIcon=random.NewId()
	police.PolicePwd="123"
	police.PoliceName="管理员"
	police.Insert()
	/*添加管理员图标*/
	photo:=new(Photos)
	photo.PhotosId=police.PoliceIcon
	photo.PhotosInfo="管理员图像"
	photo.Insert()
	/*添加管理员角色*/
	roles:=new(Roles)
	roles.RolesId=adminRoleId
	roles.RolesName="系统管理员"
	roles.RolesPermitCode

	photo.PhotosExt=".png"
	file,err:=os.Open("conf/adminIcon.png")
	if err!=nil{
		return
	}
	defer file.Close()
	fileIcon,err2:=os.OpenFile("static/photos/"+photo.PhotosId+photo.PhotosExt,os.O_CREATE|os.O_RDONLY|os.O_TRUNC,1)
	if err2!=nil{
		return
	}
	defer fileIcon.Close()
	io.Copy(file,fileIcon)

}
func init(){
	Init()

}

