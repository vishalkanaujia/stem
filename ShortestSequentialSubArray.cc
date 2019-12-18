#include <iostream>
#include <vector>
#include <unordered_set>
#include <unordered_map>
#include <limits>

using namespace std;

struct Subarray
{
    int start, end;
};

Subarray FindSmallestSequenceSubarray(const vector<string> &paragraph, const vector<string> &keywords)
{
    unordered_map<string, int> keyword_to_idx;

    for (int i = 0; i < keywords.size(); ++i)
    {
        // emplace returns a pair <iterator, true> for success
        // if key already exists, <iterator_to_existing, false>
        keyword_to_idx.emplace(keywords[i], i);
    }

    vector<int> latest_occurrences(keywords.size(), -1);

    vector<int> shortest_subarray_length(keywords.size(), numeric_limits<int>::max());

    Subarray result = Subarray{-1, -1};
    int shortest = numeric_limits<int>::max();

    for (int i = 0; i < paragraph.size(); ++i)
    {
        if (keyword_to_idx.count(paragraph[i]))
        {
            auto it = keyword_to_idx.find(paragraph[i]);
            int keyword_idx = it->second;
            cout << paragraph[i] << " " << keyword_idx << endl;
            if (keyword_idx == 0)
            { // the first keyword
                shortest_subarray_length[keyword_idx] = 1;
            } else if (shortest_subarray_length[keyword_idx - 1] != numeric_limits<int>::max())
            {
                int distance_to_previous_keyword = i - latest_occurrences[keyword_idx - 1];
                shortest_subarray_length[keyword_idx] = distance_to_previous_keyword +
                                                        shortest_subarray_length[keyword_idx - 1];
            }

            latest_occurrences[keyword_idx] = i;

            if (keyword_idx == keywords.size() - 1)
            {
                if (shortest_subarray_length.back() < shortest)
                {
                    shortest = shortest_subarray_length.back();
                    result = {i - shortest_subarray_length.back() + 1, i};
                }
            }
        }
    }
    return result;
}


int main()
{
    vector<string> paragraph;

    string str[] = {"cat", "is", "cute", "animal", "among", "all", "animals", "kind"};
    for (string s : str)
    {
        paragraph.push_back(s);
    }

    vector<string> keywords;
    string query[] = {"cute", "animal", "all"};
    for (auto it : query)
    {
        keywords.push_back(it);
    }

    Subarray result = FindSmallestSequenceSubarray(paragraph, keywords);
    cout << result.start << " " << result.end << endl;

    return 0;
}