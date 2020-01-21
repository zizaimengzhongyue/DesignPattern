package transfer

//Employee 员工
type Employee struct {
	Name   string
	Salary int
}

//GetName 获取员工姓名
//Return
//    string 员工姓名
func (e *Employee) GetName() string {
	return e.Name
}

//SetName 设置姓名
//Params
//    name string 员工姓名
func (e *Employee) SetName(name string) {
	e.Name = name
}

//GetSalary 获取薪水
//Return
//    int 薪水
func (e *Employee) GetSalary() int {
	return e.Salary
}

//SetSalary 设置薪水
//Params
//    salary int 薪水
func (e *Employee) SetSalary(salary int) {
	e.Salary = salary
}
