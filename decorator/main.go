//Description
//    面试包含第一轮和最后一轮面试，现在想增加一轮最终面
//    装饰器跟继承非常像，在 go 里面看起来没什么区别
package main

import "./interview"

func main() {
	programmer := interview.Programmer{}
	programmer.FirstInterview()
	programmer.SecondInterview()
	programmer.LastInterview()
}
