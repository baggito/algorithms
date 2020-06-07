**Sort Binary Array in Linear Time**

Given a binary array, sort it in linear time and constant space. Output should print contain all zeroes followed by all ones.


For example,

**Input:**  
arr = [1, 0, 1, 0, 1, 0, 0, 1]  

**Output:**  
[0, 0, 0, 0, 1, 1, 1, 1]

**Solution:**  
The time complexity of solution is _O(n)_ and auxiliary space used by the program is _O(1)_.  

We can solve this problem in linear time by using partitioning logic of quicksort.
