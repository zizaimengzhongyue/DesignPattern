//Description
//    创建一份菜单，用户查询食品信息，如果有则返回该食品，没有就返回一个 null 对象
package main

import (
	"fmt"

	"./food"
)

//获取菜单
func getMenu() *food.Menu {
	lists := []string{"小米", "牛肉", "面"}
	menu := &food.Menu{}
	for _, v := range lists {
		menu.Add(&food.Real{Name: v})
	}
	return menu
}

func main() {
	menu := getMenu()

	beef := menu.GetFood("牛肉")
	fmt.Println(beef.GetName())

	pepsi := menu.GetFood("百事可乐")
	fmt.Println(pepsi.GetName())
}
