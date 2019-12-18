package main

type F struct {
	J int
}

type A interface {
	a(int) int
}

func main() {
	a := 1
	b := a+1
	fmt.Println(b)
}

