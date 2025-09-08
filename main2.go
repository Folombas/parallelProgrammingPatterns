package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	doneCh := make(chan struct{})

	

	
	go func() {
		defer func() {
			doneCh<- struct{}{}
		}()
		ch <- 1
		ch <- 3
		ch <- 5
		ch <- 7
		}()

	
	go func() {
		defer func() {
			doneCh<- struct{}{}
		}()
		ch <- 2
		ch <- 4
		ch <- 6
		ch <- 8
	}()

	go func() {
		for i := 0; i < 2; i++ {
			<-doneCh
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")

}
