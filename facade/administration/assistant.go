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

//ITSupport IT 支持
func (a Assistant) ITSupport() {
	a.IT.Support()
}

//HRSupport 人力支持
func (a Assistant) HRSupport() {
	a.HR.Support()
}
