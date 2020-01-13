package candidate

import "fmt"

//Salesman 销售候选人
type Salesman struct{}

//Answer 销售面试答案
func (s Salesman) Answer() {
	fmt.Println("我的名字叫李四")
	fmt.Println("上家公司年销售额是 5000 万")
}
