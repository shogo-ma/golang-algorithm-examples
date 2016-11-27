package sort

func Selection(array []int) []int {
	// selection sort

	smallset := 0
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[smallset] {
				smallset = j
			}
		}
		array[i], array[smallset] = array[smallset], array[i]
	}

	return array
}

func Insertion(array []int) []int {
	// insertion sort
	n := len(array)

	for i := 1; i < n; i++ {
		key := array[i]

		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}

	return array
}

func Merge(array []int, p, r int) []int {
	if p >= r {
		return []int{array[p]}
	}

	q := (p + r) / 2
	lefts := Merge(array, p, q)
	rights := Merge(array, q+1, r)
	fix := merge(lefts, rights)

	return fix
}

func merge(lefts, rights []int) []int {
	fix := []int{}
	left_index, right_index := 0, 0

	for len(lefts) > left_index && len(rights) > right_index {
		if lefts[left_index] <= rights[right_index] {
			fix = append(fix, lefts[left_index])
			left_index++
		} else {
			fix = append(fix, rights[right_index])
			right_index++
		}
	}

	if len(lefts) > left_index {
		fix = append(fix, lefts[left_index:]...)
	}

	if len(rights) > right_index {
		fix = append(fix, rights[right_index:]...)
	}

	return fix
}
