#include <iostream>

int main() {
    int sum_of_square = 0;
    int square_of_sum = 0;
    for (int i = 1; i <= 100; i++) {
        sum_of_square += i * i;
        square_of_sum += i;
    }
    square_of_sum = square_of_sum * square_of_sum;

    std::cout << square_of_sum - sum_of_square << std::endl;
    return 0;
}

