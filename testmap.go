package main

import "fmt"

func main() {
	var map1 = make(map[int]int)
	map1[0] = 1
	map1[2] = 1
	map1[2] = 1
	map1[2] = 1
	map1[2] = 1
	map1[21] = 1

	fmt.Println(map1)
}
