package main

func Loop(nums []int, step int) {
	l := len(nums)
	for i := 0; i < step; i++ {
		for j := i; j < l; j += step {
			nums[j] = 4
		}
	}
}

func Local() {
	a := make(map[int]bool, 10000)
	for i := 0; i < 5000; i++ {
		a[i] = true
	}
}

func Local1() {
	a := make(map[int]bool, 10000)
	for i := 0; i < 10000; i = i + 2 {
		a[i] = true
	}
}
