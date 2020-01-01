#include "list.hpp"

template <typename T>
void List<T>::Add(T val) {
    /* void List::Add(std::string val) { */
    auto node = new Node<T>;
    node->value_ = val;
    node->next_ = nullptr;

    int size = this->Size();
    if (size == 0) {
        this->first_ = node;
        return;
    }

    Node<T>* last = this->GetNode(size - 1);
    last->next_ = node;
}

template <typename T>
void List<T>::Add(int index, T val) {
    /* void List::Add(int index, std::string val) { */
    if (index < 0) {
        return;
    }

    auto node = new Node<T>;
    node->value_ = val;
    node->next_ = nullptr;

    int size = this->Size();
    if (size <= index) {
        // append to Last
        this->Add(val);
        return;
    }

    Node<T>* after = this->GetNode(index);
    node->next_ = after;

    if (index == 0) {
        this->first_ = node;
    }

    if (index > 0) {
        Node<T>* before = this->GetNode(index - 1);
        before->next_ = node;
    }

    return;
}

template <typename T>
bool List<T>::Contains(T val) {
    /* bool List::Contains(std::string val) { */
    if (this->first_ == nullptr) {
        return false;
    }

    int size = this->Size();
    Node<T>* now = this->first_;
    for (int i = 0; i < size; i++) {
        if (now->value_ == val) {
            return true;
        }
        now = now->next_;
    }
    return false;
}

template <typename T>
T List<T>::Get(int index) {
    /* std::string List::Get(int index) { */
    Node<T>* node = this->GetNode(index);
    return node->value_;
}

template <typename T>
bool List<T>::IsEmpty() {
    return this->first_ == nullptr;
}
/* bool List::IsEmpty() { return this->first_ == nullptr; } */

template <typename T>
void List<T>::Remove(int index) {
    if (index < 0) {
        return;
    }

    int size = this->Size();
    if (size <= index) {
        return;
    }

    if (index == 0 && size == 1) {
        this->first_ = nullptr;
        return;
    }

    Node<T>* next = this->GetNode(index + 1);

    if (index == 0) {
        this->first_ = next;
    }

    if (index > 0) {
        Node<T>* previous = this->GetNode(index - 1);
        previous->next_ = next;
    }
}

template <typename T>
int List<T>::Size() {
    if (this->first_ == nullptr) {
        return 0;
    }

    Node<T>* now = this->first_;
    int size = 1;
    for (;;) {
        now = now->next_;
        if (now == nullptr) {
            break;
        }
        size++;
    }
    return size;
}

template <typename T>
void List<T>::Set(int index, T val) {
    if (index < 0) {
        return;
    }

    int size = this->Size();
    if (size <= index) {
        return;
    }

    Node<T>* target = this->GetNode(index);
    target->value_ = val;
}

template <typename T>
Node<T>* List<T>::GetNode(int index) {
    if (index == 0) {
        return this->first_;
    }

    Node<T>* target = this->first_;
    for (int i = 0; i < index; i++) {
        target = target->next_;
    }
    return target;
}
