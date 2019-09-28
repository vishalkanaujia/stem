// anagram.cpp
#include "iostream"
#include <map>
#include <cassert>

using namespace std;

string s = "aabbcc";
string w = "ababcc";

class StringChecker {
    private:
    string str;

    public:
    StringChecker() {

    }

    StringChecker(string s) {
        str = s;
    }

    bool isAnagram(string input);
};

bool StringChecker::isAnagram(string input)
{
    char c;
    std::map<std::char, int>

    for (int i = 0; i < input.size(); i++) {
        c = input[i];
        m[c] += 1;
    }

    std::map<char, int>::iterator i;
     
    return true;    
}

int main()
{
    StringChecker *s = new StringChecker("abccba");
    s->isAnagram("aabbcc");
    return 0;
}