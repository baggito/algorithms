**Print all possible solutions to N Queens Problem**

The N queens puzzle is the problem of placing N chess queens on an N × N chessboard so that no two queens threaten each other. Thus, a solution requires that no two queens share the same row, column, or diagonal.

For example, for standard 8 × 8 chessboard below is one such configuration –

_Q - - - - - - -_  
_- - - - Q - - -_  
_- - - - - - - Q_    
_- - - - - Q - -_  
_- -Q - - - - -_   
_- - - - - - Q -_   
_- Q - - - - - -_   
_- - - Q - - - -_  

Note that the solution exist for all natural numbers n with the exception of n = 2 and n = 3.

**Solution:**  
We can solve this problem with the help of backtracking. The idea very simple. We start from the first row and place Queen in each square of the first row and recursively explore remaining rows to check if they leads to the solution or not. If current configuration doesn’t result in a solution, we backtrack. Before exploring any square, we ignore the square if two queens threaten each other.
