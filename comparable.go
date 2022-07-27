package optional

type c[T comparable] struct {
	v T
}

func (o *c[T]) Get() T {
	return o.v
}

func (o *c[T]) Empty() bool {
	var empty T

	return o.v == empty
}

func (o *c[T]) IfEmpty(fn func(T)) {
	if !o.Empty() {
		return
	}

	fn(o.v)
}

func (o *c[T]) Present() bool {
	return !o.Empty()
}

func (o *c[T]) IfPresent(fn func(T)) {
	if o.Empty() {
		return
	}

	fn(o.v)
}

func (o *c[T]) If(fn func(T) bool) *c[T] {
	if o.Empty() || fn(o.v) {
		return o
	}

	return emptyComparable[T]()
}

func (o *c[T]) Else(e T) T {
	if o.Empty() {
		return e
	}

	return o.v
}

func emptyComparable[T comparable]() *c[T] {
	return &c[T]{}
}

func Comparable[T comparable](v T) *c[T] {
	return &c[T]{v}
}
