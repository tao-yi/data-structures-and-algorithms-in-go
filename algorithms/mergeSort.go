package algorithms

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	// recursive split arr into half until there only one element
	l := MergeSort(arr[:mid])
	r := MergeSort(arr[mid:])
	// merge left and right in sorted way
	// fmt.Println("merging", l, r)
	return merge(l, r)
}

// merge two sorted array into one sorted array
// O(m+n)
func merge(l []int, r []int) []int {
	var i, j int
	output := []int{}

	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			output = append(output, l[i])
			i++
		} else {
			output = append(output, r[j])
			j++
		}
	}

	for i < len(l) {
		output = append(output, l[i])
		i++
	}
	for j < len(r) {
		output = append(output, r[j])
		j++
	}
	return output
}
