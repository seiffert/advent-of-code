package main

import (
	"fmt"
	"strconv"
)

func main() {
	r := NewRing("469217538")

	r = r.Play(100)
	fmt.Println("solution 1:", r.Find[1].Next.StringUntil(1))

	r = NewRing("469217538")
	r.AddUntil(1000000)

	r = r.Play(10000000)

	one := r.Find[1]
	first, second := one.Next.v, one.Next.Next.v
	fmt.Println("solution 2:", int64(first)*int64(second))
}

type Node struct {
	v    int
	Next *Node
	Prev *Node
	Find map[int]*Node
}

func NewRing(input string) *Node {
	var n *Node
	for _, raw := range input {
		d, _ := strconv.Atoi(string(raw))
		if n == nil {
			n = &Node{v: d, Find: map[int]*Node{}}
			n.Prev, n.Next, n.Find[d] = n, n, n
			continue
		}
		o := &Node{v: d, Next: n, Prev: n.Prev, Find: n.Find}
		n.Prev.Next, n.Prev, n.Find[d] = o, o, o
	}
	return n
}

func (n *Node) Play(rounds int) *Node {
	for ; rounds > 0; rounds-- {
		removed := n.RemoveNext(3)
		dst := n.v - 1
		for n.Find[dst] == nil {
			dst--
			if dst < n.Min() {
				dst = n.Max()
			}
		}
		n.AddAfter(dst, removed)
		n = n.Next
	}
	return n
}

func (n *Node) RemoveNext(i int) *Node {
	firstRemoved := n.Next
	firstRemoved.Prev, n.Find[firstRemoved.v] = nil, nil

	lastRemoved := firstRemoved
	for j := 0; j < i-1; j++ {
		lastRemoved = lastRemoved.Next
		n.Find[lastRemoved.v] = nil
	}
	n.Next = lastRemoved.Next
	n.Next.Prev = n
	lastRemoved.Next = nil
	return firstRemoved
}

func (n *Node) Min() int {
	for i := 1; i < 5; i++ {
		if n.Find[i] != nil {
			return n.Find[i].v
		}
	}
	return 0
}

func (n *Node) Max() int {
	max := n.v
	for d := range n.Find {
		if d > max {
			max = d
		}
	}
	return max
}

func (n *Node) AddAfter(d int, o *Node) *Node {
	dst := n.Find[d]
	next := dst.Next

	dst.Next, o.Prev, n.Find[o.v] = o, dst, o

	for o.Next != nil {
		o = o.Next
		n.Find[o.v] = o
	}

	o.Next, next.Prev = next, o
	return n
}

func (n *Node) AddUntil(i int) *Node {
	for j := n.Max() + 1; j <= i; j++ {
		o := &Node{v: j, Next: n, Prev: n.Prev, Find: n.Find}
		n.Prev.Next, n.Prev, n.Find[j] = o, o, o
	}
	return n
}

func (n *Node) String() string {
	return strconv.Itoa(n.v) + n.Next.StringUntil(n.v)
}

func (n *Node) StringUntil(d int) string {
	var result string
	for n.v != d {
		result += strconv.Itoa(n.v)
		n = n.Next
	}
	return result
}
