#include <iostream>

using namespace std;

int main() {
	int x1, y1, x2, y2, x3, y3;
	cin >> x1 >> y1 >> x2 >> y2 >> x3 >> y3;

	int x, y;

	if (x1 == x2) {
		x = x3;
	} else if (x1 == x3) {
		x = x2;
	} else {
		x = x1;
	}

	if (y1 == y2) {
		y = y3;
	} else if (y1 == y3) {
		y = y2;
	} else {
		y = y1;
	}

	cout << x << " " << y << endl;
}
