package main

import (
	"strings"
	"unicode"
)

func RemoveEven(arr []int) (result []int) {
	for _, value := range arr {
		if value%2 != 0 {
			result = append(result, value)
		}
	}
	return result
}

func PowerGenerator(base int) func() int {
	lastNumber := 1
	return func() int {
		lastNumber *= base
		return lastNumber
	}
}

func DifferentWordsCount(str string) (result int) {
	str = strings.ToLower(str)
	var word string
	cnt := make(map[string]int)

	for i := 0; i < len(str); i++ {
		if !unicode.IsLetter(rune(str[i])) {
			if len(word) > 0 {
				cnt[word] += 1
			}
			word = ""
		} else {
			word += string(str[i])
		}
	}
	if len(word) > 0 {
		cnt[word] += 1
	}

	return len(cnt)
}
