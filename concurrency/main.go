package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int64

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func foo() {

	for i := 0; i < 100; i++ {
		fmt.Println("Foo :", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {

	for i := 0; i < 100; i++ {
		fmt.Println("Bar :", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

	wg.Done()

}

func incrementor(s string) {

	for i := 0; i < 20; i++ {

		time.Sleep(time.Duration(3 * time.Millisecond))
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, ":", i, ",", counter)

	}

	wg.Done()

}
