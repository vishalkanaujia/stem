#include <bits/stdc++.h> 
#include <iostream>       // std::cerr
#include <stdexcept>      // std::out_of_range
#include <vector>  

using namespace std;

int findSequence(vector<int> v) {
    unordered_set<int> hashSet;

    for(auto i = v.begin(); i != v.end(); i++) {
        hashSet.insert(*i);
    }

    int length = 0;
    int max = 0;

    for(auto i = v.begin(); i != v.end(); i++) {
        if (hashSet.find(*i - 1) != hashSet.end()) {
            cout << "*i=" << *i << endl;
            continue;
        }
        int start = *i;
        cout << "start=" << start << endl;
        while (hashSet.find(start++) != hashSet.end()) {
            length++;
        }
        cout << "length=" << length << endl;
        if (length > max) {
            max = length;
            length = 0;
        }
    }
    return max;
}

int main() {
    vector<int> v = {1,4,5,6,3};
    cout << findSequence(v) << endl;
    return 0;
}