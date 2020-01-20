//Descripiton
//    模板模式看起来就是继承
package main

import "./candidate"

func main() {
	programmer := &candidate.Programmer{Impl: candidate.Impl{Name: "张三"}}
	programmer.AboutMe()
	salesman := &candidate.Salesman{Impl: candidate.Impl{Name: "李四"}}
	salesman.AboutMe()
}
