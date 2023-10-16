package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	n := 0

	for i := len(s) - 1; i >= 0; i-- {
		if n > 0 && s[i] == ' ' {
			break
		}

		if s[i] != ' ' {
			n++
		}
	}

	return n
}

// func lengthOfLastWord(s string) int {

// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] != ' ' {
// 			continue
// 		}
// 		fmt.Println(i)
// 		if s[i] == ' ' {
// 			return len(s[i+1:])
// 		}
// 	}

// 	return 0

// }

// func splitStringBySpace(s string) int {
// 	ss := []string{}
// 	start := 0
// 	for i, v := range s {
// 		if v == ' ' {
// 			if len(s[start:i]) > 0 {
// 				ss = append(ss, s[start:i])
// 			}
// 			start = i + 1
// 		} else if i == len(s)-1 {
// 			ss = append(ss, s[start:])
// 		}
// 	}
// 	fmt.Println(ss)
// 	return len(ss[len(ss)-1])
// }

func main() {
	s := "   fly me   to   the moon     "
	// fmt.Println(lengthOfLastWord(s))
	fmt.Println(lengthOfLastWord(s))

}
