#include <iostream>
#include <algorithm>

int main() {
	std::string s;
	std::cin >> s;
	size_t cnt = std::count(s.begin(), s.end(), '1');
	std::cout << cnt << std::endl;
	return 0;
}
