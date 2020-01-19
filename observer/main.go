//Description
//    用户可以在通过手机，电脑和门店三种方式查看商品当前价格，在价格修改之后，同时向手机、电脑和门店同步
package main

import "./reserve"

func main() {
	subject := &reserve.Subject{}

	phone := &reserve.Phone{Subject: subject}
	subject.Attach(phone)
	computer := &reserve.Computer{Subject: subject}
	subject.Attach(computer)
	restaurant := &reserve.Restaurant{Subject: subject}
	subject.Attach(restaurant)

	subject.Update(15)
	subject.Update(17)
}
