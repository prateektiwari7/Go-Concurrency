package main

import (
	"fmt"
	"time"
)

func main() {
	go BooringCode("BooringCode Running!")
	fmt.Println("I'm Listening")
	time.Sleep(3 * time.Second)
	fmt.Println("You are Boring; I'm leaving")
}

func BooringCode(s string) {
	for i := 0; ; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}

}
