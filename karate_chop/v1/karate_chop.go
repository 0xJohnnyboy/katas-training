package v1

func Chop(target int, array []int) int {
	mid := 0
	n := len(array) - 1
	start := 0
	end := n
	found := false

	for found == false && start <= end {
		mid = int((start + end) / 2)
		if array[mid] == target {
			found = true
			continue
		}
		if target > array[mid] {
			start = mid + 1
			continue
		}
		end = mid - 1
	}

	if found {
		return mid
	}

	return -1
}
