#include "stack.hpp"
#include <vector>

template <typename T> bool Stack<T>::IsEmpty() {
  return this->data.size() == 0;
}

template <typename T> void Stack<T>::Push(T val) {
  this->data.push_back(val);
  ++this->top;
}

template <typename T> T Stack<T>::Pop() {
  --this->top;
  T ret = this->data.at(this->top);
  return ret;
}

template <typename T> T Stack<T>::Peek() {
  return this->data.at(this->top - 1);
}

template <typename T> int Stack<T>::Search(T val) {
  for (auto i = 0; i != this->data.size(); i++) {
    if (this->data[i] == val) {
      return i;
    }
  }

  return -1;
}

