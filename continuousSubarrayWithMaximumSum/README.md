**Print continuous subarray with maximum sum**

Given an array of integers, find and print contiguous subarray with maximum sum in it.

For example,

**Input:**  
arr = [-2, 1, -3, 4, -1, 2, 1, -5, 4]  

**Output:**  
[4, -1, 2, 1]

**Solution:**  
We can easily solve this problem in linear time using Kadane’s algorithm by maintaining maximum sum subarray ending at each index of the array. 
Then this subarray is –

1) either empty in which case its sum is zero or
 
2) consists of one more element than the maximum subarray ending at the previous index
 