package data_structures

// A bag is a collection where removing items is not supported
// Its purpose is to provide clients with the ability to collect items
// and then to iterate through the collected items
// ** The order of iteration is unspecified and should be immaterial to the client
type Bag interface {
	Add(interface{})
	IsEmpty() bool
	Size() int
	Values() []interface{}
}

type bag struct {
	items []interface{}
}

func NewBag() Bag {
	return &bag{}
}

func (b *bag) Add(item interface{}) {
	b.items = append(b.items, item)
}

func (b *bag) IsEmpty() bool {
	return len(b.items) == 0
}

func (b *bag) Size() int {
	return len(b.items)
}

func (b *bag) Values() []interface{} {
	return b.items
}
