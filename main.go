package main

/*
#include <stdio.h>

void printint() {
 	int *ptr = NULL;
	*ptr = 0;
}
*/
import "C"

func main() {

	C.printint()
}

// func t() {
// 	s := struct {
// 		a byte
// 		b byte
// 		c byte
// 		d int64
// 	}{0, 0, 0, 0}

// 	// 将结构体指针转换为通用指针
// 	p := unsafe.Pointer(&s)
// 	up0 := uintptr(p)
// 	fmt.Println(up0)
// 	up0 += 10000000000
// 	fmt.Println(up0)
// 	pp := unsafe.Pointer(up0)
// 	p1 := (*byte)(pp)
// 	fmt.Println(*p1)
// }
