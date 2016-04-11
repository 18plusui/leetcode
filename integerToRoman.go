// Source : https://oj.leetcode.com/problems/integer-to-roman/
// Author : 18plusui
// Date   : 2016-04-12

/********************************************************************************** 
* 
* Given an integer, convert it to a roman numeral.
* 
* Input is guaranteed to be within the range from 1 to 3999.
*               
**********************************************************************************/
package main

import (
   "fmt"
)

func intToRoman(num int) string {

   symbol := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
   value  := []int{1000 ,900 ,500 ,400 , 100, 90,  50, 40, 10, 9, 5, 4, 1} 
   var result string

   for i := 0; num !=0; i++ {

       for num >= value[i] {
           num -= value[i]
           result += symbol[i]
      }
  }

  return result
}

func main() {
	fmt.Println(intToRoman(10))
}