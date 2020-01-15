//Description
//    行政部分别有 IT 部门和 HR 部门提供行政支持；一般部门都会有一个助理角色专门负责对接行政团队
//    在创建 assistant 对象的时候需要手动执行初始化操作，如果没有初始化直接调用 Support 方法会有问题，这里不是很棒
package main

import "./administration"

func main() {
	assistant := administration.Assistant{}
	assistant.SupportInit()
	assistant.ITSupport()
	assistant.HRSupport()
}
