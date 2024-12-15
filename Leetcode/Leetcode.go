package leetcode

func TwoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)
	for i, I := range nums {
		diff := target - I
		if n, found := numToIndex[diff]; found {
			return []int{n, i}
		}

		numToIndex[I] = i
	}
	return []int{-1}
}

func IsPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return x == reversed || x == reversed/10
}

func RomanToInt(s string) int {
	charToInt := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	sum := 0
	for i, c := range s {
		a := charToInt[c]
		var b int
		if i != 0 {
			b = charToInt[rune(s[i-1])]
		}
		if a > b {
			sum = sum - b + a - b
		} else {
			sum += a
		}
	}
	return sum
}

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	while := true
	for ic := 0; while; ic++ {
		tmp := '\u0000'
		for _, v := range strs {
			if ic != len(v) && (tmp == '\u0000' || rune(v[ic]) == tmp) {
				tmp = rune(v[ic])
			} else {
				tmp = '\u0000'
				while = false
				break
			}
		}
		if while {
			prefix += string(tmp)
		}
	}
	return prefix
}

func IsValid(s string) bool {
	open := map[rune]int{
		'{': 1,
		'[': 2,
		'(': 3}
	close := map[rune]int{
		'}': 1,
		']': 2,
		')': 3}
	var stack []rune
	for _, c := range s {
		if _, found := open[c]; found {
			stack = append(stack, c)
		} else if _, found := close[c]; found {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if open[last] != close[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	var result []int
// 	for {

// 	}
// 	return result
// }
