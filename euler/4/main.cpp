#include <iostream>

// http://www.trytoprogram.com/cpp-examples/cplusplus-program-to-check-palindrome-number/
bool isPalindrome(int num) {
  int temp = num;
  int reverse = 0;

  while (temp != 0) {
    int reminder = temp % 10;
    reverse = reverse * 10 + reminder;
    temp = temp / 10;
  }

  return num == reverse;
}

int main() {
  int ans = 0;
  for (int p = 999; p > 99; p--) {
    for (int q = 999; q > 99; q--) {
      int now = p * q;
      if (isPalindrome(now) && ans < now) {
        ans = now;
        break;
      }
    }
  }
  std::cout << ans << std::endl;
}
