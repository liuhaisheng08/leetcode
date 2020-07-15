package math

//arr 有序 不重复
func BSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

//arr 有序 可重复 返回第一个
func BSearch2(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != target {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

//arr 有序 可重复 返回第最后一个
func BSearch3(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

//arr 有序 可重复 返回第一个 >target
func BSearch4(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid+1] != target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}
