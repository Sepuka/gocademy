package task

import (
	"encoding/json"
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

const (
	limit          = 20
	milliSecPeriod = 200 // period has length 200 mc
)

type limiterTest struct {
	suite.Suite
	limiter *Limiter
}

func (l *limiterTest) SetupTest() {
	l.limiter = NewLimiter(limit, milliSecPeriod)
}

func (l *limiterTest) Test() {
	var (
		requestCount = 0
	)
	const halfPeriod = milliSecPeriod / 2

	l.waitTheEndOfPeriod(halfPeriod)

	var curTime = time.Now()

	// work the last part of the first period and the first part of the next period
	for time.Now().Sub(curTime) < time.Millisecond*halfPeriod {
		go l.limiter.Request()
		requestCount++
	}

	limiter, _ := json.Marshal(l.limiter.getBucket())

	assert.Equal(l.T(), l.limiter.getHandled(), limit, fmt.Sprintf(`%d requests has sent, %s`, requestCount, limiter))
}

func (l *limiterTest) waitTheEndOfPeriod(halfPeriod int) {
	_, milli := getMilliSec()

	for milli%milliSecPeriod < halfPeriod {
		time.Sleep(time.Millisecond * 10)
		_, milli = getMilliSec()
	}
}

func TestBubbleSort(t *testing.T) {
	suite.Run(t, new(limiterTest))
}
