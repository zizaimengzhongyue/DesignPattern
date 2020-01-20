package employee

import "fmt"

//View 展示层
type View struct{}

//Show 展示员工信息
//Params
//    name string 员工姓名
//    salary int 员工薪资
func (v *View) Show(name string, salary int) {
	fmt.Printf("这名员工的姓名是 %s，薪资是 %d\n", name, salary)
}
