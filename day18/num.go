package day18

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

type item struct {
	n   *Node
	val int
}

type Node struct {
	parent      *Node
	left, right item
}

func NI(n *Node) item {
	return item{n: n}
}
func II(val int) item {
	return item{val: val}
}
func I(i interface{}, parent *Node) item {
	if n, ok := i.(*Node); ok {
		n.parent = parent
		return NI(n)
	} else {
		return II(i.(int))
	}
}

func N(l, r interface{}) *Node {
	n := &Node{}
	n.left, n.right = I(l, n), I(r, n)
	return n
}

func Parse(input string) (*Node, error) {
	b := bufio.NewReader(strings.NewReader(input))

	n := &Node{}
	err := n.parse(b)
	return n, err
}

func (n *Node) parse(b *bufio.Reader) error {
	r, _, err := b.ReadRune()
	if err != nil {
		return err
	} else if r != '[' {
		return fmt.Errorf("node must start with '[', got '%c'", r)
	}
	if err = n.left.parse(b, n); err != nil {
		return err
	}
	if r, _, err = b.ReadRune(); err != nil {
		return err
	} else if r != ',' {
		return fmt.Errorf("left must be followed by ',', got '%c'", r)
	}
	if err = n.right.parse(b, n); err != nil {
		return err
	}
	if r, _, err = b.ReadRune(); err != nil {
		return err
	} else if r != ']' {
		return fmt.Errorf("node must end with ']', got '%c'", r)
	}
	return nil
}

func (i *item) parse(b *bufio.Reader, parent *Node) error {
	r, _, err := b.ReadRune()
	if err != nil {
		return err
	}
	if err = b.UnreadRune(); err != nil {
		return err
	}
	if r == '[' {
		i.n = &Node{parent: parent}
		if err = i.n.parse(b); err != nil {
			return err
		}
	} else {
		if _, err = fmt.Fscanf(b, "%d", &i.val); err != nil {
			return err
		}
	}
	return nil
}

func (n *Node) String() string {
	b := strings.Builder{}
	n.WriteTo(&b)
	return b.String()
}

type RuneWriter interface {
	io.Writer
	WriteRune(rune) (int, error)
}

func (n *Node) WriteTo(w RuneWriter) {
	w.WriteRune('[')
	n.left.WriteTo(w)
	w.WriteRune(',')
	n.right.WriteTo(w)
	w.WriteRune(']')
}
func (i item) WriteTo(w RuneWriter) {
	if i.n != nil {
		i.n.WriteTo(w)
	} else {
		fmt.Fprintf(w, "%d", i.val)
	}
}

func (n *Node) Clone() *Node {
	if n == nil {
		return nil
	}
	ret := *n
	ret.left, ret.right = ret.left.Clone(), ret.right.Clone()
	return &ret
}
func (i item) Clone() item {
	i.n = i.n.Clone()
	return i
}

func add(l, r *Node) *Node {
	return &Node{left: NI(l.Clone()), right: NI(r.Clone())}
}

func (n *Node) reduce() {
	more := true
	for more {
		more = false
		if n.explodeWalk(0) {
			more = true
			continue
		}
		if n.splitWalk() {
			more = true
			continue
		}
	}
}

func Add(l, r *Node) *Node {
	s := add(l, r)
	s.reduce()
	return s
}

func (n *Node) explodeWalk(depth int) bool {
	if n == nil {
		return false
	}
	if depth >= 4 {
		n.explodeSelf()
		return true
	}
	depth++
	if n.left.explodeWalk(depth) {
		return true
	}
	if n.right.explodeWalk(depth) {
		return true
	}
	return false
}

func (i *item) explodeWalk(depth int) bool {
	if i.n == nil {
		return false
	}
	return i.n.explodeWalk(depth)
}

func (n *Node) splitWalk() bool {
	if n.left.splitWalk(n) {
		return true
	}
	if n.right.splitWalk(n) {
		return true
	}
	return false
}
func (i *item) splitWalk(container *Node) bool {
	if i.n != nil {
		return i.n.splitWalk()
	}
	if i.val >= 10 {
		half := float64(i.val) / 2
		i.n = N(int(math.Floor(half)), int(math.Ceil(half)))
		i.n.parent = container
		return true
	}
	return false
}

func (n *Node) explodeSelf() {
	n.parent.addToRightFrom(n, n.left.val)
	n.parent.addToLeftFrom(n, n.right.val)
	n.parent.zeroChild(n)
}

func (n *Node) addToRightFrom(c *Node, val int) {
	if n == nil {
		// root, we don't end up adding it anywhere
		return
	}
	if c == n.right.n {
		// the rightmost thing left of the right child is the rightmost thing in the left child
		n.left.addToRight(val)
	} else if c == n.left.n {
		// the rightmost thing left of the left child is the rightmost thing left of
		// this node in its parent
		n.parent.addToRightFrom(n, val)
	} else {
		panic(nil)
	}
}
func (n *Node) addToLeftFrom(c *Node, val int) {
	if n == nil {
		// root, we don't end up adding it anywhere
		return
	}
	if c == n.left.n {
		// the leftmost thing right of the left child is the leftmost thing in the right child
		n.right.addToLeft(val)
	} else if c == n.right.n {
		// the leftmost thing right of the right child is the leftmost thing right of
		// this node in its parent
		n.parent.addToLeftFrom(n, val)
	} else {
		panic(nil)
	}
}
func (i *item) addToRight(val int) {
	if i.n != nil {
		i.n.right.addToRight(val)
	} else {
		i.val += val
	}
}
func (i *item) addToLeft(val int) {
	if i.n != nil {
		i.n.left.addToLeft(val)
	} else {
		i.val += val
	}
}

func (n *Node) zeroChild(c *Node) {
	if c == n.left.n {
		n.left = item{}
	} else if c == n.right.n {
		n.right = item{}
	} else {
		panic(nil)
	}
}
