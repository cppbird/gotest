package main

/*
#include <stdio.h>

void printint() {
 	int *ptr = NULL;
	*ptr = 0;
}
*/
// import "C"
import "runtime"

func main() {
	runtime.NumCPU()
}
