//Description
//    需要针对候选人创建简历，其他简历可以通过复制第一份简历完成
//    还没想到具体的使用场景
package main

import "./resume"

func main() {
	resume01 := resume.Resume{}
	resume01.SetName("张三")
	resume01.SetMobile("123")

	resume02 := resume01.Clone()
	resume02.SetName("李四")
	resume02.SetMobile("456")

	resume01.GetName()
	resume01.GetMobile()

	resume02.GetName()
	resume02.GetMobile()
}
