package tree

/*
				A
		B				C
	D		E		F		G
			|0|1|2|3|4|5|6|
arr: [A,B,C,D,E,F,G]

left child: 	n*2 + 1
right child:  n*2 + 2
parent node:  (n-1) / 2
*/

type Heap interface {
	Insert(item interface{})
	Remove(item interface{}) bool
}

// type heap struct {
// }
