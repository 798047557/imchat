package main

import (
	"encoding/json"
	"fmt"
)

type student struct{
	A int `json:"id"`
	B string `json:"ww"`
}

func main() {
	//s := student{A:111,B:"BB"}
	//jsonStr,_ := json.Marshal(s)
	//fmt.Println(string(jsonStr))
	jsonStr := `{"id":111,"B":"bb"}`
	obj := &student{}
	json.Unmarshal([]byte(jsonStr),obj)
	fmt.Println(obj)



}