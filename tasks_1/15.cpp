#include <iostream>
#include <string>

using namespace std;

int main()
{
    string vv, rle;
    std::cin >> vv;
    int colb;
    for (int i = 0; i < vv.length(); i++) {
        colb = 1;
        if (vv[i] == vv[i + 1]) {
            while (vv[i] == vv[i + 1]) {
                colb += 1;
                i++;
            }
            rle += vv[i] + to_string(colb);
        }
        else {
            rle += vv[i];
        }
    }
    cout << rle;
    return 0;
}