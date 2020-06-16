**Rearrange an array with alternate high and low elements**

Given an array of integers, rearrange the array such that every second element of the array is greater than its left and right elements. Assume no duplicate elements are present in the array.

For example,

**Input:**  
arr = [1, 2, 3, 4, 5, 6, 7]  

**Output:**  
[1, 3, 2, 5, 4, 7, 6]

**Solution:**  
An efficient solution doesn’t involve sorting the array or use of auxiliary space. We start from the second element of the array and increment index by 2 for each iteration of loop. If previous element is greater than the current element, we swap the elements. Similarly if next element is greater than the current element, we swap both elements. At the end of loop, we will get the desired array that satisfies given constraints.
    
Complexity is _O(n)_
