package main

func QuickSort(left int, right int, array *[9]int) {
	l := left
	r := right

	pivot := array[(left+right)/2]
	tmp := 0

	for l < r {
		for array[l] < pivot {
			l++
		}
		for array[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		tmp = array[l]
		array[l] = array[r]
		array[r] = tmp
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {
	//arr := [9]int{-9, 78, 0, 23, -567, 70, 123, 90, -23}
}
