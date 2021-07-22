package txmanager

// TransactionManager .
type TransactionManager interface {
	Do(operation func(Tx) error) error
	DoAndEndHook(operation func(Tx) error, endHook func() error) error
}

// Tx .
type Tx = interface{}
