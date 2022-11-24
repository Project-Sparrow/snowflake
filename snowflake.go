package snowflake

import (
	"sync"
	"time"
)

var (
	workerID  int
	processID int
	epoch     time.Time
	increment int
	mtx       sync.Mutex
)

func Init(e time.Time, w, p int) {
	epoch = e
	workerID = w
	processID = p
	increment = 0
}

type Snowflake uint64

func Generate() Snowflake {
	mtx.Lock()
	defer mtx.Unlock()
	s := Snowflake(0)

	timeComp := time.Since(epoch)
	s |= Snowflake(timeComp << 22)
	s |= Snowflake(workerID << 17)
	s |= Snowflake(processID << 12)
	s |= Snowflake(increment)

	increment++

	return s
}
