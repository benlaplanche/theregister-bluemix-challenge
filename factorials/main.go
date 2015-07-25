package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide either a file path or a number input")

	} else {

		input := os.Args[1]
		file, err := os.Open(input)

		if check(err) == false {
			// looks like its a file path input
			defer file.Close()

			reader := bufio.NewReader(file)
			scanner := bufio.NewScanner(reader)

			for scanner.Scan() {
				value := scanner.Text()
				if value == "#" {
					return
				}

				value = cleanString(value)

				num := stringToInt(value)
				if num <= 15 {
					fmt.Println(factorial(num))
				}
			}

		} else if check(err) == true {
			// lets see if its a number string input
			intInput := stringToInt(input)

			output := factorial(intInput)
			fmt.Println(output)

		}

	}

}

func factorial(number int64) (factor int64) {
	factor = 1

	if number == 1 {
		return
	} else {

		var counter int64 = 1

		for i := counter; i <= number; i++ {
			factor = factor * i
		}
		return
	}

	return
}

func stringToInt(input string) int64 {
	result, _ := strconv.ParseInt(input, 10, 0)
	return result
}

func check(e error) bool {
	if e != nil {
		// encountered an error
		return true
	}
	return false
}

func cleanString(value string) (result string) {

	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllString(value, -1)

	for _, value := range matches {
		result += value
	}

	return
}
