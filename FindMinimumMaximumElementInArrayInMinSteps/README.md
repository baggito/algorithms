**Find Minimum and Maximum element in an array using minimum comparisons**

Given an array of integers, find out minimum and maximum element present using minimum comparisons.


For example,

**Input:**  
arr = [5, 7, 2, 4, 9, 6]  

**Output:**  
The minimum element in the array is 2
The maximum element in the array is 9

**Solution:**  
We can improve comparisons done by above solution by considering TWO elements at a time (i.e. consider elements in pairs). One special case we need to handle separately when array has odd number of elements.