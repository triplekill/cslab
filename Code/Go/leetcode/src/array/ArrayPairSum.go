package array


// https://leetcode-cn.com/problems/array-partition-i/description/
// 超出时间限制，说明排序算法效率太低
func arrayPairSum(nums []int) int {
	// sort the array, then make the near two integer as one group
	for i := 0; i < len(nums)-1; i++ {
		for j:=0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	// log.Fatal("test")
	sum := 0	
	for i:=0; i<len(nums)/2; i++ {
		sum += nums[2*i]
	}
	return sum
}

func quickSort(nums []int, low int, high int) {
	if low<high {
		mid := quickSplit(nums, low, high)
		quickSort(nums, low, mid-1)
		quickSort(nums, mid+1, high)
	}
}
func quickSplit(nums []int, low int, high int) int {
	key := nums[low]
	for ; low < high; {
		for low<high && nums[high] >= key { 
			high--
		}
		if low < high {
			nums[low] = nums[high]
			low++
		}
		for low<high && nums[low] <= key {
			low++
		}
		if low < high {
			nums[high] = nums[low]
			high--
		}
	}
	nums[low] = key
	return low
}

func arrayPairSum2(nums []int) int {
	// sort the array, then make the near two integer as one group
	quickSort(nums, 0, len(nums)-1)
	// log.Fatal("test")
	sum := 0	
	for i:=0; i<len(nums)/2; i++ {
		sum += nums[2*i]
	}
	return sum
}
