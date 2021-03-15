// Reader ...
type Reader interface{}

// Writer ...
type Writer interface {
	CreateNumber(ctx context.Context, n int) (int, error)
}

// Repository holds the set of method signatures
type Repository interface {
	Reader
	Writer
}