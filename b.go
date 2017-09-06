package main

import "fmt"

func main() {
	ch := make(chan int)
	go generater(ch)

	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filte(ch, ch1, prime)
		ch = ch1
	}
}

func generater(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filte(source <-chan int, dest chan<- int, prime int) {
	for i := range source {
		if i%prime != 0 {
			dest <- i
		}
	}
}
