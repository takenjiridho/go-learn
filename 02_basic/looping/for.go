package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("i adalah ", i)
	}

	sum := 0

	for x := 0; x < 10; x++ {
		fmt.Println("sum ", sum, "+= ", x)
		sum += x
	}

	fmt.Println(sum)

	sum2 := 1

	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	l := 0

	for {

		l++
		fmt.Printf("break l  %d\n", l)

		if l == 5 {
			break
		}
	}

	c := 0

	for {

		c++

		if c == 5 {
			continue
		}

		fmt.Printf("continue l %d\n", c)

		if c == 10 {
			break
		}
	}

}
