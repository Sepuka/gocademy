package lamport

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

const (
	tick       = time.Millisecond * 200
	maxAmount  = 9
	startDelay = 100
)

type (
	// transaction is a atomic domain structure which communicates between branches
	transaction struct {
		pos    int
		amount int
		unique string
	}
	// branch - is a bank branch
	// it has clock (does NOT synced) and the other properties
	branch struct {
		id          int
		pos         int
		quit        chan bool
		clock       time.Time
		trxRegistry []transaction
		pipe        chan transaction
		sync.Mutex
	}
)

var (
	branchRegistry []*branch
)

// creates a new branch
func newBranch(branchId int, clock time.Time) *branch {
	var (
		b = &branch{
			id:    branchId,
			pos:   0,
			quit:  make(chan bool),
			clock: clock,
			pipe:  make(chan transaction, 10),
		}
	)

	branchRegistry = append(branchRegistry, b)

	return b
}

// Stop sends to branch term-signal
func (b *branch) Stop() {
	b.quit <- true
}

// Work generates few transactions
func (b *branch) Work(wg *sync.WaitGroup) {
	var (
		trx    transaction
		ticker = time.NewTicker(tick)
		delay  = time.Duration(rand.Intn(startDelay))
	)

	defer wg.Done()

	time.Sleep(time.Millisecond * delay)

	// It can receive a transaction or send another
	for {
		select {
		case trx = <-b.pipe:
			b.l(trx)
		case <-b.quit:
			return
		case <-ticker.C:
			b.event()
		}
	}
}

// generates outgoing transaction
func (b *branch) event() {
	b.Mutex.Lock()
	b.pos++
	var trx = transaction{
		pos:    b.pos,
		amount: -(rand.Intn(maxAmount) + 1),
		unique: uuid.NewString(),
	}
	b.trxRegistry = append(b.trxRegistry, trx)

	b.Mutex.Unlock()

	b.sendEvent(trx)
}

func (b *branch) sendEvent(trx transaction) {
	var target = rand.Intn(len(branchRegistry))
	branchRegistry[target].pipe <- trx
}

// Lamport's clock
func (b *branch) l(trx transaction) {
	b.Mutex.Lock()
	defer b.Mutex.Unlock()

	var maxPos = 0

	if len(b.trxRegistry) > 0 {
		var lastTrx = b.trxRegistry[len(b.trxRegistry)-1]
		if lastTrx.pos > trx.pos {
			maxPos = lastTrx.pos + 1
		} else {
			maxPos = trx.pos + 1
		}
	}

	b.trxRegistry = append(b.trxRegistry, transaction{
		pos:    maxPos,
		amount: -trx.amount,
		unique: trx.unique,
	})
}
