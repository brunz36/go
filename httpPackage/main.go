package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	response, err := http.Get("https://google.com/")

	if err != nil {
		fmt.Println("Error with GET Command :: ", err)
		os.Exit(1)
	}

	//bs := make([]byte, 9999999)
	//read, err := response.Body.Read(bs)
	//if err != io.EOF {
	//	fmt.Println("ERROR ", err)
	//}
	//fmt.Println("The GET Response :: ", read, string(bs))

	lw := logWriter{}

	io.Copy(lw, response.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}
