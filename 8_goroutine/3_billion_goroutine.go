package main

import "fmt"

func main() {
	ch := make(chan bool, 1000000)

	for i := 0; i < 1000000; i++ {
		go func() {
			fmt.Println("in", i)
			ch <- true
		}()
	}
	for i := 0; i < 1000000; i++ {
		<-ch
		fmt.Println("out", i)
	}
	fmt.Println("It's over!")

}
