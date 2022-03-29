package main

import "fmt"

func main() {
	ch := make(chan []byte,2)
	ch <- []byte("wocao")
	ch <- []byte("wocao")
	//ch <- []byte("wocao")

	fmt.Println(<-ch)
	//fmt.Println([]byte("我草"))
}