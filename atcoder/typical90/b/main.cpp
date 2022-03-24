#include <iostream>
#include <math.h>

using namespace std;

bool valid(string s) {
	int count = 0;
	for (int i = 0; i < s.size(); i++) {
		if (s[i] == '(') {
			count++;
		} else {
			count--;
		}
		if (count < 0) {
			return false;
		}
	}

	return count == 0;
}

int main() {
	int N;
	cin >> N;

	if ((N % 2) != 0) {
		return 0;
	}

	for (int i = 0; i < pow(2, N); i++) {
		string s = "";
		for (int j = N - 1; j >= 0; j--) {
			if ((i & (1 << j)) == 0) {
				s += "(";
			} else {
				s += ")";
			}
		}

		if (valid(s)) {
			cout << s << "\n";
		}
	}

	return 0;
} 
