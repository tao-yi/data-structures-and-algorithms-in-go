package algorithms

func BinarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if target == arr[mid] {
			return true
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}
