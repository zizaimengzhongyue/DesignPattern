package reserve

import "fmt"

//Computer 电脑端
type Computer struct {
	Subject *Subject
	price   int
}

//Update 更新价格
func (c *Computer) Update() {
	fmt.Printf("电脑端当前价格是：%d\n", c.Subject.GetPrice())
}
