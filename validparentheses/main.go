package main

import "fmt"

func isValid(s string) bool {
	//create an empty stack to store opening brackets
	stack := make([]rune, 0)

	//note that the zero above is not necessary and it defaults to zero but its best practice
	// it could be started as a 5 and it would set it to [0 0 0 0 0] as opposed to completely empty ie []

	//define a mapping of closing brackets to their corresponding opening brackeets
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			//if its an opening bracket push it onto the stack
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			// if its a closing bracket check and see if it matches the top of the stack
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return false
			}
			//pop the matching opening bracket from the stack
			stack = stack[:len(stack)-1]
			// this is cool because it creates a new slice minus the top in a a slice
			//[12345] it would return or create[1234] and then replace stack with the new
			//slice
		}
	}

	//if the stack is empty at the end, it means all brackets were matched
	return len(stack) == 0

	// so for the gee wiz, we went through and matched whatever parens we found and
	// and if there were any left this return len(stack)==0 chooses whether it was true
	// or false if there are any left its false, if there aren't, all were matched and its
	// true

}

func main() {
	input := "()"
	isValid := isValid(input)
	fmt.Println("Is valid:", isValid)
}
