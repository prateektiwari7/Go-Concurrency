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
			select {
			case s := <-input1:
				c <- s
			// handling mutiple Case to catch
			//case s := <-input2:
			//c <- s
			case <-time.After(1 * time.Millisecond):
				fmt.Println("You are Slow Thread !")
				return

			}
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
