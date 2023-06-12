// Package ptr is a package to get a pointer for given types
package ptr

// New returns a pointer of type T
func New[T any](t T) *T {
	return &t
}
