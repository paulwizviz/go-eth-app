package eth

// Transaction structure for storing transaction data
type Transaction struct {
	Hash  string
	From  string
	To    string
	Value string
	Block int
}

type Block struct {
	Number       string        `json:"number"`
	Transactions []Transaction `json:"transactions"`
}

type Parser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []Transaction
}
