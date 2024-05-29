package result

type Result[T any] struct {
	value T
	err   error
}

func New[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.err
}
