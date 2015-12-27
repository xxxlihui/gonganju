package controllers
type Cases struct {
	Controller
}
func (c *Cases)List(){
c.TplNames="cases/list.html"
}
func (c* Cases)Add(){
	c.TplNames="cases/add.html"
}
func (c* Cases)Edit(){
c.TplNames="cases/edit.html"
}
func (c* Cases)Insert(){

}
func (c* Cases)Update(){

}


