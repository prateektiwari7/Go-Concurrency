package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quit := make(chan bool)
	c := booring("Joe", quit)
	for i := rand.Intn(100); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

func booring(s string, quit chan bool) <-chan string {
	r := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case r <- fmt.Sprintf("%s : %d", s, i):

			case <-quit:
				return
			}
		}

	}()

	return r

}
