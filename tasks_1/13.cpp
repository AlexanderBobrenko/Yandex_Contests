#include <iostream>
#include <math.h>

using namespace std;

int main()
{

    int n, b = 1, c = 0, d = 1, e = 1, f = 2;
    cin >> n;

    while (b <= n) {
        cout << b << " ";
        b += 1;
        c += 1;
        if (c == d) {
            cout << "\n";
            d += e;
            c = 0;
        }
        if (d == f) {
            e -= 2;
            f += 1;
        }
        if (d <= 1) {
            e = 1;
        }
    }
    return 0;
}