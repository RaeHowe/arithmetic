package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	var l1 ListNode
	l1.Val = 1

	var l2 ListNode
	l2.Val = 2

	var l3 ListNode
	l3.Val = 3

	var l4 ListNode
	l4.Val = 3

	var l5 ListNode
	l5.Val = 2

	var l6 ListNode
	l6.Val = 1

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6

	result := isPalindrome(&l1)

	fmt.Println(result)
}

func isPalindrome(head *ListNode) bool {
	if head == nil{
		return false
	}

	var tmpArray []int

	for head != nil{
		tmpArray = append(tmpArray, head.Val)

		head = head.Next
	}

	var length = len(tmpArray) / 2

	for i := 0; i < length; i++{
		if tmpArray[i] != tmpArray[len(tmpArray) - 1 - i]{
			return false
		}
	}

	return true
}