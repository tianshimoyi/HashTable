package test

import (
	"DataStructure/HashTable/Hashs"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	Hashs.NewHash()
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().Unix())
		v := rand.Intn(100000)
		k := 2*i + v - 1056
		key := strconv.Itoa(k)
		Hashs.SetKey(key, i)
	}
	for _, v := range Hashs.Arr {
		fmt.Printf("%+v\n", v)
	}
	
}
