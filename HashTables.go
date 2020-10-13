package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	HashSize = 4
	DefaultK = "tianshi"
	DefaultV = "tianshi"
)

var Arr [HashSize]*Node

//TODO 哈希表初始化

func NewHash() {
	for i, _ := range Arr {
		Arr[i] = CreateHeadNode(DefaultK, DefaultV)
	}
}

//TODO 散列算法 除留余数法
func HashFunc(key string) int {
	index := 0
	rand.Seed(time.Now().Unix())
	for _, v := range key {
		index = int(v)
		index += rand.Intn(100)
	}
	fmt.Println("indexbefor: ", index)
	index = index % HashSize
	fmt.Println("index: ", index)
	for {
		if Arr[index].Data.K != DefaultK || Arr[index].Data.V != DefaultV {
			index = (index%HashSize + int(math.Pow(2, 2))) % HashSize
		} else {
			break
		}
	}
	return index
}

//TODO 存入key值
func SetKey(k string, v interface{}) {
	index := HashFunc(k)
	fmt.Println(index)
	//	head := Arr[index]
	AddNode(k, v, Arr[index])
	//	Arr[index].Data.K=
}

//获取数据
func GetValueByKey(k string) {
	index := HashFunc(k)
	fmt.Printf("%v\t%+v\n", k, *Arr[index].NextNode)
}
