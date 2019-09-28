#include <bits/stdc++.h>

using namespace std;

void setRow(bool matrix[3][3], int index)
{
    cout << "Setting row number " << index << endl;
    for (auto i = 1; i < 3; i++)
        matrix[index][i] = true;
}

void setColumn(bool matrix[3][3], int index)
{
    cout << "Setting column number " << index << endl;
    for (auto i = 1; i < 3; i++)
        matrix[i][index] = true;
}

void printMatrix(bool matrix[3][3])
{
    for (auto i = 0; i < 3; i++)
    {
        for (auto j = 0; j < 3; j++)
            cout << matrix[i][j] << " ";
        cout << endl;
    }
    cout << endl;
}

// in-place approach
void zeroMatrix(bool matrix[3][3])
{
    printMatrix(matrix);

    bool firstRow, firstColumn;

    // The idea is to use the first row and the first column as a place
    // to hold the positions of 1 (true) in the matrix.

    // step 1: get the status of first row & first column
    for (auto i = 0; i < 3; i++)
    {
        if (matrix[0][i])
        {
            firstRow = true;
        }
        if (matrix[i][0])
        {
            firstColumn = true;
        }
    }

    for (auto i = 1; i < 3; i++)
    { // skip the first row & column
        for (auto j = 1; j < 3; j++)
        {
            if (matrix[i][j])
            {
                matrix[0][i] = true;
                matrix[j][0] = true;
            }
        }
    }

    for (auto i = 0; i < 3; i++) // print the first row & column
        cout << matrix[0][i] << " ";

    cout << endl;
    for (auto i = 0; i < 3; i++)
        cout << matrix[i][0] << " ";
    cout << endl;

    for (auto i = 1; i < 3; i++)
    { // modify the matrix
        if (matrix[0][i])
        {
            // set the ith row
            setRow(matrix, i);
        }
    }
    cout << "rows set complete" << endl;
    printMatrix(matrix);

    for (auto i = 1; i < 3; i++)
    { // modify the matrix
        if (matrix[i][0])
        {
            setColumn(matrix, i);
        }
    }
    cout << endl;
    printMatrix(matrix);
}

int main()
{
    bool matrix[][3] = {false, false, false, false, true, true, false, false, false};
    zeroMatrix(matrix);
    return 0;
}