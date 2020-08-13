**Print all Possible Knight’s Tours in a chessboard**

Given a chess board, print all sequences of moves of a knight on a chessboard such that the knight visits every square only once.

For example, for standard 8 × 8 chessboard below is one such tour. We have started the tour from top-leftmost of the board (marked as 1) and consecutive moves of the knight are represented by the next number.

1 50 45 62 31 18 9 64<br>
46 61 32 49 10 63 30 17<br>
51 &nbsp;2 47 44 33 28 19 8<br>
60 35 42 27 48 11 16 29<br>
41 52 &nbsp;3 34 43 24 7 20<br>
36 59 38 55 26 21 12 15<br>
53 40 57 &nbsp;4 23 14 25 6<br>
58 37 54 39 56 5 22 13<br>  


The Knight should search for a path from the starting position until it visits every square or until it exhausts all possibilities. We can easily achieve this with the help of backtracking. We start from given source square in the chessboard and recursively explore all eight paths possible to check if they leads to the solution or not. If current path doesn’t reach destination or we have explored all possible routes from the current square, we backtrack. To make sure that the path is simple and doesn’t contain any cycles, we keep track of squares involved in current path in a chessboard and before exploring any square, we ignore the square if it is already covered in current path.

**Solution:**  
We can solve this problem with the help of backtracking. The idea very simple. We start from the first row and place Queen in each square of the first row and recursively explore remaining rows to check if they leads to the solution or not. If current configuration doesn’t result in a solution, we backtrack. Before exploring any square, we ignore the square if two queens threaten each other.
