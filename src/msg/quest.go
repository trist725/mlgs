package msg

type QuestSlice []*Quest

//菜鸡客户端非要我排序了再发,肿么办。。
func (qs QuestSlice) Len() int { // 重写 Len() 方法
	return len(qs)
}
func (qs QuestSlice) Swap(i, j int) { // 重写 Swap() 方法
	qs[i], qs[j] = qs[j], qs[i]
}

// 重写 Less() 方法，以任务Id大小排序
// 从小到大排
func (qs QuestSlice) Less(i, j int) bool {
	if qs[i].Id > qs[j].Id {
		return false
	} else {
		return true
	}
}
