#include<iostream>
#include <vector>

using namespace std;

vector<int> findDuplicates(int arr[], int len) {
    vector<int> result;

    for(int i = 0; i < len; i++) {
        int index = abs(arr[i]) - 1;
        if (arr[index] < 0) {
            result.push_back(arr[i]);
        } else {
            arr[index] = -arr[index];
        }
    }
    return result;
}

int main() {
    vector<int> v;
    int arr[5] = {1,2,2,3,1};
    
    v = findDuplicates(arr, 5);

    for (auto i = v.begin(); i != v.end(); i++) {
        cout << *i << " ";
    }
    cout << endl;
    return 0;
}