#include <iostream>
#include <queue>

using namespace std;

struct Iterator {
    bool operator >(const Iterator& that) {
        return *current > *that.current;
    }

    vector<int>::const_iterator current;
    vector<int>::const_iterator end;
};

vector<int> MergeSortedArray(const vector<vector<int>>& sorted_arrays) {
    priority_queue<Iterator, vector<Iterator>, greater<>> min_heap;

    for (const vector<int>& sorted_array: sorted_arrays) {
        if (!sorted_array.empty()) {
            min_heap.emplace(Iterator{sorted_array.cbegin(), sorted_array.cend()});
        }
    }

    vector<int> result;

    while(!min_heap.empty()) {
        auto smallest_array = min_heap.top();
        min_heap.pop();
        if (smallest_array.current != smallest_array.end) {
            result.emplace_back(*smallest_array.current);
            min_heap.emplace(Iterator{next(smallest_array.current), smallest_array.end});
        }
    }
    return result;
}