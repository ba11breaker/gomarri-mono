package framework

import (
	"testing"
)

func Test_filterChildNodes(t *testing.T) {
	root := &node{
		isLeaf:  false,
		segment: "",
		handler: func(*Context) error { return nil },
		children: []*node{
			{
				isLeaf:   true,
				segment:  "FOO",
				handler:  func(*Context) error { return nil },
				children: nil,
			},
			{
				isLeaf:   false,
				segment:  ":id",
				handler:  nil,
				children: nil,
			},
		},
	}

	{
		nodes := root.filterChildNodes("FOO")
		if len(nodes) != 2 {
			t.Error("foo error")
		}
	}

	{
		nodes := root.filterChildNodes(":foo")
		if len(nodes) != 2 {
			t.Error(":foo error")
		}
	}

	{
		nodes := root.filterChildNodes("BAR")
		if len(nodes) != 1 {
			t.Error("Bar error")
		}
	}

}

func Test_matchNode(t *testing.T) {
	root := &node{
		isLeaf:  false,
		segment: "",
		handler: func(*Context) error { return nil },
		children: []*node{
			{
				isLeaf:  true,
				segment: "FOO",
				handler: nil,
				children: []*node{
					{
						isLeaf:   true,
						segment:  "BAR",
						handler:  func(*Context) error { panic("not implemented") },
						children: []*node{},
					},
				},
			},
			{
				isLeaf:   true,
				segment:  ":id",
				handler:  nil,
				children: nil,
			},
		},
	}

	{
		node := root.matchNode("foo/bar")
		if node == nil {
			t.Error("match normal node error")
		}
	}

	{
		node := root.matchNode("test")
		if node == nil {
			t.Error("match test")
		}
	}

	{
		node := root.matchNode("foo/bar/test")
		if node != nil {
			t.Error("match foo/bar test")
		}
	}

}
