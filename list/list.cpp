#include "list.hpp"

void List::Add(std::string val) {
    Node* node = new Node;
    node->value_ = val;
    node->next_ = NULL;

    int size = this->Size();
    if (size == 0) {
        this->first_ = node;
        return;
    }

    Node* last = this->GetNode(size - 1);
    last->next_ = node;
}

void List::Add(int index, std::string val) {
    if (index < 0) {
        return;
    }

    Node* node = new Node;
    node->value_ = val;
    node->next_ = NULL;

    int size = this->Size();
    if (size <= index) {
        // append to Last
        this->Add(val);
        return;
    }

    Node* after = this->GetNode(index);
    node->next_ = after;

    if (index == 0) {
        this->first_ = node;
    }

    if (index > 0) {
        Node* before = this->GetNode(index - 1);
        before->next_ = node;
    }

    return;
}

bool List::Contains(std::string val) {
    if (this->first_ == NULL) {
        return false;
    }

    int size = this->Size();
    Node* now = this->first_;
    for (int i = 0; i < size; i++) {
        if (now->value_ == val) {
            return true;
        }
        now = now->next_;
    }
    return false;
}

std::string List::Get(int index) {
    Node* node = this->GetNode(index);
    return node->value_;
}

bool List::IsEmpty() { return this->first_ == NULL; }

void List::Remove(int index) {
    if (index < 0) {
        return;
    }

    int size = this->Size();
    if (size <= index) {
        return;
    }

    if (index == 0 && size == 1) {
        this->first_ = NULL;
        return;
    }

    Node* next = this->GetNode(index + 1);

    if (index == 0) {
        this->first_ = next;
    }

    if (index > 0) {
        Node* previous = this->GetNode(index - 1);
        previous->next_ = next;
    }
}

int List::Size() {
    if (this->first_ == NULL) {
        return 0;
    }

    Node* now = this->first_;
    int size = 1;
    for (;;) {
        now = now->next_;
        if (now == NULL) {
            break;
        }
        size++;
    }
    return size;
}

void List::Set(int index, std::string val) {
    if (index < 0) {
        return;
    }

    int size = this->Size();
    if (size <= index) {
        return;
    }

    Node* target = this->GetNode(index);
    target->value_ = val;
}

Node* List::GetNode(int index) {
    if (index == 0) {
        return this->first_;
    }

    Node* target = this->first_;
    for (int i = 0; i < index; i++) {
        target = target->next_;
    }
    return target;
}
