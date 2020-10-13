package main

import (
	"fmt"
)

func main() {
	NewHash()
	SetKey("lc", 1)
	SetKey("zch", "nihao")
	SetKey("zy", 2)
	SetKey("wulaoshi", []int{10, 9, 8, 7})
	for _, v := range Arr {
		fmt.Printf("%+v\n", v.NextNode)
	}
	GetValueByKey("zch")
}
