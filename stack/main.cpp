#include "stack.cpp"
#include <iostream>
#include <vector>

int main() {
  Stack<int> s;
  std::cout << s.IsEmpty() << std::endl;

  s.Push(1);
  s.Push(2);
  s.Push(3);
  s.Push(4);
  s.Push(5);

  auto val = s.Peek();
  std::cout << val << std::endl;
  std::cout << s.IsEmpty() << std::endl;

  int pos = s.Search(3);
  std::cout << pos << std::endl;

  std::cout << s.Pop() << std::endl;
}
