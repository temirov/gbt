package stack_test

import (
	"node"
	"stack"
	"testing"
)

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
	cases := []struct {
		in   stack.Stack
		push *node.Node
		want stack.Stack
	}{
		{
			stack.Stack([]*node.Node{&node.Node{Val: 1}}),
			&node.Node{Val: 2},
			stack.Stack([]*node.Node{&node.Node{Val: 1}, &node.Node{Val: 2}}),
		},
		{
			stack.Stack([]*node.Node{}),
			&node.Node{Val: 2},
			stack.Stack([]*node.Node{&node.Node{Val: 2}}),
		},
	}

	for _, c := range cases {
		got := c.in.Push(c.push)
		if !stack.Equal(c.want, got) {
			t.Errorf("Push(%v) == %v, want %v", c.push, got, c.want)
		}
	}
}
