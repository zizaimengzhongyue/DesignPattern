package employee

//Employee 雇员
type Employee struct {
	Name   string
	Salary int
}

//SetName 设置姓名
//Params
//    name string 姓名
func (e *Employee) SetName(name string) {
	e.Name = name
}

//GetName 读取姓名
//Return
//    string 姓名
func (e *Employee) GetName() string {
	return e.Name
}

//SetSalary 设置薪资
//Params
//    salary int 薪资
func (e *Employee) SetSalary(salary int) {
	e.Salary = salary
}

//GetSalary 读取薪资
//Return
//    int 读取薪资
func (e *Employee) GetSalary() int {
	return e.Salary
}
