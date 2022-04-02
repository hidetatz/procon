#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
        int n, k, x;
        cin >> n >> k >> x;

        vector<int> as(n);

        for (int i = 0; i < n; i++) {
                cin >> as.at(i);
        }

        while (k != 0) {
                sort(as.rbegin(), as.rend());

                bool used = false;
                for (int i = 0; i < n; i++) {
                        if (k == 0) {
                                break;
                        }
                        int price = as[i];
                        if (price - x < 0) {
                                continue;
                        }
                        as[i] = price - x;
                        k--;
                        used = true;
                }
                if (!used) {
                        break;
                }
        }

        while (k != 0) {
                sort(as.rbegin(), as.rend());

                bool used = false;
                for (int i = 0; i < n; i++) {
                        if (k == 0) {
                                break;
                        }
                        int price = as[i];
                        as[i] = price - x;
                        k--;
                        used = true;
                }
                if (!used) {
                        break;
                }
        }

        int result = 0;
        for (int i = 0; i < n; i++) {
       		if (as[i] < 0) {
			continue;
                }
                result += as[i];
        }

        cout << result << endl;
}
