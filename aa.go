package main

import "fmt"

func main() {
	defer throw()

	a := 1;
	if a == 1{
		panic("炸了")
	}
}

func throw(){
	if r := recover();r != nil{
		fmt.Println(r)
	}
}
