#include<iostream>

using namespace std;

void permutationsHelper(string s, string chosen)
{
    cout << "s= " << s << " chosen= " << chosen << endl; 
    if (s.empty()) {
        cout << chosen << endl;
        return;
    }

    for (int i = 0; i < s.length(); i++)
    {
        char c = s[i];
        chosen += c;
        s.erase(i, 1);
        cout << "choosing= " << c << " i= " << i << " s= " << s <<endl;
        permutationsHelper(s, chosen);

        cout << "choice= " << c << " i= " << i << " s= " << s <<endl;
        s.insert(i, 1, c);
        chosen.erase(chosen.length() - 1, 1);
    }
}

void Permutations(string s)
{
    string chosen = "";
    permutationsHelper(s, chosen);
}

int main()
{
    Permutations("xy");
}