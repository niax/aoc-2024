package collections

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func NewSetWithCapacity[T comparable](size int) Set[T] {
	return make(Set[T], size)
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Contains(v T) bool {
	_, contains := s[v]
	return contains
}
