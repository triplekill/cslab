package algorithm

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[head], data[i] = data[i], data[head]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}

func quicksort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[tail], data[i] = data[i], data[tail]
			tail--
		} else {
			data[head], data[i] = data[i], data[head]
			head++
			i++
		}
	}
	quicksort(data[:head])
	quicksort(data[head+1:])
}
