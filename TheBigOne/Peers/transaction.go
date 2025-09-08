package peers

type Transaction struct {
	ID     string
	From   string
	To     string
	Amount int
}

func MakeTransaction(ID string, From string, To string, Amount int) *Transaction {
	t := new(Transaction)
	t.ID = ID
	t.From = From
	t.To = To
	t.Amount = Amount
	return t
}
