**Two Sum**

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

**Solution**

It turns out we can do it in one-pass. 
While we are iterating and inserting elements into the hash table, we also look back to check if current element's complement already exists in the hash table. 
If it exists, we have found a solution and return the indices immediately.

Time complexity: __O(n)__. We traverse the list containing nn elements only once. Each lookup in the table costs only __O(1)__ time.

Space complexity: __O(n)__. The extra space required depends on the number of items stored in the hash table, which stores at most nn elements.