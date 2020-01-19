//Description
//    假设一个公司的人力资源会共享人才，有面试者面试通过之后会把该面试者信息转发给公司其他部门看是否有感兴趣的，如果部门之间交叉通信就太麻烦了，
//  更合适的方式是建立一个人才库，通过这个人才库可以共享所有面试者的信息
//    感觉这个 demo 写的不是很棒
package main

import (
	"./candidate"
	"./department"
	"./talents"
)

func main() {
	talent := &talents.Talents{}

	salesman := &candidate.Impl{Name: "张三", Carrer: "销售"}
	sales := department.Sales{Talents: talent}

	programmer := &candidate.Impl{Name: "李四", Carrer: "程序员"}
	technology := department.Technology{Talents: talent}

	sales.ShowCandidate(salesman)
	technology.ShowCandidate(programmer)
}
