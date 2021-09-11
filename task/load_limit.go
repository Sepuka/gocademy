package task

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Limiter struct {
	handled int
	limit   int
	period  int
	bucket  map[string]int
	sync.Mutex
}

func NewLimiter(limit int, milliSec int) *Limiter {
	return &Limiter{
		limit:  limit,
		period: milliSec / 2,
		bucket: make(map[string]int),
	}
}

func (l *Limiter) Request() {
	var (
		sec, milliSec                   = getMilliSec()
		curBucketKey, previousBucketKey = l.getKeys(sec, milliSec/l.period)
	)

	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	cur, hasCurrentBucketInitiated := l.bucket[curBucketKey]
	prev, _ := l.bucket[previousBucketKey]
	if hasCurrentBucketInitiated {
		if cur+prev < l.limit {
			l.bucket[curBucketKey]++
			l.handled++
		}
	} else if prev < l.limit {
		l.bucket[curBucketKey] = 1
		l.handled++
	}
}

func (l *Limiter) getHandled() int {
	return l.handled
}

func (l *Limiter) getBucket() map[string]int {
	return l.bucket
}

func (l *Limiter) getKeys(currentTime int64, batch int) (string, string) {
	var (
		previousTime = currentTime
		prevBatch    = batch - 1
	)

	if prevBatch < 0 {
		prevBatch = 9
		previousTime = currentTime - 1
	}

	return fmt.Sprintf(`%d_%d`, currentTime, batch), fmt.Sprintf(`%d_%d`, previousTime, prevBatch)
}

func getMilliSec() (int64, int) {
	var (
		currentTime = time.Now()
	)
	milliString := currentTime.Format(`.000`)
	milli, _ := strconv.ParseInt(milliString[1:], 10, 64)

	return currentTime.Unix(), int(milli)
}
