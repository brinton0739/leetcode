package main

import "fmt"

func somethingSimple(numbers []int) []int {

	evenList := []int{}

	for _, num := range numbers {

		if num%2 == 0 {

			evenList = append(evenList, num)

		}

	}

	return evenList

}

func main() {

	listOne := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(somethingSimple(listOne))
}
