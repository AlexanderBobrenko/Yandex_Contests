#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

bool anagram(long num1, long num2) {
    string str1 = to_string(num1);
    string str2 = to_string(num2);
    sort(str1.begin(), str1.end());
    sort(str2.begin(), str2.end());

    return str1 == str2;
}

int main() {
    long n1, n2;
    cin >> n1 >> n2;
    anagram(n1, n2) ? cout << "YES" : cout << "NO";
}