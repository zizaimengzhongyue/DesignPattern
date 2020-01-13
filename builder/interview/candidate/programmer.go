package candidate

import "fmt"

//Programmer 程序员候选人
type Programmer struct{}

//Answer 回答问题
func (p Programmer) Answer() {
	fmt.Println("我的名字叫张三")
	fmt.Println("冒泡排序的时间复杂度是O(n^2)")
}
