package math

//n^2
func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] > arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = true
			}
		}
		if flag {
			break
		}
	}
}

//n^2
func InsertSort(arr []int) {
	tmp := 0
	for i := 1; i < len(arr); i++ {
		tmp = arr[i]
		for j := i - 1; j >= 0 && tmp < arr[i]; j-- {
			arr[i+1] = arr[i]
		}
	}
}

//n^2
func SelectSort(arr []int) {
	length, minIndex := len(arr), 0
	for i := 0; i < length; i++ {
		minIndex = i
		for j := 1; j < length-i; j++ {
			if arr[j] > arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
