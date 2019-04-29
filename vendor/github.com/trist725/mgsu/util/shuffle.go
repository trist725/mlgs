package util

//Knuth-Durstenfeld Shuffle
//时间复杂度为O(n)
//空间复杂度为O(1)
//必须知道数组长度
func KnuthDurstenfeldShuffle(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		p := RandomInt64(0, int64(i))
		a := arr[i]
		arr[i] = arr[p]
		arr[p] = a
	}
	return arr
}
