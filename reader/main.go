package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1]
	//file, err := os.Open(fileName)
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading file from disk.", err)
		os.Exit(1)
	}

	fmt.Println(string(file))

	//io.Copy(os.Stdout, file)
}
