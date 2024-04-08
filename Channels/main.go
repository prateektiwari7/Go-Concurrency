package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go BooringCode("Task!", c)
	fmt.Println("I am in Main Thread")
	for i := 0; i < 5; i++ {
		fmt.Printf("Listening....%q\n", <-c)
	}
	fmt.Println("Leaving the MainThread")

}

func BooringCode(s string, c chan string) {
	for i := 0; ; i++ {
		fmt.Println("In Parallel Universe")
		c <- fmt.Sprintf("%s,%d", s, i)
		time.Sleep(time.Duration(rand.Intn(1e3)))
	}

}
