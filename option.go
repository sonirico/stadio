package stadio

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
func (o Option[T]) Unwrap() T {
	if !o.isSome {
		panic("option is none")
	}
	return o.value
}

func Some[T any](t T) Option[T] {
	return Option[T]{value: t, isSome: true}
}

func None[T any]() Option[T] {
	return Option[T]{}
}
