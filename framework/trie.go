package framework

// Represent the Tree structure
type Tree struct {
	// The root node of the tree
	root *node
}

// Represent the node of the tree
type node struct {
	isLeaf   bool              // Is the leaf node
	segement string            // The uri string
	handler  ControllerHandler // The controller handler
	children []*node           // The children of the node
}
