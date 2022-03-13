package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var (
		str      string = ""
		number   string = ""
		i        float64
		Fizz     string = "Fizz"
		Buzz     string = "Buzz"
		FizzBuzz string = "FizzBuzz"
	)
	for i = 1; i <= 100; i++ {
		var indexModThree = math.Mod(i, 3)
		var indexModFive = math.Mod(i, 5)
		if indexModThree == 0 && indexModFive == 0 {
			str += FizzBuzz
			continue
		}
		if indexModThree == 0 {
			str += Fizz
			continue
		}

		if indexModFive == 0 {
			str += Buzz
			continue
		}
		number = strconv.FormatFloat(i, 'f', -1, 64)
		str += number
	}

	fmt.Println(str)
}
