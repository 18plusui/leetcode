// Source : https://oj.leetcode.com/problems/letter-combinations-of-a-phone-number/
// Author : 18plusui
// Date   : 2016-04-13

/********************************************************************************** 
* 
* Given a digit string, return all possible letter combinations that the number could represent.
* 
* A mapping of digit to letters (just like on the telephone buttons) is given below.
* 
* Input:Digit string "23"
* Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
* 
* Note:
* Although the above answer is in lexicographical order, your answer could be in any order you want.
*               
**********************************************************************************/
package main

import (
	"fmt"
	"unicode"
	"strconv"
)

func letterCombinations(str string) []string {
	match := [][]string{{" ","","",""},  //0
	                    {"","","",""},   //1
	                    {"a","b","c",""},//2
    					{"d","e","f",""},//3
    					{"g","h","i",""},//4
    					{"j","k","l",""},//5
    					{"m","n","o",""},//6
    					{"p","q","r","s"},//7
    					{"t","u","v",""},//8
    					{"w","x","y","z"},//9
   					   }

    var result []string
    if len(str) <= 0 {
    	
    	return append(result,"")
    }

    for i := 0; i < len(str); i++ {
    	if !unicode.IsDigit(rune(str[i])) {
    		return []string{""}
    	}

    	d,_ := strconv.Atoi(string(str[i]))

    	if len(result) <= 0 {
    		for j := 0; j < 4 && match[d][j]!=""; j++ {
    			result = append(result , match[d][j])
    	
    		}
    		continue
    	}
    
	    var r []string
	    for z := 0; z < len(result); z++ {

	    	for k := 0; k < 4 && match[d][k]!=""; k++ {
	    		r = append(r,result[z]+match[d][k])
	    		
	    	}
	    	
	    }
	    result = r
	}

    return result
}

func main() {
  	fmt.Println(letterCombinations("23"))
}
