package service

type Rsp struct {
	ResultID int
	ResultMsg string
}
func (this *Rsp)SetResultId(resultId int){
	this.ResultID=resultId
}
func (this *Rsp)SetResultMsg(resultMsg string,){
	this.ResultMsg =resultMsg
}
func (this *Rsp)GetResultId() int{
	return this.ResultID
}
func (this *Rsp)GetResultMsg() string{
	return this.ResultMsg
}
type RspInterface interface {
	SetResultId(resultId int)
	SetResultMsg(resultMsg string)
	GetResultId() int
	GetResultMsg() string
}
