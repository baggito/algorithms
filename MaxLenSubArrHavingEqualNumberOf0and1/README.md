**Find maximum length sub-array having equal number of 0’s and 1’s**

Given a binary array containing 0 and 1, find maximum length sub-array having equal number of 0’s and 1’s.  

**Input:**  
arr = [0, 0, 1, 0, 1, 0, 0]  

**Output:**  
The largest subarray length is 4

**Solution:**  
Complexity is _O(n)_

We can use map to solve this problem in linear time. The idea is to replace _0_ with _-1_ and find out the largest subarray with _0_ sum. To find largest subarray with 0 sum, we create an empty map which stores the ending index of first sub-array having given sum. We traverse the given array, and maintain sum of elements seen so far.
 
 1) If sum is seen for first time, insert the sum with its index into the map.
 2) If sum is seen before, there exists a sub-array with 0 sum which ends at current index and we update maximum length sub-array if current sub-array has more length.  
