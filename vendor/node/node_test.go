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
		{
			&node.Node{Val: 1},
			[]*node.Node{},
		},
		{
			&node.Node{Val: 1, Left: &node.Node{Val: 2}},
			[]*node.Node{&node.Node{Val: 2}},
		},
		{
			&node.Node{Val: 1, Right: &node.Node{Val: 2}},
			[]*node.Node{&node.Node{Val: 2}},
		},
		{
			&node.Node{Val: 1, Left: &node.Node{Val: 2}, Right: &node.Node{Val: 3}},
			[]*node.Node{&node.Node{Val: 2}, &node.Node{Val: 3}},
		},
	}

	for _, c := range cases {
		got := c.in.Children()
		if !node.Equal(c.want, got) {
			t.Errorf("Children(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestHasChildren(t *testing.T) {
	cases := []struct {
		in   *node.Node
		want bool
	}{
		{&node.Node{Val: 1}, false},
		{&node.Node{Val: 1, Left: &node.Node{Val: 2}}, true},
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
		{&node.Node{Val: 1}, true},
		{&node.Node{Val: 1, Left: &node.Node{Val: 2}}, false},
	}

	for _, c := range cases {
		got := c.in.NoChildren()
		if got != c.want {
			t.Errorf("NoChildren(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestString(t *testing.T) {
	cases := []struct {
		in   *node.Node
		want string
	}{
		{&node.Node{Val: 1}, "1"},
	}

	for _, c := range cases {
		got := c.in.String()
		if got != c.want {
			t.Errorf("String(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		in   []*node.Node
		want bool
	}{
		{[]*node.Node{&node.Node{Val: 2}, &node.Node{Val: 2}}, true},
		{[]*node.Node{&node.Node{Val: 1}, &node.Node{Val: 2}}, false},
	}

	for _, c := range cases {
		got := c.in[0].Equal(c.in[1])
		if got != c.want {
			t.Errorf("Equal(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
