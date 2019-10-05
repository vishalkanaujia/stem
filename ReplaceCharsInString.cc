#include<iostream>
#include<algorithm>

using namespace std;

char* ReplaceChars(int size, char input[])
{
    int write_idx = 0, b_count = 0;
    
    for (int i = 0; i < size; i++) {
        if (input[i] != 'a') {
            input[write_idx] = input[i];
            write_idx++;
        }

        if (input[i] == 'b') {
            ++b_count;
        }
    }

    // add 'dd' for each b
    int count = write_idx - 1;

    for (int i = write_idx + b_count - 1; i >= 0; i--)
    {
        if (input[count] == 'b') {
            input[i] = 'd';
            input[--i] = 'd';
        } else {
            input[i] = input[count];
        }
        count--;
    }

    //input[write_idx] = '\0';
    return input;
}

int main()
{
    char s[] = "baab";
    cout << ReplaceChars(sizeof(s), s) << endl;
}