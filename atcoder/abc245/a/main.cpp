#include <iostream>

using namespace std;

int main() {
	int a;
	int b;
	int c;
	int d;
	cin >> a >> b >> c >> d;

	if (a < c) {
		cout << "Takahashi" << endl;
	} else if (a > c) {
		cout << "Aoki" << endl;
	} else {
		if (b > d) {
			cout << "Aoki" << endl;
		} else {
			cout << "Takahashi" << endl;
		}
	}
	
	return 0;
}
