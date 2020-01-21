#include "queue.cpp"
#include <iostream>

int main() {
  Queue q(5);

  q.Add(1);
  q.Add(2);
  q.Add(3);
  q.Add(4);
  q.Add(5);
  q.Add(6); // won't be added

  std::cout << q.Element() << std::endl;
  std::cout << q.Element() << std::endl;
  std::cout << q.Remove() << std::endl;
  std::cout << q.Remove() << std::endl;
  std::cout << q.Remove() << std::endl;
  std::cout << q.Remove() << std::endl;
  std::cout << q.Remove() << std::endl;
  std::cout << q.Remove() << std::endl;
}
