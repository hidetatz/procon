#include "queue.hpp"

void Queue::Add(int val) {
  if (this->size == this->max) {
    return;
  }

  this->data[(this->size + this->top) % this->max] = val;
  ++this->size;
}

int Queue::Remove() {
  int ret = this->data[this->top];
  this->top = this->top + 1 % this->max;
  --this->size;
  return ret;
}

int Queue::Element() { return this->data[this->top]; }
