package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := FanIn(BooringCode("Joe"), BooringCode("An"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Boring Code Leaving")

}

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func BooringCode(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s,%d", s, i)
			time.Sleep(time.Duration(rand.Intn(1e3)))
		}

	}()

	return c

}
