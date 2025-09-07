package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")

}
