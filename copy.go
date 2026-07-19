package main

import (
	"fmt"
)

func main() {
	fmt.Println("Valid Parenthesis Checker Started...")
	//testing := "((){})"
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	//fmt.Println(checker(testing))
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) (int, []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			fmt.Println("kth value: ", nums[k])
			fmt.Println("ith value: ", nums[i])
			fmt.Println("Spaced")
			nums[k] = nums[i]
			k++
		}
	}
	return k, nums[:k]
}

// func checker(s string) bool {
// 	if len(s)%2 != 0 {
// 		return false
// 	}

// 	pairs := map[rune]rune{
// 		')': '(',
// 		'}': '{',
// 		']': '[',
// 	}

// 	stack := []rune{}

// 	for _, ch := range s {
// 		switch ch {
// 		case '(', '{', '[':
// 			stack = append(stack, ch) // ["(", "("] || ["(", "{"]
// 		case ')', '}', ']':
// 			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1]
// 		}
// 	}

// 	return len(stack) == 0
// }
