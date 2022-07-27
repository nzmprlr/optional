package optional

type m[K comparable, V any] struct {
	v map[K]V
}

func (o *m[K, V]) Get() map[K]V {
	return o.v
}

func (o *m[K, V]) Empty() bool {
	return len(o.v) == 0
}

func (o *m[K, V]) IfEmpty(fn func(map[K]V)) {
	if !o.Empty() {
		return
	}

	fn(o.v)
}

func (o *m[K, V]) Present() bool {
	return !o.Empty()
}

func (o *m[K, V]) IfPresent(fn func(map[K]V)) {
	if o.Empty() {
		return
	}

	fn(o.v)
}

func (o *m[K, V]) If(fn func(map[K]V) bool) *m[K, V] {
	if o.Empty() || fn(o.v) {
		return o
	}

	return emptyMap[K, V]()
}

func (o *m[K, V]) Else(e map[K]V) map[K]V {
	if o.Empty() {
		return e
	}

	return o.v
}

func emptyMap[K comparable, V any]() *m[K, V] {
	return &m[K, V]{}
}

func Map[K comparable, V any](v map[K]V) *m[K, V] {
	return &m[K, V]{v}
}
