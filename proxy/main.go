//Description
//    通过一个代理类来访问原始类，看起来跟桥接模式、适配器模式、装饰器模式很像
package main

import "./interview"

func main() {
	p := interview.Proxy{Interv: interview.Programmer{}}
	p.JobDescription()
}
