#include <iostream>
using namespace std;

int findMaxNonAdjacentSum(int *arr, int idx)
{
    if (idx == 0)
        return arr[0];
    if (idx < 0)
        return 0;
    int selected = findMaxNonAdjacentSum(arr, idx - 2) + arr[idx];
    int notSelected = findMaxNonAdjacentSum(arr, idx - 1);
    return max(notSelected, selected);       
}

// Optimized sum
int findMaxSumNonAdjOptimized(int *arr, int idx)
{
    int len = idx;
    int prev2 = arr[0];
    int prev1 = max(arr[0], arr[1]);

    if (idx > 2)
    {
        for (int i = 2; i <= len; i++)
        {
            int tmp = prev1;
            prev1 = max(arr[i] + prev2, prev1);
            prev2 = tmp;
        }
    }
    return prev1;
}

int main() 
{
    int arr[] = {1,22,11,5,6,67};
    cout << findMaxNonAdjacentSum(arr, sizeof(arr)/sizeof(int) - 1) << endl;
    cout << findMaxSumNonAdjOptimized(arr, sizeof(arr)/sizeof(int) - 1) << endl;
}