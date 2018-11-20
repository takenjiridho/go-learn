package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	fmt.Println("skrg adalah jam  : ", t.Hour(), t.Minute())

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning !! ")
	case t.Hour() < 17:
		fmt.Println("Good aftarnoon !!")
	default:
		fmt.Println("good evening")

	}

}
