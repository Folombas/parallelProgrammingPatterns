package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	doneCh := make(chan struct{})
	

	
	go func() {
		
		ch <- 1
		ch <- 3
		ch <- 5
		ch <- 7
	}()

	
	go func() {
		
		ch <- 2
		ch <- 4
		ch <- 6
		ch <- 8
	}()

	go func() {
		
		<-doneCh
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")

}
