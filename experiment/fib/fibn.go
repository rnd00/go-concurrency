package fibn

import "fmt"

// linked list?

type fibn struct {
	Order int
	Count int
	Prev  *fibn
	Next  *fibn
}

func new() *fibn {
	return &fibn{
		Order: 1,
		Count: 1,
	}
}

func (f *fibn) createNext() *fibn {
	if f.Next != nil {
		return f.Next
	}
	next := &fibn{
		Order: f.nextOrdr(),
		Count: f.nextCnt(),
		Prev:  f,
	}
	f.Next = next
	return next
}

func (f *fibn) nextCnt() int {
	var nextCount int
	if f.Prev != nil {
		nextCount = f.Count + f.Prev.Count
	} else {
		nextCount = f.Count
	}
	return nextCount
}

func (f *fibn) nextOrdr() int {
	return f.Order + 1
}

func (f *fibn) printFromHere() {
	a := f
	for a.Next != nil {
		fmt.Printf("%d, ", a.Count)
		a = a.Next
	}
	fmt.Print("\n")
}
