package typeutil

// Number limits balance to either int or float64
type Number interface {
	int | float64
}
