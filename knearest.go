package knearest

import (
	"sort"
)

type Comparison int

const (
	LT Comparison = iota
	GT
	ET
)

type Comparable interface {
	Compare(other Comparable) Comparison
}

type KNearestStorable interface {
	Add(s Comparable)
	Get() []Comparable
}

func NewKNearestStorable(k int) KNearestStorable {
	l := sliceKNearest{
		k:              k,
		currentNearest: []Comparable{},
	}

	return &l
}

type sliceKNearest struct {
	currentNearest []Comparable
	k              int
}

func (l *sliceKNearest) Add(s Comparable) {
	i := sort.Search(len(l.currentNearest), func(i int) bool {
		return s.Compare(l.currentNearest[i]) == LT
	})

	inserted := append(l.currentNearest[:i], append([]Comparable{s}, l.currentNearest[i:]...)...)

	if len(inserted) > l.k {
		l.currentNearest = inserted[:l.k]
	} else {
		l.currentNearest = inserted
	}
}

func (l *sliceKNearest) Get() []Comparable {
	return l.currentNearest
}
