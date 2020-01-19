package reserve

import "fmt"

//Phone 手机端
type Phone struct {
	Subject *Subject
	price   int
}

//Update 更新价格
func (p *Phone) Update() {
	fmt.Printf("手机端当前价格是：%d\n", p.Subject.GetPrice())
}
