package stack_test

import (
	"node"
	"stack"
	"testing"
)

func TestPopPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	stack.Stack([]*node.Node{}).Pop()
}

func TestPop(t *testing.T) {
	cases := []struct {
		in   stack.Stack
		want stack.Stack
	}{
		{
			stack.Stack([]*node.Node{&node.Node{Val: 1}}),
			stack.Stack([]*node.Node{}),
		},
	}

	for _, c := range cases {
		got, _ := c.in.Pop()
		if !stack.Equal(c.want, got) {
			t.Errorf("Pop() == %v, want %v", got, c.want)
		}
	}
}

func TestPush(t *testing.T) {
	var emptyStack stack.Stack
	cases := []struct {
		in struct {
			s    stack.Stack
			push node.Node
		}
		want struct {
			s stack.Stack
			l int
		}
	}{
		{
			in: struct {
				s    stack.Stack
				push node.Node
			}{
				stack.Stack([]*node.Node{&node.Node{Val: 1}}),
				node.Node{Val: 2},
			},
			want: struct {
				s stack.Stack
				l int
			}{
				stack.Stack([]*node.Node{&node.Node{Val: 1}, &node.Node{Val: 2}}),
				2,
			},
		},
		{
			in: struct {
				s    stack.Stack
				push node.Node
			}{
				stack.Stack([]*node.Node{}),
				node.Node{Val: 2},
			},
			want: struct {
				s stack.Stack
				l int
			}{
				stack.Stack([]*node.Node{&node.Node{Val: 2}}),
				1,
			},
		},
		{
			in: struct {
				s    stack.Stack
				push node.Node
			}{
				emptyStack,
				node.Node{Val: 1},
			},
			want: struct {
				s stack.Stack
				l int
			}{
				stack.Stack([]*node.Node{&node.Node{Val: 1}}),
				1,
			},
		},
	}

	for _, c := range cases {
		got := c.in.s.Push(&c.in.push)
		if len(got) != c.want.l {
			t.Errorf("len(%v) == %v, want %v", got, len(got), c.want.l)
		}
		if !stack.Equal(c.want.s, got) {
			t.Errorf("Push(%v) == %v, want %v", c.in.push, got, c.want.s)
		}
	}
}

func recursion(n node.Node, s *stack.Stack) {
	switch {
	case n.Val > 0:
		*s = (*s).Push(&n)
	default:
		return
	}

	n = node.Node{Val: n.Val - 1}
	recursion(n, s)
}

func TestPushRecursion(t *testing.T) {
	node := node.Node{Val: 10}
	var st stack.Stack
	recursion(node, &st)
	got := len(st)
	if got != 10 {
		t.Errorf("len(%v) == %v, want %v", st, got, 10)
	}
}
