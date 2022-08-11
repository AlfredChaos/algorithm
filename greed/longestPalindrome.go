package main

import "fmt"

func longestPalindrome(s string) int {
    if len(s) == 1 {
        return 1
    }
    result := 0
    duplicateMap := make(map[byte]bool)
    for i:=0; i<len(s); i++ {
        if _, ok := duplicateMap[s[i]]; ok {
            result += 2
            delete(duplicateMap, s[i])
			continue
        }
        duplicateMap[s[i]] = true
    }
    if len(duplicateMap) != 0 {
        result += 1
    }
    return result
}

func main() {
	s := "bb"
	fmt.Println(longestPalindrome(s))
}