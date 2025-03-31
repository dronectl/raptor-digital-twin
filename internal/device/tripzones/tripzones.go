package tripzones

import (
	"cmp"
	"sync"
)

type Tripzone[T cmp.Ordered] struct {
	tripped bool
	min     T
	max     T
	m       sync.Mutex
}

func (tz *Tripzone[T]) Min(v T) {
	tz.m.Lock()
	tz.min = v
	tz.m.Unlock()
}

func (tz *Tripzone[T]) Max(v T) {
	tz.m.Lock()
	tz.max = v
	tz.m.Unlock()
}

func (tz *Tripzone[T]) CheckTripCondition(v T) bool {
	tz.m.Lock()
	defer tz.m.Unlock()
	if v < tz.min || v > tz.max {
		tz.tripped = true
	} else {
		tz.tripped = false
	}
	return tz.tripped
}

func New[T cmp.Ordered](mx, mi T) *Tripzone[T] {
	return &Tripzone[T]{min: mi, max: mx, tripped: false}
}
