package controllers
type PoliceInfo struct {
	Controller
}
func (c *PoliceInfo)Index(){
	c.Data["Paths"]=[...] paths {{Name:"个人信息统计"}}
	c.Data["PanelName"]="个人统计"
	c.Data["PanelTitle"]=c.Police.PoliceName+"的统计信息"
	c.TplNames="police/policeinfo.html"
}