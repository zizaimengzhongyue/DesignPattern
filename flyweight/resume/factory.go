package resume

//Factory ProgrammerResume 工厂类
type Factory struct {
	MapResume map[string]*ProgrammerResume
}

//GetInstance 创建 ProgrammerResume 实例
//Params
//    name string 简历姓名
//Return
//    *ProgrammerResume 程序员简历
func (f Factory) GetInstance(name string) *ProgrammerResume {
	if _, ok := f.MapResume[name]; ok {
		return f.MapResume[name]
	}
	resume := &ProgrammerResume{}
	resume.SetName(name)
	f.MapResume[name] = resume
	return f.MapResume[name]
}
