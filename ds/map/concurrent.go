package _map

import (
	"sync"

	"github.com/sonirico/stadio/fp"
	"github.com/sonirico/stadio/slices"
	"github.com/sonirico/stadio/tuples"
)

type (
	MapInner[K comparable, V any] Map[K, V]

	Concurrent[K comparable, V any] struct {
		L sync.RWMutex
		MapInner[K, V]
	}
)

func NewConcurrent[K comparable, V any](inner Map[K, V]) *Concurrent[K, V] {
	return &Concurrent[K, V]{MapInner: inner}
}

func (m *Concurrent[K, V]) Get(k K) (v V, ok bool) {
	m.L.RLock()
	v, ok = m.MapInner.Get(k)
	m.L.RUnlock()
	return
}

func (m *Concurrent[K, V]) Has(k K) (ok bool) {
	m.L.RLock()
	_, ok = m.MapInner.Get(k)
	m.L.RUnlock()
	return
}

func (m *Concurrent[K, V]) Set(k K, v V) {
	m.L.Lock()
	m.MapInner.Set(k, v)
	m.L.Unlock()
	return
}

func (m *Concurrent[K, V]) Range(fn func(K, V, int) bool) {
	m.L.RLock()
	defer m.L.RUnlock()
	m.MapInner.Range(fn)
}

func (m *Concurrent[K, V]) Delete(k K) {
	m.L.Lock()
	m.MapInner.Delete(k)
	m.L.Unlock()
}

func (m *Concurrent[K, V]) GetOrSet(k K, def V) (v V, ok bool) {
	m.L.Lock()
	v, ok = m.MapInner.GetOrSet(k, def)
	m.L.Unlock()
	return
}

func (m *Concurrent[K, V]) Map(fn func(K, V) (K, V)) Map[K, V] {
	m.L.RLock()
	defer m.L.RUnlock()
	return &Concurrent[K, V]{MapInner: m.MapInner.Map(fn)}
}

func (m *Concurrent[K, V]) FilterMap(
	fn func(K, V) fp.Option[tuples.Tuple2[K, V]],
) Map[K, V] {
	m.L.RLock()
	defer m.L.RUnlock()
	return &Concurrent[K, V]{MapInner: m.MapInner.FilterMap(fn)}
}

func (m *Concurrent[K, V]) Filter(fn func(K, V) bool) Map[K, V] {
	m.L.RLock()
	defer m.L.RUnlock()
	return &Concurrent[K, V]{MapInner: m.MapInner.Filter(fn)}
}

func (m *Concurrent[K, V]) Values() slices.Slice[V] {
	m.L.RLock()
	res := m.MapInner.Values()
	m.L.RUnlock()
	return res
}

func (m *Concurrent[K, V]) Keys() slices.Slice[K] {
	m.L.RLock()
	res := m.MapInner.Keys()
	m.L.RUnlock()
	return res
}

func (m *Concurrent[K, V]) Entries() slices.Slice[Entry[K, V]] {
	m.L.RLock()
	res := m.MapInner.Entries()
	m.L.RUnlock()
	return res
}
