package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("when 's sunday?")
	today := time.Now().Weekday()

	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far away", today+1, " ")
	}

}
