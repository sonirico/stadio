package fp

type (
	Option[T any] struct {
		value  T
		isSome bool
	}
)

func (o Option[T]) IsSome() bool {
	return o.isSome
}

func (o Option[T]) IsNone() bool {
	return !o.isSome
}

func (o Option[T]) Unwrap() (T, bool) {
	return o.value, o.isSome
}

func (o Option[T]) UnwrapOr(value T) T {
	if o.isSome {
		return o.value
	}
	return value
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	if o.isSome {
		return o.value
	}
	return fn()
}

func (o Option[T]) UnwrapOrDefault() T {
	if o.isSome {
		return o.value
	}
	var res T
	return res
}

func (o Option[T]) UnwrapUnsafe() T {
	if !o.isSome {
		panic("option is none")
	}
	return o.value
}

func (o Option[T]) Or(other Option[T]) Option[T] {
	if !o.isSome {
		return other
	}
	return o
}

func (o Option[T]) OrElse(fn func() Option[T]) Option[T] {
	if !o.isSome {
		return fn()
	}
	return o
}

func (o Option[T]) Map(fn func(T) T) Option[T] {
	if o.isSome {
		return Some(fn(o.value))
	}
	return o
}

func (o Option[T]) MapOr(value T, fn func(T) T) T {
	if o.isSome {
		return fn(o.value)
	}
	return value
}

func (o Option[T]) MapOrElse(handleNone func() T, handleSome func(T) T) T {
	if o.isSome {
		return handleSome(o.value)
	}
	return handleNone()
}

func (o Option[T]) OkOr(err error) Result[T] {
	if o.isSome {
		return Ok(o.value)
	}
	return Err[T](err)
}

func (o Option[T]) OkOrElse(fn func() error) Result[T] {
	if o.isSome {
		return Ok(o.value)
	}
	return Err[T](fn())
}

func (o Option[T]) Match(handleSome func(T) Option[T], handleNone func() Option[T]) Option[T] {
	if o.isSome {
		return handleSome(o.value)
	}

	return handleNone()
}

func Some[T any](t T) Option[T] {
	return Option[T]{value: t, isSome: true}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func OptionFromTuple[T any](x T, ok bool) Option[T] {
	if ok {
		return Some(x)
	}
	return None[T]()
}

func OptionFromPtr[T any](x *T) Option[T] {
	if x == nil {
		return None[T]()
	}
	return Some(*x)
}

func OptionFromZero[T comparable](x T) Option[T] {
	var zero T
	if x == zero {
		return None[T]()
	}
	return Some(x)
}
