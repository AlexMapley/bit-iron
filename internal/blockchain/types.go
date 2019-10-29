package blockchain

type BitIron struct {
	Genesis string
	Block   Block
}

// A Block represents a single transaction
// that may or may not have generated new BitIron
type Block struct {
	ID          int64
	Transaction Transaction
}
