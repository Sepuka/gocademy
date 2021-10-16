package lamport

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	handlers  = 10
	maxNsec   = 999999999
	ttlBranch = time.Second
)

func TestLamport(t *testing.T) {
	type (
		operation struct {
			uuid   string
			amount int
		}
	)
	var (
		i        = 0
		b        *branch
		branches []*branch
		clock    time.Time
		nsec     int
		waiter   = &sync.WaitGroup{}
		result   = map[int][]operation{}
		knownOps = map[string]int{}
	)

	waiter.Add(handlers * 2)

	for i = 0; i < handlers; i++ {
		rand.Seed(int64(i))
		nsec = rand.Intn(maxNsec) + maxNsec
		clock = time.Date(2021, 10, 11, 23, 03, 0, nsec, time.UTC)
		b = newBranch(i, clock)
		branches = append(branches, b)

		go b.Work(waiter)
		go func(b *branch) {
			defer waiter.Done()
			time.Sleep(ttlBranch)
			b.Stop()
		}(b)
	}

	waiter.Wait()

	for _, b = range branches {
		for _, trx := range b.trxRegistry {
			result[trx.pos] = append(result[trx.pos], operation{
				uuid:   trx.unique,
				amount: trx.amount,
			})
		}
	}

	for _, op := range result {
		for _, o := range op {
			if el, ok := knownOps[o.uuid]; ok {
				assert.Positive(t, o.amount)
				assert.Negative(t, el)
			} else {
				knownOps[o.uuid] = o.amount
			}
		}
	}
}
