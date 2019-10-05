#include <iostream>
#include <unordered_set>
#include <unordered_map>
#include <vector>

using namespace std;

typedef pair<int, int> indexes;

indexes FindMinimalSubarray(const vector<string> &paragraph,
                            const unordered_set<string> &keywords,
                            unordered_map<string, int> keywords_to_cover)
{

    // prepare the hashmap of search query strings
    for (auto s : keywords)
    {
        keywords_to_cover[s]++;
    }

    indexes result;
    result.first = result.second = -1;

    int to_cover = keywords_to_cover.size();

    // keep left & right index, increment right till we cover all search
    // strings. Then we start increasing left till we do not cover all
    // search strings.
    for (int left = 0, right = 0; right < paragraph.size(); ++right)
    {
        string current = paragraph[right];
        if (keywords.count(current) &&
            --keywords_to_cover[current] >= 0)
        {
            --to_cover;
        }

        while (to_cover == 0)
        {
            // adjust the result indexes
            if ((result.first == -1 && result.second == -1) ||
                (result.second - result.first > right - left))
            {
                result.first = left;
                result.second = right;
            }

            // move the left index till we do not cover all query strings.
            if (left <= right) {
                if (keywords.count(paragraph[left]) &&
                (keywords_to_cover[paragraph[left]] == 0)) {
                   ++keywords_to_cover[paragraph[left]];
                   ++to_cover;
                }
                ++left;
            }
        }
    }
    return result;
}

int main()
{
    vector<string> paragraph;

    string str[] = {"cat", "is", "cute", "animal", "among", "all", "animal", "kind"};
    for (string s : str)
    {
        paragraph.push_back(s);
    }

    unordered_set<string> keywords;
    string query[] = {"cute", "animal", "all"};
    for (auto it: query) {
        keywords.emplace(it);
    }

    unordered_map<string, int> map;

    auto result = FindMinimalSubarray(paragraph, keywords, map);
    cout << result.first << " " << result.second << endl;

    return 0;
}