package main

import (
	"fmt"
)

// 剑指 Offer 20. 表示数值的字符串
func isNumber(s string) bool {
	hasNum := false
	hasE := false
	hasSign := false
	hasDot := false
	i := 0
	// 跳过空格
	for i < len(s) {
		if s[i] == ' ' {
			i++
			continue
		}
		break
	}
	// 进入循环
	for i < len(s) {
		if s[i] <= '9' && s[i] >= '0' {
			hasNum = true
			i++
		} else if i == len(s) {
			break
		} else if s[i] == 'e' || s[i] == 'E' {
			// 两个e或者没有之前没有数字
			if hasE || !hasNum {
				return false
			}
			hasE = true
			// 遍历后面的新数字
			hasNum = false
			hasSign = false
			hasDot = false
			i++
		} else if s[i] == '-' || s[i] == '+' {
			if hasDot || hasNum || hasSign {
				return false
			}
			hasSign = true
			i++
		} else if s[i] == '.' {
			// 有两个点或者e之后跟着小数
			if hasDot || hasE {
				return false
			}
			hasDot = true
			i++
		} else if s[i] == ' ' {
			// 末尾的空格或者数字中出现空格
			break
		} else {
			return false
		}
	}
	// 处理数字后的空格
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// 判断空格是数字中还是数字末尾，当为末尾的时候，还需要判断一下整个字符串是不是都是字符或者e之后没有整数
	return i == len(s) && hasNum == true
}

func exchange(nums []int) []int {
	head := 0
	tail := len(nums) - 1
	for head < tail {
		testHead := nums[head] % 2
		testTail := nums[tail] % 2
		if testHead == 0 && testTail != 0 {
			temp := nums[head]
			nums[head] = nums[tail]
			nums[tail] = temp
			head++
			tail--
			continue
		}
		if testHead != 0 && testTail != 0 {
			head++
			continue
		}
		if testHead == 0 && testTail == 0 {
			tail--
			continue
		}
		head++
		tail--
	}
	return nums
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 剑指 Offer 25. 合并两个排序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dum := new(ListNode)
	temp := dum
	for l1 != nil && l2 != nil {
		switch l1.Val < l2.Val {
		case true:
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
		case false:
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
		}
	}
	switch l1 == nil {
	case true:
		temp.Next = l2
	case false:
		temp.Next = l1
	}
	return dum.Next
}

func main() {
	// print(isNumber("6e6.5"))
	test := []int{1, 2, 3, 4}
	fmt.Printf("the new arrays is %v\n", exchange(test))
}
