package optional

type i struct {
	v any
}

func (o *i) Get() any {
	return o.v
}

func (o *i) Empty() bool {
	return o.v == nil
}

func (o *i) IfEmpty(fn func(any)) {
	if !o.Empty() {
		return
	}

	fn(o.v)
}

func (o *i) Present() bool {
	return !o.Empty()
}

func (o *i) IfPresent(fn func(any)) {
	if o.Empty() {
		return
	}

	fn(o.v)
}

func (o *i) If(fn func(any) bool) *i {
	if o.Empty() || fn(o.v) {
		return o
	}

	return emptyInterface()
}

func (o *i) Else(e any) any {
	if o.Empty() {
		return e
	}

	return o.v
}

func emptyInterface() *i {
	return &i{}
}

func Interface(v any) *i {
	return &i{v}
}

func Error(v any) *i {
	return &i{v}
}
