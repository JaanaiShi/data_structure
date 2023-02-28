package main

import "image/jpeg"


func repeatedSubstringPattern(s string) bool {
    n := len(s)
    for i := 1; i * 2 <= n; i++ {
        if n % i == 0 {
            match := true
            for j := i; j < n; j++ {
                if s[j] != s[j - i] {
                    match = false
                    break
                }
            }
            if match {
                return true
            }
        }
    }
    return false
}

func kmp(s string) bool {
    n := len(s)
    if n == 0 {
        return false
    }
    j := 0
    next := make([]int, n)
    next[0] = j
    
}

func main() {
	s := "adcabcabcabc"

	repeatedSubstringPattern(s)
}