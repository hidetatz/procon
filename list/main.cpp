#include <iostream>
#include <string>
#include "list.cpp"

bool testAdd1() {
    List* list = new List;
    if (!list->IsEmpty()) {
        return false;
    }

    list->Add("a");
    if (list->Get(0) != "a") {
        return false;
    }

    list->Add("b");
    if (list->Get(0) != "a") {
        return false;
    }
    if (list->Get(1) != "b") {
        return false;
    }

    return true;
}

bool testAdd2() {
    List* list = new List;
    if (!list->IsEmpty()) {
        return false;
    }
    list->Add("a");
    list->Add("b");
    if (list->Get(0) != "a") {
        return false;
    }
    if (list->Get(1) != "b") {
        return false;
    }

    list->Add(0, "c");
    if (list->Get(0) != "c") {
        return false;
    }
    if (list->Get(1) != "a") {
        return false;
    }
    if (list->Get(2) != "b") {
        return false;
    }

    list->Add(1, "d");
    if (list->Get(0) != "c") {
        return false;
    }
    if (list->Get(1) != "d") {
        return false;
    }
    if (list->Get(2) != "a") {
        return false;
    }
    if (list->Get(3) != "b") {
        return false;
    }

    list->Add(4, "e");
    if (list->Get(0) != "c") {
        return false;
    }
    if (list->Get(1) != "d") {
        return false;
    }
    if (list->Get(2) != "a") {
        return false;
    }
    if (list->Get(3) != "b") {
        return false;
    }
    if (list->Get(4) != "e") {
        return false;
    }

    return true;
}

bool testContains() {
    List* list = new List;
    list->Add("a");
    list->Add("b");

    if (!list->Contains("a")) {
        return false;
    }

    if (!list->Contains("b")) {
        return false;
    }

    if (list->Contains("c")) {
        return false;
    }

    return true;
}

bool testIsEmpty() {
    List* list = new List;
    if (!list->IsEmpty()) {
        return false;
    }

    list->Add("a");
    if (list->IsEmpty()) {
        return false;
    }

    return true;
}

bool testRemove() {
    List* list = new List;
    list->Add("a");
    list->Add("b");
    list->Add("c");
    if (list->Size() != 3) {
        return false;
    }

    list->Remove(0);
    if (list->Size() != 2) {
        return false;
    }
    if (list->Get(0) != "b") {
        return false;
    }
    if (list->Get(1) != "c") {
        return false;
    }

    list->Remove(2);
    if (list->Size() != 2) {
        return false;
    }
    if (list->Get(0) != "b") {
        return false;
    }
    if (list->Get(1) != "c") {
        return false;
    }

    list->Remove(-1);
    if (list->Size() != 2) {
        return false;
    }
    if (list->Get(0) != "b") {
        return false;
    }
    if (list->Get(1) != "c") {
        return false;
    }

    return true;
}

bool testSet() {
    List* list = new List;
    list->Add("a");
    list->Add("b");
    list->Add("c");
    if (list->Size() != 3) {
        return false;
    }

    list->Set(0, "d");
    if (list->Size() != 3) {
        return false;
    }
    if (list->Get(0) != "d") {
        return false;
    }
    if (list->Get(1) != "b") {
        return false;
    }
    if (list->Get(2) != "c") {
        return false;
    }

    list->Set(-1, "e");
    if (list->Size() != 3) {
        return false;
    }
    if (list->Get(0) != "d") {
        return false;
    }
    if (list->Get(1) != "b") {
        return false;
    }
    if (list->Get(2) != "c") {
        return false;
    }

    list->Set(3, "f");
    if (list->Size() != 3) {
        return false;
    }
    if (list->Get(0) != "d") {
        return false;
    }
    if (list->Get(1) != "b") {
        return false;
    }
    if (list->Get(2) != "c") {
        return false;
    }

    return true;
}

bool testSize() {
    List* list = new List;
    list->Add("a");
    list->Add("b");
    list->Add("c");
    if (list->Size() != 3) {
        return false;
    }

    list->Add("d");
    if (list->Size() != 4) {
        return false;
    }

    list->Remove(2);
    if (list->Size() != 3) {
        return false;
    }

    return true;
}

int main() {
    if (!testAdd1()) {
        std::cout << "testAdd1 failed" << std::endl;
    }
    if (!testAdd2()) {
        std::cout << "testAdd2 failed" << std::endl;
    }
    if (!testContains()) {
        std::cout << "testContains failed" << std::endl;
    }
    if (!testIsEmpty()) {
        std::cout << "testIsEmpty failed" << std::endl;
    }
    if (!testRemove()) {
        std::cout << "testRemove failed" << std::endl;
    }
    if (!testSet()) {
        std::cout << "testSet failed" << std::endl;
    }
    if (!testSize()) {
        std::cout << "testSize failed" << std::endl;
    }
}
