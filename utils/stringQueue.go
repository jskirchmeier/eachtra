package utils

import (
	"strings"
)

type StringQueue interface {
	Next() string
	Peek() string
	More() bool
	Remaining() int
	Source() []string
}

type defaultStringQueue struct {
	q   []string
	idx int
}

func NewStringQueueFromLine(s string) StringQueue {
	return NewStringQueue(strings.Fields(s))
}

func NewStringQueue(s []string) StringQueue {
	return &defaultStringQueue{s, 0}
}

func (sq *defaultStringQueue) Next() (s string) {
	if sq.idx < len(sq.q) {
		s = sq.q[sq.idx]
		sq.idx++
	}
	return
}

func (sq *defaultStringQueue) Peek() (s string) {
	if sq.idx < len(sq.q) {
		s = sq.q[sq.idx]
	}
	return
}

func (sq *defaultStringQueue) More() bool {
	return sq.idx < len(sq.q)
}

func (sq *defaultStringQueue) Remaining() int {
	return len(sq.q) - sq.idx
}

func (sq *defaultStringQueue) Source() []string {
	return sq.q
}
