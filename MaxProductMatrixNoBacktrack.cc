#include <iostream>

using namespace std;

void printMatrix(int **matrix, int r, int c) {
    for (int i = 0; i < r; i++) {
        for (int j = 0; j < c; j++) 
            cout << matrix[i][j] << " ";
       cout << endl;
    }
}

int matrixProduct(int matrix[][3], int r, int c) {
    // check for zero sized matrix
    if (r == 0 && c == 0) {
        return 0;
    }

    int **maxCache = new int*[r];
    for (int i =0 ; i < r; i++) {
        maxCache[i] = new int[c];
    }

    int temp1=0, temp2=0;

    for (int i = 0; i < r; i++) {
        for(int j = 0; j < c; j++) {
            if (i == 0 && j == 0) {
                maxCache[i][j] = matrix[i][j];
                continue;
            }

            if (i > 0) {
                temp1 = maxCache[i-1][j]*matrix[i][j];
            }

            if (j > 0) {
                temp2 = maxCache[i][j-1] * matrix[i][j];
            }

            maxCache[i][j] = max(temp1, temp2);
            temp1 = temp2 = -1;
        }
    }

    // print matrix
    printMatrix(maxCache, r, c);

    return maxCache[r-1][c-1];
}

int main() {
    int matrix[][3] = {{1,2,3},{4,5,6},{7,8,9}};
    cout << matrixProduct(matrix, 3,3) << endl;
}