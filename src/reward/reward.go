package reward

//func (logic *Logic) Reward(reward sd.Reward, n int, notify bool) (err error) {
//	if reward.Len() == 0 {
//		err = fmt.Errorf("empty reward")
//		logger.Error("[%s] reward %+v fail, %v", logic.Sign(), reward, err)
//		return
//	}
//	if n <= 0 {
//		err = fmt.Errorf("n <= 0")
//		logger.Error("[%s] reward %+v fail, %v", logic.Sign(), reward, err)
//		return
//	}
//	if n > 1 {
//		reward = reward.Multiple(n)
//	}
//
//	reward = reward.Format()
//
//	reward.Each(func(item sd.RewardItem) {
//		err = logic.item.Gain(item.TID(), item.Num(), notify)
//		if err != nil {
//			logger.Error("[%s] reward %+v fail, %v", logic.Sign(), reward, err)
//			return
//		}
//	})
//
//	return
//}
//
//func (logic *Logic) RollRewardGroup(id int) (reward sd.Reward, err error) {
//	s, d := sd.GetRewardGroup(id)
//	if s == nil {
//		err = fmt.Errorf("static data not found")
//		logger.Error("[%s] roll reward group [%d] fail, %v", logic.Sign(), id, err)
//		return
//	}
//	if d == nil {
//		err = fmt.Errorf("dice found")
//		logger.Error("[%s] roll reward group [%d] fail, %v", logic.Sign(), id, err)
//		return
//	}
//
//	for i := 0; i < s.RollNum; i++ {
//		reward = reward.AppendItem(d.Roll().(sd.RewardGroupItem).ToRewardItem())
//	}
//
//	return
//}
