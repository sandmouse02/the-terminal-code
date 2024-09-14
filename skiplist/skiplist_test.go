package skiplist

import (
	"testing"

	"math/rand/v2"
)

const maxLevel = 16
const nkeys = 65536

func shuffle(keys []uint32) {
	rand.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})
}

func makekeys(nkeys int) []uint32 {
	keys := make([]uint32, nkeys)

	for i := 0; i < len(keys); i++ {
		keys[i] = uint32(i + 1)
	}

	return keys
}

func setup_v1(maxLevel int, nkeys int) (*Skiplist_v1, []uint32) {
	s := New_v1(maxLevel)
	keys := makekeys(nkeys)

	return s, keys
}

func TestSkiplist_v1(t *testing.T) {
	s, keys := setup_v1(maxLevel, nkeys)

	if s == nil {
		t.Fatalf("New_v1() returned nil")
	}

	shuffle(keys)

	for i := 0; i < len(keys); i++ {
		s.Insert_v1(keys[i])
	}

	shuffle(keys)

	for i := 0; i < len(keys); i++ {
		n := s.Search_v1(keys[i])
		if n == nil {
			t.Fatalf("Search_v1() returned nil")
		}
	}
}

func setup_v2(maxLevel int, nkeys int) (*Skiplist_v2, []uint32) {
	s := New_v2(int32(maxLevel))
	keys := makekeys(nkeys)

	return s, keys
}

func TestSkiplist_v2(t *testing.T) {
	s, keys := setup_v2(maxLevel, nkeys)

	if s == nil {
		t.Fatalf("New_v2() returned nil")
	}

	shuffle(keys)

	for i := 0; i < len(keys); i++ {
		s.Insert_v2(keys[i])
	}

	shuffle(keys)

	for i := 0; i < len(keys); i++ {
		n := s.Search_v2(keys[i])
		if n == nil {
			t.Fatalf("Search_v2() returned nil")
		}
	}
}

func BenchmarkSkiplist_insert_v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		s, keys := setup_v1(maxLevel, nkeys)
		if s == nil {
			b.Fatalf("New_v1() returned nil")
		}

		shuffle(keys)

		b.StartTimer()

		for i := 0; i < len(keys); i++ {
			s.Insert_v1(keys[i])
		}
	}
}

func BenchmarkSkiplist_search_v1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		s, keys := setup_v1(maxLevel, nkeys)
		if s == nil {
			b.Fatalf("New_v1() returned nil")
		}

		for i := 0; i < len(keys); i++ {
			s.Insert_v1(keys[i])
		}
		shuffle(keys)

		b.StartTimer()

		for i := 0; i < len(keys); i++ {
			n := s.Search_v1(keys[i])
			if n == nil {
				b.Fatalf("Search_v1() returned nil")
			}
		}
	}
}

func BenchmarkSkiplist_insert_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		s, keys := setup_v2(maxLevel, nkeys)
		if s == nil {
			b.Fatalf("New_v2() returned nil")
		}

		shuffle(keys)

		b.StartTimer()

		for i := 0; i < len(keys); i++ {
			s.Insert_v2(keys[i])
		}
	}
}

func BenchmarkSkiplist_search_v2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		s, keys := setup_v2(maxLevel, nkeys)
		if s == nil {
			b.Fatalf("New_v2() returned nil")
		}

		for i := 0; i < len(keys); i++ {
			s.Insert_v2(keys[i])
		}
		shuffle(keys)

		b.StartTimer()

		for i := 0; i < len(keys); i++ {
			n := s.Search_v2(keys[i])
			if n == nil {
				b.Fatalf("Search_v2() returned nil")
			}
		}
	}
}
