package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0;i < 3; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")
	go f("go2")
	go f("go3")
	go f("go4")

	for i:=0; i<1000; i++ {
		go f("gf")
	}

	for i:=0; i<1000; i++ {
		fmt.Println("g", ":", i)
	}

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

