package binary_search

func BinarySearch[T ~int64 | ~int | ~float32](e T, array []T) bool {
	high := len(array) - 1
	low := 0
	done := false
	for !done {
		mid := low + (high-low)/2
		switch {
		case low > high:
			done = true
		case e > array[mid]:
			low = mid + 1
		case e < array[mid]:
			high = mid - 1
		case e == array[mid]:
			return true
		}

	}
	return false
}
