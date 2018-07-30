package syntax

func factorial(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func sliceAppend() {
	s := make([]int, 0)
	for i := 0; i < 32; i++ {
		s = append(s, i)
	}
}

func maxSubsequnceSum(A []int) int {
	thisSum, maxSum := 0, 0
	for i := 0; i < len(A); i++ {
		thisSum = 0
		for j := i; j < len(A); j++ {
			thisSum += A[j]
			if thisSum > maxSum {
				maxSum = thisSum
			}
		}
	}
	return maxSum
}

func insertSort(A []int, n int) {
	j := 0
	for i := 1; i < n; i++ {
		tmp := A[i]
		for j = i; j > 0 && tmp < A[j-1]; j-- {
			A[j] = A[j-1]
		}
		A[j] = tmp
	}
}
