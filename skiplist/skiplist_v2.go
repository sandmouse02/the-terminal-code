package skiplist

import (
	"math/rand/v2"
)

const NLOWPTRS = 4

type Node_v2 struct {
	key      uint32
	level    int32
	forwLow  [NLOWPTRS]*Node_v2
	forwHigh []*Node_v2
}

type Skiplist_v2 struct {
	level    int32
	maxLevel int32
	p        float64
	header   *Node_v2
}

func (s *Skiplist_v2) MakeNode_v2(key uint32, level int32) *Node_v2 {
	n := &Node_v2{key: key, level: level}
	if level >= NLOWPTRS {
		n.forwHigh = make([]*Node_v2, level-NLOWPTRS+1)
	}
	return n
}

func New_v2(maxLevel int32) *Skiplist_v2 {
	/* adjust for array indexing */
	maxLevel = maxLevel - 1

	/* header node */
	h := &Node_v2{
		key:      0,
		level:    maxLevel,
		forwHigh: make([]*Node_v2, maxLevel-NLOWPTRS+1)}

	/* last node */
	n := &Node_v2{key: 1 << 31, level: maxLevel}

	var i int32
	for i = 0; i <= maxLevel; i++ {
		if i < NLOWPTRS {
			h.forwLow[i] = n
		} else {
			h.forwHigh[i-NLOWPTRS] = n
		}
	}

	s := &Skiplist_v2{level: 0, maxLevel: maxLevel, p: 0.5, header: h}
	return s
}

func (s *Skiplist_v2) Search_v2(searchKey uint32) *Node_v2 {
	x := s.header

	var i int32
	for i = s.level; i >= NLOWPTRS; i-- {
		for x.forwHigh[i-NLOWPTRS].key < searchKey {
			x = x.forwHigh[i-NLOWPTRS]
		}
	}

	for i = NLOWPTRS - 1; i >= 0; i-- {
		for x.forwLow[i].key < searchKey {
			x = x.forwLow[i]
		}
	}

	x = x.forwLow[0]
	if x.key == searchKey {
		return x
	} else {
		return nil
	}
}

func (s *Skiplist_v2) RandomeLevel_v2() int32 {
	var level int32 = 0

	for rand.Float64() < s.p && level < s.maxLevel {
		level = level + 1
	}

	return level
}

func (s *Skiplist_v2) Insert_v2(searchKey uint32) {
	var i int32

	update := make([]*Node_v2, s.maxLevel+1)
	x := s.header

	/* search */
	for i = s.level; i >= NLOWPTRS; i-- {
		for x.forwHigh[i-NLOWPTRS].key < searchKey {
			x = x.forwHigh[i-NLOWPTRS]
		}
		update[i] = x
	}

	for i = NLOWPTRS - 1; i >= 0; i-- {
		for x.forwLow[i].key < searchKey {
			x = x.forwLow[i]
		}
		update[i] = x
	}

	x = x.forwLow[0]

	/* insert */
	if x.key != searchKey {
		level := s.RandomeLevel_v2()
		if level > s.level {
			for i = s.level + 1; i <= level; i++ {
				update[i] = s.header
			}
			s.level = level
		}

		n := s.MakeNode_v2(searchKey, level)
		for i = 0; i <= level; i++ {
			if i < NLOWPTRS {
				n.forwLow[i] = update[i].forwLow[i]
				update[i].forwLow[i] = n
			} else {
				n.forwHigh[i-NLOWPTRS] = update[i].forwHigh[i-NLOWPTRS]
				update[i].forwHigh[i-NLOWPTRS] = n
			}
		}
	}
}
