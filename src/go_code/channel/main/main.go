package main

import (
	"fmt"
	"sync"
	"time"
)

//利用互斥所解决goroutine资源竞争的问题。
var (
	myMap = make(map[int]int)
	lock  sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}
func main() {
	for i := 0; i < 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 10)
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("Mymap[%v] = %v\n", k, v)
	}
	lock.Unlock()
}
