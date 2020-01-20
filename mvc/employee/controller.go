package employee

//Controller 控制器
type Controller struct {
	Emp  *Employee
	View *View
}

//UpdateEmployeeSalary 更新雇员薪资
//Params
//    salary int 薪资
func (c *Controller) UpdateEmployeeSalary(salary int) {
	c.Emp.SetSalary(salary)
}

//Show 展示员工信息
func (c *Controller) Show() {
	c.View.Show(c.Emp.GetName(), c.Emp.GetSalary())
}
