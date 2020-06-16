**Sort an array of 0’s, 1’s and 2’s (Dutch National Flag Problem)**

Given an array containing only 0’s, 1’s and 2’s, sort the array in linear time and using constant space.

We can rearrange the array in single traversal using an alternative <strong>linear-time partition routine</strong> can be used that separates the values into three groups:

    values less than the pivot
    values equal to the pivot and
    values greater than the pivot.
    
Complexity is _O(n)_
