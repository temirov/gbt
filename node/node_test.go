package node_test

import (
	"node"
	"testing"
)

func TestChildren(t *testing.T) {
	cases := []struct {
		in   *node.Node
		want []*node.Node
	}{
		{&node.Node{val: 1}, []*Node{}},
		{&node.Node{val: 1, left: &node.Node{val: 2}}, []*Node{&node.Node{val: 2}}},
	}

	for _, c := range cases {
		got := c.in.Children()
		if got != c.want {
			t.Errorf("Children(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestHasChildren(t *testing.T) {
	cases := []struct {
		in   *node.Node
		want bool
	}{
		{&node.Node{val: 1}, false},
		{&node.Node{val: 1, left: &node.Node{val: 2}}, true},
	}

	for _, c := range cases {
		got := c.in.HasChildren()
		if got != c.want {
			t.Errorf("HasChildren(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestNoChildren(t *testing.T) {
	cases := []struct {
		in   *node.Node
		want bool
	}{
		{&node.Node{val: 1}, true},
		{&node.Node{val: 1, left: &node.Node{val: 2}}, false},
	}

	for _, c := range cases {
		got := c.in.NoChildren()
		if got != c.want {
			t.Errorf("NoChildren(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
