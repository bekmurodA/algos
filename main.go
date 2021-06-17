package main

import (
	"fmt"
	"log"
	"strconv"
	"os"
)

func fib(n int) int {

	if n < 2 {
		return n
	} else {
		x := make(chan int)
		y := make(chan int)

		go func() {
			x <- fib(n - 1)
		}()

		go func() {
			y <- fib(n - 2)
		}()
		xVal := <-x
		yVal := <-y
		return (xVal + yVal)
	}

}
func main() {
	var n int
	var err error
	if len(os.Args) < 2 {
		fmt.Print("Please provide a number: ")
		_, err = fmt.Scanln(&n)
		if err != nil {
			log.Println(err)
		}
	} else {
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println(err)
		}
	}
	result := fib(n)
	fmt.Printf("fib of %d is %d\n", n, result)
}
