package disjointset

// Element is an element of a set.
type Element struct {
	parent *Element
	rank   int
	// The value stored within this element
	Value interface{}
}

// New creates the singleton {x}.
func New(x interface{}) *Element {
	e := new(Element)
	e.parent = e
	e.rank = 0
	e.Value = x
	return e
}

// Find returns the representative of the set containing x.
func Find(x *Element) *Element {
	if x != x.parent { // not the root
		x.parent = Find(x.parent)
	}
	return x.parent
}

// Union merges the two sets containing x and y.
// If x and y are already in the same set, Union does nothing.
func Union(x, y *Element) {
	xRoot, yRoot := Find(x), Find(y)
	if xRoot == yRoot { // in the same set already
		return
	}

	// x and y are not in the same set: merge
	if xRoot.rank < yRoot.rank {
		xRoot.parent = yRoot
	} else if yRoot.rank < xRoot.rank {
		yRoot.parent = xRoot
	} else {
		xRoot.parent = yRoot
		yRoot.rank++
	}
}
