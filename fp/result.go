package fp

type (
	Result[T any] struct {
		value T
		err   error
	}
)

var (
	OkAny = Result[any]{}
)

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) UnwrapUnsafe() T {
	if r.err != nil {
		panic("result is error: " + r.err.Error())
	}

	return r.value
}

func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.err
}

func (r Result[T]) UnwrapOr(other T) T {
	if r.err == nil {
		return r.value
	}

	return other
}

func (r Result[T]) UnwrapOrElse(fn func() T) T {
	if r.err == nil {
		return r.value
	}

	return fn()
}

func (r Result[T]) UnwrapOrDefault() T {
	if r.err == nil {
		return r.value
	}

	var res T
	return res
}

func (r Result[T]) Or(other Result[T]) Result[T] {
	if r.err == nil {
		return r
	}

	return other
}

func (r Result[T]) OrElse(fn func() Result[T]) Result[T] {
	if r.err == nil {
		return r
	}

	return fn()
}

func (r Result[T]) Match(
	handleOk func(T) Result[T],
	handleErr func(error) Result[T],
) Result[T] {
	if r.err == nil {
		return handleOk(r.value)
	}

	return handleErr(r.err)
}

func (r Result[T]) And(other Result[T]) Result[T] {
	if r.err == nil {
		return other
	}

	return r
}

func (r Result[T]) AndThen(fn func() T) Result[T] {
	if r.err == nil {
		return Ok(fn())
	}

	return r
}

func (r Result[T]) Map(fn func(T) T) Result[T] {
	if r.err == nil {
		return Ok(fn(r.value))
	}

	return r
}

func (r Result[T]) MapOr(value T, fn func(T) T) Result[T] {
	if r.err == nil {
		return Ok(fn(r.value))
	}

	return Ok(value)
}

func (r Result[T]) MapOrElse(
	handleErr func(error) T,
	handleOk func(T) T,
) Result[T] {
	if r.err == nil {
		return Ok(handleOk(r.value))
	}

	return Ok(handleErr(r.err))
}

func Ok[T any](v T) Result[T] {
	return Result[T]{value: v, err: nil}
}

func OkZero[T any]() Result[T] {
	return Result[T]{err: nil}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}
