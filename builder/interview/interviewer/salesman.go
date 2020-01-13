package interviewer

import "fmt"

//Salesman 销售面试官
type Salesman struct{}

//Salesman 销售面试问题
func (s Salesman) Question() {
	fmt.Println("你叫什么名字")
	fmt.Println("上家公司的年销售额是多少")
}
