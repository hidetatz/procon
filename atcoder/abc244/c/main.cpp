#include <bits/stdc++.h>

using namespace std;

// using ll = long long;

int main() {
  int n;
  cin >> n;

  vector<int> as(n * 2 + 1);
  for (int i = 0; i < n * 2 + 1; i++) as[i] = 0;

  for (int i = 0; i < 2 * n + 1; i++) {
    for (int j = 0; j < n * 2 + 1; j++) {
      if (as[j] == 0) {
        as[j] = 1;
        cout << j + 1 << endl;
        break;
      }
    }

    int n;
    cin >> n;

    if (n == 0) {
      break;
    }
    as[n - 1] = 1;
  }
  return 0;
}
