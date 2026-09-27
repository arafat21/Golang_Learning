package main

import "fmt"

func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i= i+ 1
	}

	// infinite loop
	// for{
	// 	println("1")
	// }

	// Classic for loop
	for i:=0;i<5;i++{
		if i ==2{
			continue
		}
		fmt.Println(i)
	}

	// range
	for i:= range 10{
		fmt.Println(i)
	}

}