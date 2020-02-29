#include <iostream>
#include <vector>

int main() {
    int want = 19690720;
    std::vector<int> input{
        1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,10,23,2,10,23,27,1,27,6,31,1,13,31,35,1,13,35,39,1,39,10,43,2,43,13,47,1,47,9,51,2,51,13,55,1,5,55,59,2,59,9,63,1,13,63,67,2,13,67,71,1,71,5,75,2,75,13,79,1,79,6,83,1,83,5,87,2,87,6,91,1,5,91,95,1,95,13,99,2,99,6,103,1,5,103,107,1,107,9,111,2,6,111,115,1,5,115,119,1,119,2,123,1,6,123,0,99,2,14,0,0
    };

    for (auto n = 0; n < 99; ++n) {
        for (auto m = 0; m < 99; ++m) {
            std::vector<int> copied{input};
            copied[1] = n;
            copied[2] = m;

            for (auto i = 0; i < 125; i += 4) {
                if (copied[i] == 99) {
                    break;
                }
                if (copied[i] == 1) {
                    copied[copied[i+3]] = copied[copied[i+1]] + copied[copied[i+2]];
                    continue;
                }
                if (copied[i] == 2) {
                    copied[copied[i+3]] = copied[copied[i+1]] * copied[copied[i+2]];
                    continue;
                }
            }
            if (copied[0] == want) {
                std::cout << 100 * n + m << std::endl;
                break;
            }
        }
    }
}
