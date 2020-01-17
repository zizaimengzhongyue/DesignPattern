package resume

//Programmer 程序员简历
type Programmer struct {
	Name   string
	Status int
}

//GetName 简历姓名
func (p *Programmer) GetName() string {
	return p.Name
}

//GetStatus 简历状态
func (p *Programmer) GetStatus() int {
	return p.Status
}
