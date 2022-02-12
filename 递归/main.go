package main

import "fmt"

type Item struct {
	Id    int
	Type  string // doll diamond
	Child *Item
}

// 递归查询
func main() {
	items := Item{
		Type: "doll",
		Id:   1,
		Child: &Item{
			Type: "doll",
			Id:   2,
			Child: &Item{
				Type:  "diamond",
				Id:    3,
				Child: nil,
			},
		},
	}

	fmt.Println("查到的宝石是：", findDiamond(items))
}

// 递归查询
func findDiamond(item Item) Item {
	if item.Type == "doll" {
		return findDiamond(*item.Child)
	} else {
		return item
	}
}
