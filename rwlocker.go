package rwlocker

import "sync"

// A RWLocker represents an object that can be locked and unlocked
// by multiple readers or a single writer.
//
// It adds two extra methods RLock and RUnlock and embeds [sync.Locker]
type RWLocker interface {
	sync.Locker
	RLock()
	RUnlock()
}

// NullRWLocker is a null implementation of RWLocker.
type NullRWLocker struct{}

// Lock locks the NullRWLocker.
func (*NullRWLocker) Lock() {}

// Unlock unlocks the NullRWLocker.
func (*NullRWLocker) Unlock() {}

// RLock locks the NullRWLocker for reading.
func (*NullRWLocker) RLock() {}

// RUnlock unlocks the NullRWLocker for reading.
func (*NullRWLocker) RUnlock() {}
