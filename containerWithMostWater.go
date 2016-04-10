// Source : https://oj.leetcode.com/problems/container-with-most-water/
// Author : 18plusui
// Date   : 2016-04-11

/********************************************************************************** 
* 
* Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). 
* n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). 
* 
* Find two lines, which together with x-axis forms a container, such that the container contains the most water.
* 
* Note: You may not slant the container.
*       
**********************************************************************************/

package main

import (
	"fmt"
)

func maxArea(height []int) int {
  var maxarea,left,right,area  int  =  0,0,len(height)-1,0

  for left < right{

    	area = (right - left) * calu3(height[left], height[right])
    	maxarea = calu3(area, maxarea)

      if height[left] < height[right] {

      	left++
      	
        for left < right && height[left-1] >= height[left] {
      		left++
      	}

      } else {
      	
        right--
      	
        for right > left && height[right+1] >= height[right] {
      		right--
      	}

      }
  }
  
  return maxarea

}

func calu3(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	inputArr := []int{1,2,6,5,2,1}
	fmt.Print(maxArea(inputArr))
	
}