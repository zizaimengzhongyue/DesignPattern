//Description
//    后台提供了多种服务，客户端不需要去关注这些服务，只需要通过关键词调用这些服务就行
//    我还没理解为什么要再做一个 client，而不是用 delegate 直接执行
package main

import "./business"

func main() {
	delegate := &business.DeleGate{}
	delegate.SetServiceType("it")

	client := &business.Client{DeleGate: delegate}
	client.DoTask()

	delegate.SetServiceType("hr")
	client = &business.Client{DeleGate: delegate}
	client.DoTask()
}
