#include <iostream>
#include <map>
#include <vector>

using namespace std;

void printVector(vector<int> const& input) {
    for (int i = 0; i < (int)input.size(); i++) {
        cout << input.at(i) << ' ';
        cout << endl;
    }
}

class CustomStack {
   public:
    vector<int> st{};
    int maxSize;

    void print() { printVector(st); }

    CustomStack(int maxSize) { this->maxSize = maxSize; }

    void push(int x) {
        if ((int)st.size() == maxSize) {
            return;
        }

        st.push_back(x);
    }

    int pop() {
        if ((int)st.size() == 0) {
            std::cout << -1 << std::endl;
            return -1;
        }

        int ret = st[(int)st.size() - 1];
        st.erase(st.begin() + (int)st.size() - 1);
        std::cout << ret << std::endl;
        return ret;
    }

    void increment(int k, int val) {
        if (k > (int)st.size()) {
            k = (int)st.size();
        }

        for (int i = 0; i < k; i++) {
            st[i] += val;
        }
    }
};

int main() {
    CustomStack customStack(3);  // Stack is Empty []
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.push(1);
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.push(2);
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.pop();
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.push(2);
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.push(3);
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.push(4);
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.increment(5, 100);
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.increment(2, 100);
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.pop();
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.pop();
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.pop();
    customStack.print();
    std::cout << "-----" << std::endl;

    customStack.pop();
    std::cout << "-----" << std::endl;
}
