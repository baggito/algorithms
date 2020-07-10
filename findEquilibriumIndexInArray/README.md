**Find Equilibrium Index of an Array**

Given an array of integers, find equilibrium index in it.


For an array A consisting n elements, index i is an equilibrium index if sum of elements of sub-array A[0..i-1] is equal to the sum of elements of sub-array A[i+1..n-1] i.e.


(A[0] + A[1] + ... + A[i-1]) = (A[i+1] + A[i+2] + ... + A[n-1])   where 0 < i < n-1

 
Similarly, 0 is an equilibrium index if (A[1] + A[2] + ... + A[n-1]) = 0 and n-1 is an equilibrium index if (A[0] + A[1] + ... + A[n-2]) = 0

**Solution:** 

The idea is to calculate the sum of all elements of the array. Then we start from the last element of the array and maintain right sub-array sum. We can calculate left sub-array sum in constant time by subtracting right sub-array sum and current element from total sum. i.e.

Sum of left subarray A[0..i-1] = total - (A[i] + sum of right subarray A[i+1..n-1])