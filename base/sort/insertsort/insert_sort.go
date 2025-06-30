package insertsort

// InsertSort 插入排序算法
// 时间复杂度：O(n²)，空间复杂度：O(1)
// 算法思想：将数组分为已排序和未排序两部分，每次从未排序部分取出一个元素插入到已排序部分的正确位置
func InsertSort(nums []int) {
	// 从第二个元素开始遍历，因为第一个元素可以认为是已排序的
	for i := 1; i < len(nums); i++ {
		e := nums[i] // 当前要插入的元素
		var j int
		// 从当前位置向前查找插入位置，同时将大于当前元素的元素向后移动
		for j = i; j > 0 && nums[j-1] > e; j-- {
			nums[j] = nums[j-1] // 将大于当前元素的元素向后移动一位
		}
		nums[j] = e // 将当前元素插入到正确位置
	}
}
