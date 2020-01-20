#include "node.hpp"

template <class T>
class List {
public:
    Node<T>* first_;

    void Add(T val);
    void Add(int index, T val);
    bool Contains(T val);
    T Get(int index);
    bool IsEmpty();
    void Remove(int index);
    void Set(int index, T val);
    int Size();

    List() {
        first_ = nullptr;
    }

private:
    Node<T>* GetNode(int index);
};
