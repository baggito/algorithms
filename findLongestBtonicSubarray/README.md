**Find Longest Bitonic Subarray in an array**

The longest bitonic subarray problem is to find a subarray of a given sequence in which the subarray’s elements are first sorted in in increasing order, then in decreasing order, and the subarray is as long as possible. Strictly ascending or descending subarrays are also accepted.


For example,

**Input:**  
arr = [3, 5, 8, 4, 5, 9, 10, 8, 5, 3, 4]  

**Output:**  
[4, 5, 9, 10, 8, 5, 3 ]

**Solution:** 

We can solve this problem without using extra space. The idea is to check for longest bitonic subarray starting at A[i]. If longest bitonic subarray starting at A[i] ends at A[j], the trick is to skip all elements between i and j as longest bitonic subarray starting from them will have less length. So, next we check for longest bitonic subarray starting at A[j]. We continue this process till end of array is reached and also keep track of longest bitonic subarray found so far.

The time complexity of above solution is _O(n)_ and auxiliary space used by the program is _O(1)_.