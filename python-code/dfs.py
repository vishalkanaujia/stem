matrix = [
    [0,0,0],
    [0,0,0],
    [0,0,0]
]

ROWS = len(matrix)
COLS = len(matrix[0])

def isValid(x, y):
    return (x >= 0 and x < COLS and y >=0 and y < ROWS)

visited = []

def dfs(current):
    if  