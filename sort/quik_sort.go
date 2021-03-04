package sort

func QuickSort(arr []int) {
	//skip
	if len(arr) < 1 {
		return
	}
	flag := arr[0]
	left, right := 0, len(arr)-1

	for i := 1; i <= right; {
		if arr[i] > flag {
			arr[i], arr[right] = arr[right], arr[i]
			right--
		} else {
			arr[i], arr[left] = arr[left], arr[i]
			i++
			left++
		}
	}
	QuickSort(arr[:right])
	QuickSort(arr[right+1:])
}
