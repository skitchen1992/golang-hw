package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var envName = os.Getenv("USER")

	if envName == "" {
		envName = "Guest"
	}

	fmt.Println("Аргументы CLI:")

	for index, argument := range os.Args[1:] {
		fmt.Printf("%d: %s\n", index+1, argument)
	}

	fmt.Println("Версия Go:", runtime.Version())

}
