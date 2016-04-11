// Source : https://oj.leetcode.com/problems/roman-to-integer/
// Author : 18plusui
// Date   : 2016-04-12

/********************************************************************************** 
* 
* Given a roman numeral, convert it to an integer.
* 
* Input is guaranteed to be within the range from 1 to 3999.
*               
**********************************************************************************/

package main

import (
	"fmt"
)

func romanToIntHelper(str string) int{
	var d int 
	switch(str){
	
	case "I" :
		d = 1
		break
	case "V" :
		d = 5
		break
	case "X" :
		d = 10
		break
	case "L" :
		d = 50
		break
	case "C" :
		d = 100
		break
	case "D" :
		d = 500
		break
	case "M" :
		d = 100
		break
	}

	return d
}

func romanToInt (s string) int {
    
    if len(s) <= 0 {return 0}

    result := romanToIntHelper(s)
    for i := 1; i < len(s); i++ {
    	prev := romanToIntHelper(string(s[i-1]))
    	curr := romanToIntHelper(string(s[i]))

    	if prev < curr {
    		result = result - prev + (curr - prev)
    	} else {
    		result += curr
    	}
    }

    return result

}

func main() {
	fmt.Println(romanToInt("XL"))
}
