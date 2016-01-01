package controllers
type Cases struct {
	Controller
}

func (c *Cases)Assigned(){
	c.Data["Paths"] = [...] paths{{Name:"案件指派"}}
	c.Data["PanelName"] = "案件列表"
	c.Data["PanelTitle"] = "案件指派列表"
	c.TplNames = "cases/assigned.html"
}
func (c *Cases)List() {
	c.Data["Paths"] = [...] paths{{Name:"案件指派"}}
	c.Data["PanelName"] = "案件列表"
	c.Data["PanelTitle"] = "案件指派列表"
	c.TplNames = "cases/list.html"
}
func (c* Cases)Add() {
	c.TplNames = "cases/add.html"
}
func (c* Cases)Edit() {
	c.TplNames = "cases/edit.html"
}
func (c* Cases)Insert() {

}
func (c* Cases)Update() {

}


