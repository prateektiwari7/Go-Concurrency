package main

// Code Highlights sync between Go routines !!

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := BooringCode("Joe!")
	c1 := BooringCode("An!")
	for i := 0; i < 5; i++ {
		fmt.Printf(" %q \n", <-c)
		fmt.Printf(" %q \n", <-c1)
	}
	fmt.Println("Leaving Main Thread")

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
