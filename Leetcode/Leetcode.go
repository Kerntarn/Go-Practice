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
