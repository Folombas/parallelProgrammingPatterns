package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
		ch <- 3
		ch <- 5
		ch <- 7
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 2
		ch <- 4
		ch <- 6
		ch <- 8
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")

}
