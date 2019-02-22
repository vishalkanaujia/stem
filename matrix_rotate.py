matrix = [[1, 2, 3, 7],
          [4, 5, 6, 8],
          [4, 5, 6, 9],
          [7, 8, 9, 9]]


def cnext(r, c, degree):
    if r == c:
        if r == 0 and (c + 1) < degree:
            return (r, c + 1)
        if r == degree - 1:
            return (r, c - 1)

    if (r < c):
        if c + 1 < degree:
            return (r, c + 1)
        else:
            return (r + 1, c)

    if (r > c):
        if c == 0:
            return (r - 1, c)
        else:
            return (r, c - 1)


def rotateMatrix(matrix, length):
    if length == 0 or (len(matrix[0]) != length):
        return None

    degree = length
    layers = length / 2

    for i in range(layers + 1):
        if degree == 0 or degree == 1:
            break

        print("starting at degree = ", degree)
        print(i, i)
        next = cnext(i, i, degree)
        while True:
            print(next)
            next = cnext(next[0], next[1], degree)
            if next == (i, i):
                break

        degree = degree / 2


rotateMatrix(matrix, 4)
