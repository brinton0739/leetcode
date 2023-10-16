package main

import (
	"fmt"
)

func romanToInt(roman string) int {

	//we create a map of what each letter means in integers

	romanMap := map[byte]int{

		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	// iterate through the roman numeral string.
	for i := 0; i < len(roman); i++ {
		currentValue := romanMap[roman[i]]

		// in roman numerals if the number immediately preceeding a larger number it is
		//subtracted from the number in front of it other wise it is added to the number
		if i+1 < len(roman) {
			nextValue := romanMap[roman[i+1]]
			if currentValue < nextValue {
				//if it is, subtract current value.
				result -= currentValue
			} else {
				//otherwise add it
				result += currentValue
			}
		} else {
			// if there are no more symbols to compare add the current value
			result += currentValue
		}

	}
	return result

}

func main() {
	romanNumeral := "MCMLXXIX"
	integerValue := romanToInt(romanNumeral)

	fmt.Printf("Roman Numeral: %s\n", romanNumeral)
	fmt.Printf("Integer Value: %d\n", integerValue)
}
