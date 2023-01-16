package main

import (
	"fmt"

	"github.com/ardhisaif/golang_intro/src"
	// "time"
)

func main() {
	fmt.Println("masuk")
	//! task 1
	// src.PrintSegitiga(4)

	//! task 2
	genPass, err := src.GeneratePassword("UFDDHGS", "strong")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(genPass)

	//! task 3
	// var data = []int{1, 7, 3, 4, 8, 9}
	// result, err := src.CheckFilm(10, data)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(result)

}
