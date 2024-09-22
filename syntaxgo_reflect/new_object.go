package syntaxgo_reflect

// GetObject 能够在编译阶段预防T是指针类型（比如*A）的情况
func GetObject[T any]() (a T) {
	return a
}

func NewObject[T any]() (a *T) {
	return new(T)
}
