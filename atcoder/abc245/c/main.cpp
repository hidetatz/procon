#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
	int N;
	int K;
	cin >> N >> K;
	vector<int> A(N);
	vector<int> B(N);
	for (int i = 0; i < N; i++) {
		cin >> A[i];
	}
	for (int i = 0; i < N; i++) {
		cin >> B[i];
	}

	cout << "K: " << K << endl;

	cout << "A: ";
	for (int i: A) {
		cout << i << " ";
	}

	cout << "\n";

	cout << "B: ";
	for (int i: B) {
		cout << i << " ";
	}
	cout << "\n";

	for (int bit = 0; bit < (1 << N); bit++) {
		vector<int> vec(N);
		for (int j = 0; j < N; j++) {
			// pick from A if 0, else pick from B, 
			if (bit & (1 << j)) {
				vec[j] = B[j];
			} else {
				vec[j] = A[j];
			}
		}


		cout << "vec: ";
		for (int i: vec) {
			cout << i << " ";
		}
		cout << "\n";

		int cur = vec[0];
		bool no = false;

		for (int j = 1; j < N; j++) {
			int v = vec[j] - cur;
			if (v < 0) {
				v = -v;
			}

			if (v > K) {
				no = true;	
				break;
				// goto next sequence
			}
			cur = vec[j];
		}

		if (no) {
			continue;
		}

		cout << "Yes" << endl;
		return 0;
	}

	cout << "No" << endl;

	return 0;
}
