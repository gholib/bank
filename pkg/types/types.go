package types

//Money
type Money int64

//Category
type Category string

//Currency currency code
type Currency string

//currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN card number
type PAN string

//Card payments card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

//Status payment status
type Status string

//payment statuses
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

//Payment payment information
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
