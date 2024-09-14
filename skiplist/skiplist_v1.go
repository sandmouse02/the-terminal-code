package skiplist

import (
	"math/rand/v2"
)

type Node_v1 struct {
	key     uint32
	level   int
	forward []*Node_v1
}

type Skiplist_v1 struct {
	level    int
	maxLevel int
	p        float64
	header   *Node_v1
}

func (s *Skiplist_v1) MakeNode_v1(key uint32, level int) *Node_v1 {
	n := &Node_v1{key: key, level: level, forward: make([]*Node_v1, level+1)}
	return n
}

func New_v1(maxLevel int) *Skiplist_v1 {
	/* adjust for array indexing */
	maxLevel = maxLevel - 1

	/* header node */
	h := &Node_v1{
		key:     0,
		level:   maxLevel,
		forward: make([]*Node_v1, maxLevel+1)}

	/* last node */
	n := &Node_v1{key: 1 << 31, level: maxLevel, forward: nil}

	for i := 0; i <= maxLevel; i++ {
		h.forward[i] = n
	}

	s := &Skiplist_v1{level: 0, maxLevel: maxLevel, p: 0.5, header: h}
	return s
}

func (s *Skiplist_v1) Search_v1(searchKey uint32) *Node_v1 {
	x := s.header

	for i := s.level; i >= 0; i-- {
		for x.forward[i].key < searchKey {
			x = x.forward[i]
		}
	}

	x = x.forward[0]
	if x.key == searchKey {
		return x
	} else {
		return nil
	}
}

func (s *Skiplist_v1) RandomLevel_v1() int {
	level := 0

	for rand.Float64() < s.p && level < s.maxLevel {
		level = level + 1
	}

	return level
}

func (s *Skiplist_v1) Insert_v1(searchKey uint32) {
	update := make([]*Node_v1, s.maxLevel+1)
	x := s.header

	/* search */
	for i := s.level; i >= 0; i-- {
		for x.forward[i].key < searchKey {
			x = x.forward[i]
		}
		update[i] = x
	}

	x = x.forward[0]

	/* insert */
	if x.key != searchKey {
		level := s.RandomLevel_v1()
		if level > s.level {
			for i := s.level + 1; i <= level; i++ {
				update[i] = s.header
			}
			s.level = level
		}

		n := s.MakeNode_v1(searchKey, level)
		for i := 0; i <= level; i++ {
			n.forward[i] = update[i].forward[i]
			update[i].forward[i] = n
		}
	}
}
