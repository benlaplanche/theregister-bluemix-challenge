package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	stringInput := os.Args[1]
	intInput, _ := strconv.ParseInt(stringInput, 10, 0)

	output := factorial(intInput)
	fmt.Println(output)
}

func factorial(number int64) (factor int64) {
	factor = 1

	if number == 1 {
		return
	} else {

		var counter int64 = 1

		for i := counter; i <= number; i++ {
			fmt.Println(factor)
			fmt.Println(i)
			factor = factor * i

		}
	}

	return factor
}
