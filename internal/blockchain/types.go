package blockchain

type BitIron struct {
	genesis string
	block   Block
}

type Block struct {
	id int64
}
