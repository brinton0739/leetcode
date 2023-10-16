package main

import "fmt"

//so first we want to create a struct for the single merged list
// this creates more dynamic use of pointers and list size that isn't
// bound to rules that slices and arrays are

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(listOne *ListNode, listTwo *ListNode) *ListNode {

	// we need a dummy node to simplify the code
	dummy := &ListNode{}
	current := dummy

	//iterate through both lists

	for listOne != nil && listTwo != nil {
		//loop will keep running while both lists are not nil
		if listOne.Val < listTwo.Val {
			// when the number in list one is less
			current.Next = listOne
			//place whatever is in list one to the current position
			listOne = listOne.Next
			//move to the next position in list one
		} else {
			current.Next = listTwo
			//if the list one value wasn't less then place the number in list two onto
			//the current list
			listTwo = listTwo.Next
			//move to the next place in list two and move
		}

		current = current.Next
	}
	//If one of the lists is not exhausted, append the remaning nodes

	if listOne != nil {
		current.Next = listOne
	} else {
		current.Next = listTwo
	}

	return dummy.Next
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}

	// Merge the two lists
	merged := mergeTwoLists(list1, list2)

	// Print the merged list
	fmt.Println("Merged Linked List:")
	printLinkedList(merged)
}
