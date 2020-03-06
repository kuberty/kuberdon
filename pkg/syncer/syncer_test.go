package syncer

import (
	"testing"
	"time"
)

func getMockFunction(delay time.Duration, callCh chan struct{}) func() {
	return func() {
		time.Sleep(delay)
		callCh <- struct{}{}
	}
}

func TestClosedInTime(t *testing.T) {
	t.Parallel()
	syncer := LocalSyncer{}
	callCh := make(chan struct{})

	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh))

	select {
	case <-callCh: //It is closed, yey!
		select {
		case <-callCh:
			t.Fatalf("Function should only be called once")
		case <-time.After(110 * time.Millisecond):
		}
	case <-time.After(110 * time.Millisecond):
		t.Fatalf("Function did not complete in time.")
	}
}

func TestScheduledCompletesInTime(t *testing.T) {
	t.Parallel()
	syncer := LocalSyncer{}
	callCh1 := make(chan struct{})
	callCh2 := make(chan struct{})

	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh1))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh2))
	<-callCh1 // Wait for the first to complete

	select {
	case <-callCh2: //It is closed, yey!
		select {
		case <-callCh1:
			t.Fatalf("Function 1 should only be called once")
		case <-callCh2:
			t.Fatalf("Function 2 should only be called once")
		case <-time.After(110 * time.Millisecond):
		}
	case <-time.After(110 * time.Second):
		t.Fatalf("Function 2 did not complete in time.")
	}
}
func TestLastFunctionTakesPriority(t *testing.T) {
	t.Parallel()
	syncer := LocalSyncer{}
	callCh1 := make(chan struct{})
	callCh2 := make(chan struct{})
	callCh3 := make(chan struct{})

	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh1))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh2))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh3))
	<-callCh1
	select {
	case <-time.After(110 * time.Millisecond):
		t.Fatalf("Third function did not complete in time, was second function prioritized?")
	case <-callCh3: // it is called!

	}
}

func TestIntermediateScheduledFunctionNeverCalled(t *testing.T) {
	t.Parallel()
	syncer := LocalSyncer{}
	callCh1 := make(chan struct{})
	callCh2 := make(chan struct{})
	callCh3 := make(chan struct{})

	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh1))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh2))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh3))
	<-callCh1 // Wait for the first to complete

	select {
	case <-time.After(110 * time.Millisecond):
	case <-callCh2:
		t.Fatalf("Second call schould never be ran when function is already scheduled!")
	}
}
func TestOnlyLastFunctionCalled(t *testing.T) {
	t.Parallel()
	syncer := LocalSyncer{}
	callCh1 := make(chan struct{})
	callCh2 := make(chan struct{})
	callCh2a := make(chan struct{})
	callCh2b := make(chan struct{})
	callCh2c := make(chan struct{})
	callCh3 := make(chan struct{})

	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh1))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh2))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh2a))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh2b))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh2c))
	syncer.ScheduleSync(getMockFunction(100*time.Millisecond, callCh3))
	<-callCh1 // Wait for the first to complete

	select {
	case <-time.After(110 * time.Millisecond):
		select {
		case <-time.After(110 * time.Millisecond):
			t.Fatalf("Last function was not called in time.")
		case <-callCh3:
		}
	case <-callCh2:
		t.Fatalf("Second call schould never be ran when function is already scheduled!")
	case <-callCh2a:
		t.Fatalf("Intermediate call schould never be ran when function is already scheduled!")
	case <-callCh2b:
		t.Fatalf("Intermediate call schould never be ran when function is already scheduled!")
	case <-callCh2c:
		t.Fatalf("Intermediate call schould never be ran when function is already scheduled!")
	}
}
