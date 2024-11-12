package mutex_impl

import (
	"runtime"
	"sync/atomic"
)

const (
	spin = 10
)

type Mutex struct {
	state int32
}

type MutexWithChan struct {
	state  int32
	waitCh chan bool
}

type Locker interface {
	Lock()
	Unlock()
}

func NewMutexWithChan() *MutexWithChan {
	return &MutexWithChan{
		state:  0,
		waitCh: make(chan bool, 32),
	}
}

func (mc *MutexWithChan) Lock() {
	for i := 0; ; i++ {
		if atomic.CompareAndSwapInt32(&mc.state, 0, 1) {
			return
		}
		if i >= spin {
			<-mc.waitCh
			i = 0
		}
		runtime.Gosched()
	}
}

func (mc *MutexWithChan) Unlock() {
	atomic.StoreInt32(&mc.state, 0)
	select {
	case mc.waitCh <- true:
	default:
	}
}

func (m *Mutex) Lock() {
	for {
		for {
			//fmt.Println("comp")
			if atomic.CompareAndSwapInt32(&m.state, 0, 1) {
				return
			}
			runtime.Gosched()

			{
			}
		}
	}
}

func (m *Mutex) Unlock() {
	atomic.StoreInt32(&m.state, 0)

}
