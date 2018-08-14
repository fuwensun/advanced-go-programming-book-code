// hello.cpp

#include <iostream.h>

extern "C" {
	#include "hello.h"
}

void SayHello(const char* s) {
	std::cout << s;
}