#include <algorithm>
#include <iostream>
#include <vector>

typedef long long ll;

using namespace std;

int main() {
  ll n, k, x;
  cin >> n >> k >> x;

  vector<ll> as(n);

  for (int i = 0; i < n; i++) {
    cin >> as.at(i);
  }

  ll ans = 0;
  for (int i = 0; i < n; i++) {
    ans += as[i];
  }

  ll m = 0;
  for (int i = 0; i < n; i++) {
    m += as[i] / x;
    as[i] %= x;
  }
  if (k < m) {
    m = k;
  }
  k -= m;
  ans -= m * x;

  if (k <= 0) {
    cout << ans << endl;
    return 0;
  }

  sort(as.rbegin(), as.rend());
  for (int i = 0; i < n; i++) {
    if (k == 0) {
      break;
    }
    ans -= as[i];
    k--;
  }

  cout << ans << endl;
}
