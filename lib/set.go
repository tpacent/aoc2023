package lib

func NewSet[T comparable](items []T) Set[T] {
	set := make(Set[T], len(items))
	set.Add(items...)
	return set
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Has(key T) (ok bool) {
	_, ok = s[key]
	return
}

func (s Set[T]) Add(items ...T) {
	for _, v := range items {
		s[v] = struct{}{}
	}
}
