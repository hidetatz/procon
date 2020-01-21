#include <vector>

class Queue{
public:
    std::vector<int> data;
    int top;
    int size;
    int max;

    void Add(int val);
    int Remove();
    int Element();

    Queue() {
        top = 0;
        size = 0;
        max = 10;
        data.resize(10);
    }

    Queue(int m) {
        top = 0;
        size = 0;
        max = m;
        data.resize(m);
    }
};
