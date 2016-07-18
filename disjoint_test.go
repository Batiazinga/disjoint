package disjointsets

import (
	"fmt"
	"testing"
)

func Example() {
	// build three singletons
	a := MakeSet("a")
	b := MakeSet("b")
	c := MakeSet("c")

	// union a and b
	Union(a, b)

	// a and b are in the same set, c is in another set
	fmt.Printf("a and b are in the same set: %v\nc is in another set: %v\n", Find(a) == Find(b), Find(a) != Find(c))
	// Output:
	// a and b are in the same set: true
	// c is in another set: true
}

func TestDisjoint(t *testing.T) {
	sets := make(map[int]*Element)
	sets[0] = MakeSet(0)
	sets[1] = MakeSet(1)
	sets[2] = MakeSet(2)

	for i, elt := range sets {
		if elt.parent != elt {
			t.Errorf("Element %v should be the Root.\n", i)
		}
		if Find(elt) != elt {
			t.Errorf("Wrong returned value for Find: %v should be its own parent (root).\n", i)
		}
		if elt.rank != 0 {
			t.Errorf("Rank of %v should be zero.\n", i)
		}
	}

	if Find(sets[0]) == Find(sets[1]) {
		t.Error("0 and 1 should be in different sets.\n")
	}
	Union(sets[0], sets[1])
	if Find(sets[0]) != Find(sets[1]) {
		t.Error("0 and 1 should be in the same set.\n")
	}
	if Find(sets[0]) == Find(sets[2]) {
		t.Error("0 and 2 should be in different sets.\n")
	}
	Union(sets[0], sets[2])
	if Find(sets[0]) != Find(sets[2]) {
		t.Error("0 and 2 should be in the same set.\n")
	}
}
