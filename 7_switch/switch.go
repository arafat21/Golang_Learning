package main

import (
	"fmt"
	//"time"
)

func main() {
	// Basic

	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("others")
	// }

	// multiple condition

	// switch time.Now().Weekday(){
	// case time.Friday,time.Saturday:
	// 	fmt.Println("It's Weekend")
	// default:
	// 	fmt.Println("Working day")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("It's string")
		case bool:
			fmt.Println("It's bool")
		default:
			fmt.Println("others", t)

		}
	}

	whoAmI("55.6")
}
