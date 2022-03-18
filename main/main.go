package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// d(5)
	fmt.Println(d(5))
	c := time.After(5 * time.Second)

	ch := make(chan int)
	ch1 := make(chan int)
	// wg.Add(2)
	go func() {
		test(ch, 0)
	}()
	go func() {
		test(ch1, 1)
	}()

flag:
	for {
		select {
		// case j, ok := <-ch:
		// 	if !ok {
		// 		// break flag
		// 	}
		// 	fmt.Print("x:", j, ",")
		case j, ok := <-ch1:
			if !ok {
				break flag
			}
			fmt.Print("y:", j, ",")
		case <-time.After(3 * time.Second):
			fmt.Println("超时")
		case tt := <-c:
			fmt.Println(tt)
		}
		fmt.Println("----")

	}

}

func test(number chan int, typee int) {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		if typee == 0 {
			number <- x
			time.Sleep(1 * time.Second)
		} else {
			number <- y
			time.Sleep(5 * time.Second)
		}
		x, y = y, x+y

	}
	defer close(number)
}

func d(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return d(n-1) + (n - 2)
}
