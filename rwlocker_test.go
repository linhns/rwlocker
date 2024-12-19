package rwlocker_test

import (
	"sync"
	"testing"

	"github.com/linhns/rwlocker"
)

func TestRWLocker(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		locker rwlocker.RWLocker
	}{
		"sync.RWMutex": {
			locker: &sync.RWMutex{},
		},
		"rwlocker.NullRWLocker": {
			locker: &rwlocker.NullRWLocker{},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			tt.locker.Lock()
			tt.locker.Unlock()

			tt.locker.RLock()
			tt.locker.RUnlock()
		})
	}
}
