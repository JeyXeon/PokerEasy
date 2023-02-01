package common

import "sync"

type RWLockerMap[T any] struct {
	mx sync.RWMutex
	m  map[int]T
}

func (rwLockerMap *RWLockerMap[T]) Load(key int) (T, bool) {
	rwLockerMap.mx.RLock()
	defer rwLockerMap.mx.RUnlock()
	val, ok := rwLockerMap.m[key]
	return val, ok
}

func (rwLockerMap *RWLockerMap[T]) Store(key int, value T) {
	rwLockerMap.mx.Lock()
	defer rwLockerMap.mx.Unlock()
	rwLockerMap.m[key] = value
}

func (rwLockerMap *RWLockerMap[T]) Delete(key int) {
	rwLockerMap.mx.Lock()
	defer rwLockerMap.mx.Unlock()
	delete(rwLockerMap.m, key)
}

func NewRWLockerMap[T any](containment map[int]T) *RWLockerMap[T] {
	rwLockerMap := new(RWLockerMap[T])
	rwLockerMap.m = containment
	return rwLockerMap
}
