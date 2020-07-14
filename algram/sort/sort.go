// Package sort is used to organize all sorting algrorithms
package sort

// Arr is int array for sorting
type Arr []int

// MergeSort sort tc:nlogn sc:n stable
func (arr Arr) MergeSort() []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	key := n / 2
	left := arr[0:key].MergeSort()
	right := arr.MergeSort()
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	newArr := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for {
		if left[i] > right[j] {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}
		} else {
			newArr[index] = left[i]
			index++
			i++
			if i == len(left) {
				copy(newArr[index:], right[j:])
				break
			}
		}
	}
	return newArr
}

// tc:n^2 insertion,bubble stable

// Insertion sort tc:n^2 sc:1 stable
func (arr Arr) Insertion() {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		for j := i; j > 0 && arr[j] > arr[j-1]; j++ {
			arr[j] = arr[j-1]
		}
		arr[i] = tmp
	}
}

// Bubble sort tc:n^2 sc:1 stable
func (arr Arr) Bubble() {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	for i := 0; i < arrLen; i++ {
		for j := arrLen - 1; j > i; j++ {
			if arr[j] < arr[j-1] {
				tmp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return
}

// Selection sort tc:n^2 sc:1 unstable
func (arr Arr) Selection() {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

// Shell sort tc:n^3/2 sc:1 unstable
func (arr Arr) Shell() {

}

// Heap sort tc:nlogn sc:n unstable
func (arr Arr) Heap() {

}

// Quick sort tc:nlogn sc:n stable : unconfirm
func (arr Arr) Quick() {

}
