#include<iostream>

using namespace std;

void MergeArrays(int *bigger, int big_size, int total, int *smaller, int small_size)
{
    int final_index = total - 1, small_index = small_size - 1, big_index = big_size - 1;
    for (int i = final_index; i >= 0; i--)
    {
        while (small_index >= 0 && big_index >= 0) {
            bigger[i--] = (smaller[small_index] > bigger[big_index]) ? smaller[small_index--] : bigger[big_index--];
            cout << "bigger:"<< i << "[]=>" << bigger[i] << endl;
        }

        cout << "small_index=" << small_index << " big_index=" << big_index << endl;

        if (small_index >= 0)
            bigger[i] = smaller[small_index--];

        if (big_index >= 0)
            bigger[i] = bigger[big_index--];
    }

    for (int i = 0; i < final_index; i++)
        cout << bigger[i] << " ";
    cout << endl;
}

int main()
{
    int *arr1 = (int*)malloc(sizeof(int)*5);
    int *arr2 = (int*)malloc(sizeof(int)*3);
    for (int i = 0; i < 3; i++)
        arr1[i] = i;

    for (int i = 0; i < 2; i++)
        arr2[i] = -(i*i+7);

    MergeArrays(arr1, 3, 5, arr2, 2);
}