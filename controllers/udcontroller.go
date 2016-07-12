package controllers

type UdController struct {
	BaseController
}

func (this *UdController) Test() {
	test := this.Ctx.Input.Params()
	s := ""
	for k, v := range test {
		s += k + ":" + v + "\n"
	}
	this.Ctx.WriteString(s)
}

func (this *UdController) Test1() {
	this.Ctx.WriteString("test1")
}
func (this *UdController) Test2() {
	this.Ctx.WriteString("test2")
}
func (this *UdController) Test3() {
	this.Ctx.WriteString("test3")
}
func (this *UdController) Test4() {
	this.Ctx.WriteString("test4")
}
