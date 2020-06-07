**Print all sub-arrays with 0 sum**

Given an unsorted array of integers, print all sub-arrays with 0 sum.


For example,

**Input:**  
arr = [4, 2, -3, -1, 0, 4]  

**Output:**  
{-3, -1, 0, 4}
{0}

**Solution:**  
_O(n^2) solution using MultiMap._  

We can use MultiMap to print all sub-arrays with 0 sum present in the given array. The idea is to create an empty multimap to store ending index of all sub-arrays having given sum. We traverse the given array, and maintain sum of elements seen so far. If sum is seen before, there exists at-least one sub-array with 0 sum which ends at current index and we print all such sub-arrays
