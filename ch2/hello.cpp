#include <iostream>

extern "C" {
    #include "hello.h"
}

void SayHello3(const char* s) {
    std::cout << s;
}
