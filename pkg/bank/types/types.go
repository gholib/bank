package types

//Money
type Money int64

//Category
type Category string

//Payment payment information
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
