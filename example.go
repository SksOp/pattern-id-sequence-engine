package main

import (
	"fmt"

	"github.com/pattern-id-sequence-engine/utils"
)

// Test one with built in configration
func main() {
	/*
	 max N supported in sequence is 18
	*/
	pattern := "DDYYYMMN"

	next, err := utils.GetNextId(pattern, "")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(next)

	nextId, err := utils.GetNextId(pattern, next)
	if err != nil {
		fmt.Println("error in creating next id", err)
		return
	}
	fmt.Println(nextId)

	for i := 1; i < 10; i++ {
		nextId, err = utils.GetNextId(pattern, nextId)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(nextId)
	}
}

// // Test second with built in configration with custom fields

// func main() {
// 	/*
// 	 max N supported in sequence is 18
// 	*/

// 	pattern := "DDYYKKYMMUUUUUUN"

// 	custom := utils.CustomeSupport{
// 		Customs: map[string]string{
// 			"KK":     "00",
// 			"UUUUUU": "696969",
// 		},
// 	}

// 	next, err := utils.GetNextId(pattern, "", custom)
// 	if err != nil {
// 		fmt.Printf("%s\n", err)
// 	}
// 	fmt.Println(next)

// 	nextId, err := utils.GetNextId(pattern, next, custom)
// 	if err != nil {
// 		fmt.Println("error in creating next id", err)
// 		return
// 	}
// 	fmt.Println(nextId)

// 	for i := 1; i < 10; i++ {
// 		nextId, err = utils.GetNextId(pattern, nextId, custom)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println(nextId)
// 	}
// 	fmt.Println(nextId)

// }
