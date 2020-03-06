package syncer

import (
	"fmt"
	"sync"
)

type Syncer interface {
	/**
	 The syncer simply runs one function at a time. It can buffer up to one extra function to run after this.
	 When scheduling while a function is running, a LIFO (last-in-first-out) policy is used, this means that the last function scheduled will be the one actually in the (size-one) queue.

	Within Kuberdon we use this simple scheduler to schedule our synchronization task.
	*/
	ScheduleSync(syncFunc func())
}

//NOTE: The implementation could be slightly improved by storing the mutex inside the actual kubernetes api.
type LocalSyncer struct {
	sync.RWMutex
	scheduledFunc func()
	running       bool
}

/**
The single exposed method for the Syncer struct. This is the method you should use to schedule your functions.

If the syncFunc is not currently running, this will start running it asynchronously. If it is running, this will make sure it is ran once more after it is finished.
*/
func (s *LocalSyncer) ScheduleSync(syncFunc func()) {
	s.Lock()
	defer s.Unlock()

	// If the sync func is not yet running, start it. If it is running, make sure the scheduled function is overwritten with the latest syncFunc
	if !s.running {
		s.running = true
		go s.startSync(syncFunc)
	} else {
		s.scheduledFunc = syncFunc
	}

}

func (s *LocalSyncer) startSync(syncFunc func()) {
	//We make sure that our syncer stops "running" once it is finished, and we schedule a new one if necessary:
	defer func() {
		s.Lock()
		defer s.Unlock()

		s.running = s.isScheduled() // Set the syncer state to whether or not we are running another function
		if s.isScheduled() {
			nextSyncFunc := s.scheduledFunc
			s.scheduledFunc = nil
			go s.startSync(nextSyncFunc)
		}
	}()

	runSyncFuncWithRecovery(syncFunc)
}

func runSyncFuncWithRecovery(syncFunc func()) {
	defer func() { // recovery script
		if r := recover(); r != nil {
			fmt.Printf("Recovered from %v. \n", r)
		}
	}()

	syncFunc() // Call the actual function, if an panic occurs, this will be recovered.
}

/**
Checks whether there is a scheduled function, note that this function is not thread safe!
*/
func (s LocalSyncer) isScheduled() bool {
	return s.scheduledFunc != nil
}
