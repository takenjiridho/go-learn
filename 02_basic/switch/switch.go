package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("  go run on : ")
	switch v := runtime.GOOS; v {
	case "darwin":
		fmt.Println("mac osx ")
	case "linux":
		fmt.Println("linux.")
	default:
		fmt.Printf("%s.\n", v)
	}

}
