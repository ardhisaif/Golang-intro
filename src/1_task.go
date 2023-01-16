package src

import "fmt"

func PrintSegitiga(num int) {
	for i := 0; i < num; i++ { //untuk baris
		for j := 1; j <= num*2-1; j++ { //untuk print
			if j <= i || j >= num*2-i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}