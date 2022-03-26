#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
	int N;
	cin >> N;
	vector<int> vec(N);
	for (int i = 0; i < N; i++) {
		cin >> vec[i];
	}

	sort(vec.begin(), vec.end());

	if (vec[0] != 0) {
		cout << 0 << endl;
		return 0;
	}

	int cur = 0;

	for (int i = 0; i < N; i++) {
		if (vec[i] == cur) {
			continue;
		}

		if (vec[i] == cur + 1) {
			cur++;
			continue;
		}

		cout << cur + 1 << endl;
		return 0;
	}
	
	cout << cur + 1 << endl;
	return 0;
}
