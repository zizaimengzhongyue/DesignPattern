package administration

//Administration 行政部门
type Administration interface {
	Support()
}

//Assistant 部门助理
type Assistant struct {
	IT IT
	HR HR
}

//SupportInit 支持部门初始化
func (a Assistant) SupportInit() {
	a.IT = IT{}
	a.HR = HR{}
}

//ITSupport IT 支持
func (a Assistant) ITSupport() {
	a.IT.Support()
}

//HRSupport 人力支持
func (a Assistant) HRSupport() {
	a.HR.Support()
}
