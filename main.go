package main

/*
#include <stdio.h>
#include <stdlib.h>

void say_hello() {
        printf("Hello World!\n");
}
*/
import "C"

func main() {
	C.say_hello()
}
