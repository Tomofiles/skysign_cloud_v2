package txmanager

// TransactionManager .
type TransactionManager interface {
	Do(func(tx Tx) error) error
}

// Tx .
type Tx = interface{}
