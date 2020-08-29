**Group elements of an array based on their first occurrence**

Given an unsorted array of integers containing many duplicates elements, rearrange the given array such that same element appears together and relative order of first occurrence of each element remains unchanged.

**Input:**  
arr = [5, 4, 5, 5, 3, 1, 2, 2, 4]  

**Output:**  
[5 5 5 4 4 3 1 2 2]

**Solution:**  
The idea is to use <a href="https://www.techiedelight.com/hashing-in-data-structure/" target="_blank" rel="noopener noreferrer">hashing</a> to solve this problem. We store frequency of each element present in the input array in a map. Then we traverse the input array and for each element A[i], two cases are possible –</p>
<ul>
<li>If A[i] exists in the map, then this is first occurrence of A[i] in input array. We print element A[i] k times where k is the frequency of A[i] in the input array (stored in map). Finally we delete A[i] from the map so it would not get processed again.<br>
&nbsp;</li>
<li>If A[i] is not present in the map, then this is repeated occurrence of A[i], and move to the next element.</li>
</ul>

The time complexity of above solution is _O(n)_ and auxiliary space used by the program is _O(n)_