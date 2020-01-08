// https://projecteuler.net/problem=3

#include <iostream>

int main() {
  long long n = 600851475143LL;

  for (long long i = 2LL; i < n; i++) {
    if (n % i == 0) {
      n /= i;
    }
  }

  std::cout << n << std::endl;
}
