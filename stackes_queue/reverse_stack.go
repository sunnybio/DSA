package stackes_queue

func Reverse_stack(arr []int) []int {

	if len(arr) == 2 {
		arr[0], arr[1] = arr[1], arr[0]
		return arr
	}

	return append(Reverse_stack(arr[1:]), arr[0])

}
