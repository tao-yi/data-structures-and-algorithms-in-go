package tree

type Comparator interface {
	Compare(interface{}) int
}
