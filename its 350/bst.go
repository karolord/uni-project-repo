package main

type bst struct {
	x     int
	left  *bst
	right *bst
}

var head *bst
var tail *bst

func (b *bst) insertbst(x int) {
	if x < b.x {
		if b.left == nil {
			b.left = &bst{x: x}
		} else {
			b.left.insertbst(x)
		}
	} else if x > b.x {
		if b.right == nil {
			b.right = &bst{x: x}
		} else {
			b.right.insertbst(x)
		}
	}

}

func (b *bst) bsttodll() {
	if b != nil {

		if b.left != nil {
			b.left.bsttodll()
		}
		if tail != nil {
			tail.right = b
			b.left = tail
		} else {
			head = b
		}
	}
	tail = b
	if b.right != nil {
		b.right.bsttodll()
	}

}
func main() {
	binarysearch := &bst{x: 5}
	binarysearch.insertbst(6)
	binarysearch.insertbst(2)
	binarysearch.insertbst(4)
	binarysearch.insertbst(3)
	binarysearch.insertbst(1)
	binarysearch.insertbst(7)
	binarysearch.insertbst(8)
	binarysearch.bsttodll()

}
