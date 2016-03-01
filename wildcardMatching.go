package main

import (
    "fmt"
)

func IsMatch(s string, mask string) bool {
    if len(mask) == 0 && len(s) == 0 {
        return true
    }
    if len(mask) > 1 && mask[0] == '*' && len(s) == 0 {
        return false
    }
    if (len(mask) > 1 && mask[0] == '?') || 
       (len(mask) > 0 && len(s) > 0 && mask[0] == s[0]) {
        return IsMatch(s[1:], mask[1:])
    }
    if len(mask) > 0 && mask[0] == '*' {
        return IsMatch(s, mask[1:]) || IsMatch(s[1:], mask)
    }
    return false
}

func main() {
    fmt.Println(IsMatch("hello", "h?llo"))
    fmt.Println(IsMatch("hello", "h?lo"))
    fmt.Println(IsMatch("hello", "h*"))
    fmt.Println(IsMatch("hello", "h*lo"))
    fmt.Println(IsMatch("hello", "h*1"))
}
