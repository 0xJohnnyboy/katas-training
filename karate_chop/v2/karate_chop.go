package v2

func Chop(target int, array []int) int {
	low, high := 0, len(array)-1
	return binarySearch(target, array, low, high)
}

func binarySearch(target int, array []int, low int, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2
	if array[mid] == target {
		return mid
	}
	if target < array[mid] {
		return binarySearch(target, array, low, mid-1)
	}
	return binarySearch(target, array, mid+1, high)
}
