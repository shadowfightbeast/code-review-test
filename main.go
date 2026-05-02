package main

import "fmt"

func main() {
	// Intentionally messy code for the bot to review
	var x int = 10
	if x > 5 {
		fmt.Println("X is greater than 5")
	}

	// Unused variable
	var unusedStr string = "I am not used"

	// Potential nil pointer issue
	//bot will find this issue.

	var data *string
	fmt.Println(*data)
}
