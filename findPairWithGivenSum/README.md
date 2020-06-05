**Find pair with given sum in an array**

Given an unsorted array of integers, find a pair with given sum in it.


For example,

**Input:**  
arr = [8, 7, 2, 5, 3, 1]  
sum = 10

**Output:**  
Pair found at index 0 and 2

**Solution:**  
_O(n) solution using Map._  

We can use map to easily solve this problem in linear time. 
The idea is to insert each element of the array `arr[i]`
in a map. We also check if difference `(arr[i], sum-arr[i])`  already exists in the map or not. If the difference is seen before, we print the pair and return.