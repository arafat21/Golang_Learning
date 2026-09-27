package main

import "fmt"

func main() {
	// age := 22

	// if age > 18 {
	// 	fmt.Println("Voter")
	// } else{
	// 	fmt.Println("Not eligible for voting")
	// }

	// Else if uses

	// age:=7
	// if age>=18{
	// 	fmt.Println("Adult")
	// }else if age>=12{
	// 	fmt.Println("teenager")
	// }else{
	// 	fmt.Println("kid")
	// }

	// var role = "admin"
	// var hasPermission = true
	// if role == "admin" || hasPermission{
	// 	fmt.Println("yes")
	// }

	if age:=15; age>=18{
		fmt.Println("Person is adult",age)
	} else if age>=12{
		fmt.Println("Person is teenager",age)
	} else{
		fmt.Println("Person is kid",age)
	}
}
