package main

func MergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
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

func main() {}

// MergeSortP .
func MergeSortP(arr []int) []int {
	l := len(arr)
	if l < 2 {
		return arr
	}
	exci := l / 2
	left := MergeSortP(arr[:exci])
	right := MergeSortP(arr[exci:])

	return mergeP(left, right)
}

func mergeP(l, r []int) []int {
	ll := len(l)
	lr := len(r)
	tmp := make([]int, ll+lr)
	i, j, index := 0, 0, 0
	for {
		if l[i] > r[j] {
			tmp[index] = r[j]
			index++
			j++
			if j == lr {
				copy(tmp[index:], r[i:])
				break
			}
		} else {
			tmp[index] = l[i]
			index++
			i++
			if i == ll {
				copy(tmp[index:], l[j:])
				break
			}
		}
	}
	return tmp
}
