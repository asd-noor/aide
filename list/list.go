package list

type transformFn[T any, U any] func(T) U
type applyFn[T any] func(T) T
type predicateFn[T any] func(T) bool
type foldFn[T any] func(T, T) T

// Convert takes in a slice and transformer function, and
// returns a different slice by applying the tranformer function on each element.
func Convert[T any, U any](arr []T, fn transformFn[T, U]) (returnValue []U) {
	for _, v := range arr {
		returnValue = append(returnValue, fn(v))
	}

	return
}

// Map takes in a slice and a function to apply on every element of slice
// and return the updated slice.
func Map[T any](arr []T, fn applyFn[T]) (returnValue []T) {
	for _, v := range arr {
		returnValue = append(returnValue, fn(v))
	}

	return
}

// Filter takes in a slice and a predicate function to return a filtered slice
// based on the condition in predicate function.
func Filter[T any](arr []T, fn predicateFn[T]) (returnValue []T) {
	for _, v := range arr {
		if fn(v) {
			returnValue = append(returnValue, v)
		}
	}

	return
}

// Fold reduces a slice to a single value based on given function to merge elements.
func Fold[T any](arr []T, fn foldFn[T]) (returnValue T) {
	arrLen := len(arr)
	if arrLen == 0 {
		return
	}

	returnValue = arr[0]
	if arrLen == 1 {
		return
	}

	for i := 1; i < arrLen-1; i++ {
		returnValue = fn(returnValue, arr[i])
	}

	return
}
