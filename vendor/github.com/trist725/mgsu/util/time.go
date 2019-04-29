package util

import "time"

// 获取这一分钟时间戳
func GetThisMinute() time.Time {
	return time.Now().Truncate(time.Minute)
}

// 获取下一分钟时间戳
func GetNextMinute() time.Time {
	return GetThisMinute().Add(1 * time.Minute)
}

// 获取当前和下一分钟的时差
func GetDurationToNextMinute() time.Duration {
	nextMinute := GetNextMinute()
	return nextMinute.Sub(time.Now())
}

// 获取距离某时刻n天的开始时刻
func GetTheDayBeginTime(src time.Time, days int) (desc time.Time) {
	desc = src.AddDate(0, 0, days)
	desc = time.Date(desc.Year(),
		desc.Month(),
		desc.Day(),
		0, 0, 0, 0, desc.Location())
	return
}

// 获取整点和半小时
// 例如: 当前是19:20, 则获取到的是19:30, 如果是19:40, 获取到的是20:00
func GetNextHalfHour() time.Time {
	now := time.Now()
	next := now.Truncate(30 * time.Minute)
	if now.After(next) {
		next = next.Add(30 * time.Minute)
	}
	return next
}

// 获取下个整点或半小时到当前的时差
func GetDurationToNextHalfHour() time.Duration {
	return GetNextHalfHour().Sub(time.Now())
}

const DefaultTimeFormat = "2006-01-02 15:04:05"

// 解析时间字符串, 返回UTC时间
func ParseTimeStringByFormat(fmt string, strTime string) (time.Time, error) {
	t, err := time.Parse(fmt, strTime)
	return t, err
}
func ParseTimeString(strTime string) (time.Time, error) {
	t, err := time.Parse(DefaultTimeFormat, strTime)
	return t, err
}

// 解析时间字符串, 返回本地时间
func ParseTimeStringInLocationByFormat(fmt string, strTime string) (time.Time, error) {
	t, err := time.ParseInLocation(fmt, strTime, time.Local)
	return t, err
}

func ParseTimeStringInLocation(strTime string) (time.Time, error) {
	t, err := time.ParseInLocation(DefaultTimeFormat, strTime, time.Local)
	return t, err
}
