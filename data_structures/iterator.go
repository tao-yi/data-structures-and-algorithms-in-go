package data_structures

type Iterator interface {
	HasNext() bool
	Next() int
}
