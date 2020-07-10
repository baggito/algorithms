**Move all zeros present in an array to the end**

Given an array of integers, move all zeros present in the array to the end. The solution should maintain the relative order of items in the array.
For example,

**Input:**  
arr = [6, 0, 8, 2, 3, 0, 4, 0, 1]  

**Output:**  
[6, 8, 2, 3, 4, 1, 0, 0, 0 ]

**Solution:**  
Using partitioning logic of QuickSort:
We can also solve this problem in one scan of array by modifying partitioning logic of quicksort. The idea is to use 0 as a pivot element and make one pass of partition process. The partitioning logic will read all elements and each time we encounter a non-pivot element, that element is swapped with the first occurrence of pivot.