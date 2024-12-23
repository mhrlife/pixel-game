package pt

func Value[T any](val T) *T {
	return &val
}
