// https://leetcode.com/problems/lexicographical-numbers/
// Medium

#include <vector>
#include <iostream>

using namespace std; 

void print(std::vector<int> const &input) {
    for (int i = 0; i < input.size(); i++) {
        cout << input.at(i) << ' ';
    }
}

class Solution {
    public:
        vector<int> lexicalOrder(int n) {
            vector<int> res;
            for (int i = 1; i < 10; i++) {
                dfs(res, i, n);
            }
            return res;
        }

    private:
        void dfs(vector<int>& res, int now, int n) {
            if (now > n) {
                return;
            }
            res.push_back(now);
            for (int i = now * 10; i < now * 10 + 10; i++) {
                dfs(res, i, n);
            }
        }
};

int main() {
    Solution sol;
    auto ans = sol.lexicalOrder(13);
    print(ans);
    return 0;
}
