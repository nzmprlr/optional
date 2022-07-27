package optional

type s[T any] struct {
	v []T
}

func (o *s[T]) Get() []T {
	return o.v
}

func (o *s[T]) Empty() bool {
	return len(o.v) == 0
}

func (o *s[T]) IfEmpty(fn func([]T)) {
	if !o.Empty() {
		return
	}

	fn(o.v)
}

func (o *s[T]) Present() bool {
	return !o.Empty()
}

func (o *s[T]) IfPresent(fn func([]T)) {
	if o.Empty() {
		return
	}

	fn(o.v)
}

func (o *s[T]) If(fn func([]T) bool) *s[T] {
	if o.Empty() || fn(o.v) {
		return o
	}

	return emptySlice[T]()
}

func (o *s[T]) Else(e []T) []T {
	if o.Empty() {
		return e
	}

	return o.v
}

func emptySlice[T any]() *s[T] {
	return &s[T]{}
}

func Slice[T any](v []T) *s[T] {
	return &s[T]{v}
}
