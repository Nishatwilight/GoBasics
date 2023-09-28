package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string `json:"full_name"`
	Weapon Weapon `json:"weapon"`
	Level  int    `json:"level"`
}
type Weapon struct {
	Name  string `json:"weapon_name"`
	Level int    `json:"weapon_level"`
}

func main() {
	fmt.Println("JSON Conversion")
	stringForm := `{"full_name": "Wallace","weapon":{"weapon_name":"weponName1, weponName2","weapon_level":6},"level": 1}`

	fmt.Println("stringForm", stringForm)
	var wallace Ninja
	fmt.Println("wallace", wallace)

	//Unmarshal Use json to make  Go  ==>converts json to go
	err := json.Unmarshal([]byte(stringForm), &wallace)
	if err != nil {
		fmt.Println("ohh noo!!!")
	}
	fmt.Println("After Unmarshal wallace", wallace)

	//Marshal Make Json use Go  ==>converts go to json
	jsontype, err := json.Marshal(wallace)
	if err != nil {
		fmt.Println("opps error: ", err)
	}
	fmt.Println("converts go to json: ", string(jsontype))
	fmt.Printf("using fmt.Printf %T\n", jsontype)
	fmt.Printf("using fmt.Printf %s\n", jsontype)
}
