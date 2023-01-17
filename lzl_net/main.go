package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	write, err := w.Write([]byte("hello world"))
	if err != nil {
		return
	}
	fmt.Println(write)

}

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
