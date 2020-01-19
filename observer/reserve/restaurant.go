package reserve

import "fmt"

//Restaurant 店面端
type Restaurant struct {
	Subject *Subject
	price   int
}

//Update 更新价格
func (r *Restaurant) Update() {
	fmt.Printf("店面端当前价格是：%d\n", r.Subject.GetPrice())
}
