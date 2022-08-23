package main

import "fmt"

func main() {

	var numbers [3]int = [3]int{4, 5, 6}

	// This would also work, althoug IDEA would auto remove it.
	// var numbers [3]int = [3]int{4, 5, 6,}

	fmt.Println(numbers)

	// If the closing curly bracket is on the next line the previous line needs to have
	// a comma, other wise the declaration would be incorrect. This is probably due to
	// go compiler automatically adding semicolon at the end of line. Notice that with the
	// declaration on line#7 the comma wasn't necesaary although adding it wouldn't cause any
	// error.
	var numsGreaterThanTen [4]int = [4]int{
		11,
		12,
		13,
	}
	fmt.Println(numsGreaterThanTen)
}
