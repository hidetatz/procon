#include <vector>

template <class T>
class Stack{
public:
    std::vector<T> data;
    int top;

    bool IsEmpty();
    void Push(T val);
    T Pop();
    T Peek();
    int Search(T val);

    Stack() {
        top = 0;
    }
};
