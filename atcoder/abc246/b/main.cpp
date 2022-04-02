#include <iostream>
#include <cmath>

using namespace std;

int main() {
	double a, b;
	cin >> a >> b;

	double len = sqrt(a * a + b * b);


	double x, y;
	// a:len = x:1
	x = a / len;
	// b:len = y:1
	y = b / len;
	cout << a << " " << b << endl;
}
