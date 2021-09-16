**Find the smallest missing element from a sorted array**

Given a sorted array of distinct non-negative integers, find the smallest missing element in it.

**Input:**  
arr = [0, 1, 2, 6, 9, 11, 15]  

**Output:**  
The smallest missing element is 3

**Input:**  
arr = [1, 2, 3, 4, 6, 9, 11, 15]  

**Output:**  
The smallest missing element is 0

**Input:**  
arr = [0, 1, 2, 3, 4, 5, 6]  

**Output:**  
The smallest missing element is 7

**Solution:**  
We go through the given array and on each iteration check the difference `arr[i+1] - arr[i]`, 
if the difference greater than `1` it means there is at least one missing number.  
Complexity is _O(n)_
