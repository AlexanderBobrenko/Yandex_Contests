#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;


int main() {
    vector <string> words;
    while (true) {
        string word;
        cin >> word;
        if (word == "end") {
            break;
        }
        words.push_back(word);
    }
    vector <string> alf;
    for (int i = 0; i < words.size(); i++) {
        for (int j = 0; j < words.size(); j++) {
            if (words[i] == words[j] && i != j) {
                alf.push_back(words[i]);
            }
        }
    }
    sort(alf.begin(), alf.end());
    auto dubl = unique(alf.begin(), alf.end());
    alf.erase(dubl, alf.end());
    for (int k = 0; k < alf.size(); k++) {
        cout << alf[k] << " ";
    }
}