package main

import "fmt"

func multires(x, y string) (string, string) {
	return x, y

}

func main() {
	a, b := multires("ganteng", "banget ya")
	fmt.Printf("ridho %v , %v\n", a, b)
}
