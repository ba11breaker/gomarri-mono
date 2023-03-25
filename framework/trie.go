package framework

import (
	"errors"
	"strings"
)

// Represent the Tree structure
type Tree struct {
	// The root node of the tree
	root *node
}

// Represent the node of the tree
type node struct {
	isLeaf   bool              // Is the leaf node
	segment  string            // The uri string
	handler  ControllerHandler // The controller handler
	children []*node           // The children of the node
}

func newNode() *node {
	return &node{
		isLeaf:   false,
		segment:  "",
		children: []*node{},
	}
}

func NewTree() *Tree {
	return &Tree{
		root: newNode(),
	}
}

// Check if the segment is a wild segment
// Start with ':'
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// Filter the child nodes by segment rules
func (n *node) filterChildNodes(segment string) []*node {
	// If the node's children is empty, return nil
	if len(n.children) == 0 {
		return nil
	}

	// If the segment is a wild segment, return all children
	if isWildSegment(segment) {
		return n.children
	}

	nodes := make([]*node, 0, len(n.children))
	for _, child := range n.children {
		// If the child's segment is a wild segment, append it
		if isWildSegment(child.segment) {
			nodes = append(nodes, child)
			//
		} else if child.segment == segment {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

func (n *node) matchNode(uri string) *node {

	segments := strings.SplitN(uri, "/", 2)

	segment := segments[0]

	// If the segment is not a wild segment, convert it to uppercase
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	// Find the children nodes by segment rules
	children := n.filterChildNodes(segment)

	if len(children) == 0 {
		return nil
	}

	// If there is only one segment, return the first leaf node
	if len(segments) == 1 {
		for _, child := range children {
			if child.isLeaf {
				return child
			}
		}
		return nil
	}

	// If there are two segments, find the node recursively
	for _, child := range children {
		if node := child.matchNode(segments[1]); node != nil {
			return node
		}
	}

	return nil
}

func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := tree.root

	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")

	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}

		isLast := index == len(segments)-1

		var objNode *node

		childrenNodes := n.filterChildNodes(segment)

		if len(childrenNodes) > 0 {
			for _, child := range childrenNodes {
				if child.segment == segment {
					objNode = child
					break
				}
			}
		}

		if objNode == nil {
			childNode := newNode()
			childNode.segment = segment
			if isLast {
				childNode.isLeaf = true
				childNode.handler = handler
			}
			n.children = append(n.children, childNode)
			objNode = childNode
		}

		n = objNode
	}

	return nil
}

func (tree *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}
