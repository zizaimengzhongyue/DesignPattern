//Description
//    看到解释器模式的第一个反应就是可以用来校验参数是否合法，实现了一个判断年龄是否为青少年的解释器
package main

import (
	"fmt"

	"./expression"
	"./expression/expressions"
)

//getTeenExpression 获取青少年解释器
//Return
//    *expression.AndExpression and 解释器，需要满足多个 expression
func getTeenExpression() *expression.AndExpression {
	teen := &expression.AndExpression{}
	age := &expressions.Match{Str: "Age"}
	teen.Add(age)
	ageRange := &expressions.Range{Min: 12, Max: 18}
	teen.Add(ageRange)
	return teen
}

//isTeen 是否青少年
//Params
//    val []interface{} 条件集合
func isTeen(val []interface{}) {
	teen := getTeenExpression()
	if teen.Interpret(val) {
		fmt.Printf("这是一个青少年：%v\n", val)
	} else {
		fmt.Printf("这不是一个青少年：%v\n", val)
	}
}

func main() {
	val01 := []interface{}{"Age", 12}
	isTeen(val01)
	val02 := []interface{}{"ID", 001}
	isTeen(val02)
	val03 := []interface{}{"Age", 22}
	isTeen(val03)
}
