#include <bits/stdc++.h>

using namespace std;

// using ll = long long;

int main() {
  int n;
  cin >> n;
  string s;
  cin >> s;

  string dir = "east";
  int x = 0;
  int y = 0;
  for (int i = 0; i < n; i++) {
    char v = s[i];
    cerr << v << " " << dir << endl;

    if (v == 'S') {
      if (dir == "east") x++;
      if (dir == "west") x--;
      if (dir == "north") y++;
      if (dir == "south") y--;
      continue;
    }

    if (dir == "east")
      dir = "south";
    else if (dir == "west")
      dir = "north";
    else if (dir == "north")
      dir = "east";
    else if (dir == "south")
      dir = "west";
  }

  cout << x << " " << y << endl;

  return 0;
}
