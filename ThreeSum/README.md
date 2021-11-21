**Three Sum**

Given an array of distinct integers nums and an integer target, return indices of the three numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

**Solution**

It turns out we can do it in one-pass.
While we are iterating and inserting elements into the hash table, we also look back to check if current element's complement already exists in the hash table.
If it exists, we have found a solution and return the indices immediately.

Time complexity: __O(NlogN)__.

Space complexity: __O(n)__. 