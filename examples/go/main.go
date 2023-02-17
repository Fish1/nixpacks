package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello from Go!")

	fmt.Println(len(os.Args), os.Args)
}
