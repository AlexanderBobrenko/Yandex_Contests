#include <iostream>
#include <math.h>

int main() {
	float res = sqrt(12) * (1 - (1 / (3.0 * 3.0)) + (1 / (5.0 * 3.0 * 3.0)) - (1 / (7.0 * 3.0 * 3.0 * 3.0 * 3.0)) + (1 / (9.0 * 3.0 * 3.0 * 3.0 * 3.0)) - (1 / (11.0 * 3.0 * 3.0 * 3.0 * 3.0 * 3.0)));
	std::cout << res << std::endl;

	return 0;
}
