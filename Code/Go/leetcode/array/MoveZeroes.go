package array

func moveZeroes(nums []int)  {
	if len(nums) <2 {
		return
	}
	index := 0
	worker:= 0
	for worker < len(nums) {
		if nums[worker]!=0 {
			nums[index], nums[worker] = nums[worker], nums[index]
			index++
		}
		worker++
	}
}
