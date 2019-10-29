package transactions

// A Transaction represents an exchange of BitIron
// between a sender and a receiver
type Transaction struct {
	SenderAddress   string
	ReceiverAddress string
	Amount          int64
}
