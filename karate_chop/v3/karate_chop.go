package v3

func Chop(target int, array []int) int {
    var search func(low, high int) int
    
    search = func(low, high int) int {
        if low > high {
            return -1
        }
        
        mid := low + (high - low) / 2
        
        if array[mid] == target {
            return mid
        } else if array[mid] < target {
            return search(mid + 1, high)
        } else {
            return search(low, mid - 1)
        }
    }
    
    return search(0, len(array) - 1)
}
