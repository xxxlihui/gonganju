package controllers
type PoliceInfo struct {
	Controller
}
func (c *PoliceInfo)Index(){
	c.Data["Paths"]=[...] paths {{Name:"个人信息统计"}}
	c.TplNames="police/policeinfo.html"
}