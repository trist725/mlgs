package util

import (
	"math/rand"
	"sync"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// 生成一定长度的随机字节数组
func GenRandomByteArray(size int) []byte {
	diceLock.Lock()
	defer diceLock.Unlock()

	b := make([]byte, size)
	for i, cache, remain := size-1, dice.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = dice.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

// 生成一定长度的随机字符串
func GenRandomString(size int) string {
	return string(GenRandomByteArray(size))
}

// rand不是线程/协程安全，必须加锁
var dice = rand.New(rand.NewSource(time.Now().UnixNano()))
var diceLock = &sync.Mutex{}

// 注意一下RandomXX函数族得到的是 [min, max)
func RandomTimeDuration(min time.Duration, max time.Duration) time.Duration {
	if min == max {
		return min
	}

	diceLock.Lock()
	defer diceLock.Unlock()

	if min < max {
		return min + time.Duration(dice.Int63n(int64(max-min)))
	}
	return max + time.Duration(dice.Int63n(int64(min-max)))
}

func RandomInt32(min int32, max int32) int32 {
	if min == max {
		return min
	}

	diceLock.Lock()
	defer diceLock.Unlock()

	if min < max {
		return min + dice.Int31n(max-min)
	}
	return max + dice.Int31n(min-max)
}

func RandomInt(min int, max int) int {
	if min == max {
		return min
	}

	diceLock.Lock()
	defer diceLock.Unlock()

	if min < max {
		return min + dice.Intn(max-min)
	}
	return max + dice.Intn(min-max)
}

func RandomInt64(min int64, max int64) int64 {
	if min == max {
		return min
	}

	diceLock.Lock()
	defer diceLock.Unlock()

	if min < max {
		return min + dice.Int63n(max-min)
	}
	return max + dice.Int63n(min-max)
}
