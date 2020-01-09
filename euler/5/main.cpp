#include <iostream>

 long long gcd(long long a, long long b) {
  if (a == 0)
    return b;
  return gcd(b % a, a);
 }

long long lcm(long a, long b) { return (a * b) / gcd(a, b); }

int main() {
  // calculate LCM of 1...20
  long long curr = lcm(20, 19);
  for (int i = 18; i > 1; i--) {
    curr = lcm(curr, i);
  }
  std::cout << curr << std::endl;
  return 0;
}

