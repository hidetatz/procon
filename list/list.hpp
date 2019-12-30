#include "node.hpp"

class List {
public:
    Node* first_;

    void Add(std::string val);
    void Add(int index, std::string val);
    bool Contains(std::string val);
    std::string Get(int index);
    bool IsEmpty();
    void Remove(int index);
    void Set(int index, std::string val);
    int Size();

private:
    Node* GetNode(int index);
};
