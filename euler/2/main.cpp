// https://projecteuler.net/problem=2
#include <iostream>

int main() {
  int max = 4000000;
  int sum = 2; // because loop starts from 3
  int prevprev = 1;
  int prev = 2;

  for (;;) {
    int now = prev + prevprev;
    if (now > max) {
      break;
    }
    if (now % 2 == 0) {
      sum += now;
    }
    prevprev = prev;
    prev = now;
  }

  std::cout << sum << std::endl;
  return 0;
}
