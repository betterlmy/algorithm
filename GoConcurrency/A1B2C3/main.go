package main

import (
	"fmt"
	"sync"
	"time"
)

var abcChannel, onettChannel = make(chan struct{}), make(chan struct{})

// 交替打印cat dog hello

func printABC(wg *sync.WaitGroup) {
	defer wg.Done()
	i := int32(0)
	x := 'A'

	for {
		<-abcChannel
		fmt.Printf("%v", string(x+i))
		i++
		i %= 26
		time.Sleep(time.Second)
		onettChannel <- struct{}{}
	}
}
func print123(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	for {
		<-onettChannel
		fmt.Printf("%v", i)
		i++
		i %= 26
		time.Sleep(time.Second)
		abcChannel <- struct{}{}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go printABC(wg)
	go print123(wg)
	onettChannel <- struct{}{}
	wg.Wait()
}
