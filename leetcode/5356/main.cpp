// https://leetcode.com/problems/lexicographical-numbers/
// Medium

#include <iostream>
#include <map>
#include <vector>

using namespace std;

void print(vector<int> const& input) {
    for (int i = 0; i < (int)input.size(); i++) {
        cout << input.at(i) << ' ';
    }
}

void printMap(map<int, int> const& input) {
    for (auto const& [k, v] : input) printf("(%d,%d)\n", k, v);
}

class Solution {
   public:
    map<int, int> minInRow;
    map<int, int> maxInCol;

    void createMap(vector<vector<int>>& matrix) {
        for (int i = 0; i < (int)matrix.size(); i++) {
            for (int j = 0; j < (int)matrix[i].size(); j++) {
                if (minInRow.count(i) == 0 || matrix[i][j] < minInRow[i]) {
                    minInRow[i] = matrix[i][j];
                }

                if (maxInCol.count(j) == 0 || matrix[i][j] > maxInCol[j]) {
                    maxInCol[j] = matrix[i][j];
                }
            }
        }
    }

    int getMinInRow(int i) { return minInRow[i]; }

    int getMaxInCol(int j) { return maxInCol[j]; }

    vector<int> luckyNumbers(vector<vector<int>>& matrix) {
        createMap(matrix);

        vector<int> ans;
        for (int i = 0; i < (int)matrix.size(); i++) {
            for (int j = 0; j < (int)matrix[i].size(); j++) {
                if (matrix[i][j] == getMinInRow(i)) {
                    if (matrix[i][j] == getMaxInCol(j)) {
                        ans.push_back(matrix[i][j]);
                    }
                }
            }
        }

        return ans;
    }
};

int main() {
    Solution sol;
    /* vector<vector<int>> input = {{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}; */
    /* vector<vector<int>> input = {{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17,
     * 12}}; */
    vector<vector<int>> input = {{36376, 85652, 21002, 4510},
                                 {68246, 64237, 42962, 9974},
                                 {32768, 97721, 47338, 5841},
                                 {55103, 18179, 79062, 46542}};
    auto ans = sol.luckyNumbers(input);
    print(ans);
    return 0;
}
