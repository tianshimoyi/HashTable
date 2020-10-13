package Hashs

import (
	"fmt"
	"math"
)

const (
	HashSize = 1003
	DefaultK = "tianshi"
	DefaultV = "tianshi"
)

var (
	Arr [HashSize]*Node
)

//TODO 哈希表初始化

func NewHash() {
	for i, _ := range Arr {
		Arr[i] = CreateHeadNode(DefaultK, DefaultV)
	}
}

//TODO 散列算法 除留余数法
func HashFunc(key string) int {
	index := 0
	for _, v := range key {
		index = int(v)
		index += int(math.Pow(float64(index), 3))
	}
	index = index % HashSize
	fmt.Println("初始index: ", index)
	return index
}

//开放地址法
func InsertHashKey(key string) int {
	index := HashFunc(key)
	for {
		if Arr[index].NextNode != nil {
			index = (index%HashSize + int(math.Pow(2, 2))) % HashSize
			fmt.Println("每次迭代的index: ", index)
			//time.Sleep(time.Second * 1)
		} else {
			break
		}
	}
	return index
}

//TODO 存入key值
func SetKey(k string, v interface{}) {
	index := InsertHashKey(k)
	fmt.Println("解决冲突后的index: ", index)
	AddNode(k, v, Arr[index])
	fmt.Printf("修改后2 %+v\n", Arr[index].NextNode)
}

//获取数据
func GetValueByKey(k string) {
	index := HashFunc(k)
	fmt.Println(index)
	for {
		if Arr[index].NextNode == nil || Arr[index].NextNode.Data.K != k {
			index = (index%HashSize + int(math.Pow(2, 2))) % HashSize
		} else {
			break
		}
	}
	fmt.Printf("%v\t%+v\n", k, Arr[index].NextNode)
}
