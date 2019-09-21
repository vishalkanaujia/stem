#include<iostream>

using namespace std;

bool MatchQuestion(const char * pattern, const char *input);
bool MatchStar(const char *pattern, const char *input);

bool matchOne(char pattern, char input)
{
  cout << "matchOne" << "patt=" << pattern << "input=" << input << endl;
 if (pattern == '\0')
     return true;

  if (input == '\0')
      return false;

  return pattern == input;
}

bool MatchSameLength(const char *pattern, const char* input) 
{
    if (*pattern == '\0')
        return true;

    return matchOne(pattern[0], input[0]) && MatchSameLength(pattern+1, input+1);
}

bool MatchAnywhere(const char *pattern, const char *input)
{
    //cout << "patt=" << *pattern << "input=" << *input << endl;
    if (*pattern == '\0')
        return true;

    if (*pattern == '$' && *input == '\0')
        return true;

    if (pattern[1] == '?') {
        return MatchQuestion(pattern, input);
    }

    if (pattern[1] == '*') {
        return MatchStar(pattern, input);
    }

    return matchOne(pattern[0], input[0]) && MatchAnywhere(pattern+1, input+1);
}

bool MatchQuestion(const char * pattern, const char *input)
{
    return (
        MatchAnywhere(pattern+2, input) ||
            (matchOne(*pattern, *input) && MatchAnywhere(pattern+2, input))
            );
}

bool MatchStar(const char *pattern, const char *input)
{
 return (
     MatchAnywhere(pattern+2, input) ||
     (matchOne(*pattern, *input) && MatchAnywhere(pattern, input+1))
 );
}

bool search(const char *pattern, const char *input )
{
    if (*pattern == '^') {
        return MatchAnywhere(pattern+1, input);
    }

    while (*input) {
        if (MatchAnywhere(pattern, input))
            return true;
        ++input;
        }
    return false;    
}

int main() {
    //cout << MatchSameLength("aab", "aab") << endl;
    //cout << MatchAnywhere("^b$", "b") << endl;

    cout << search("ax*b", "abc") << endl;
    return 1;
}