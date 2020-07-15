**Find length of smallest subarray whose sum of elements is greater than the given number**

Given an array of positive integers, find the length of smallest subarray whose sum of elements is greater than the given number.

For example,

**Input:** {1, 2, 3, 4, 5, 6, 7, 8}, k = 20

**Output:** {6, 7, 8}

**Solution:** 

We can solve this problem by using a sliding window. The idea is to maintain a window that ends at the current element and sum of its elements is less than or equal to the given sum. If current window’s sum becomes more than the given sum at any point of time, then the window is unstable and continue removing elements from the window’ left till it becomes stable again. We also update the result if unstable window’s length is less than minimum found so far.

The time complexity of above solution is _O(n)_ and auxiliary space used by the program is _O(1)_.