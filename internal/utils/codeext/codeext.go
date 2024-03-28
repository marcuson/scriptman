package codeext

func Tern[T any](cond bool, a T, b T) T {
	if cond {
		return a
	}

	return b
}

func SetIf[T any](prop *T, cond bool, value T) {
	if cond {
		*prop = value
	}
}
