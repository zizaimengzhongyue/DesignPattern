package interviewer

import "fmt"

//Programmer 程序员面试官
type Programmer struct{}

//Question 问题
func (p Programmer) Question() {
	fmt.Println("你叫什么名字")
	fmt.Println("冒泡排序的复杂度是多少？")
}
