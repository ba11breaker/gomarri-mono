package framework

// IGroup is the interface of group
type IGroup interface {
	// Implement HttpMethods
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	// Implememt group nesting
	Group(string) IGroup
}

// Group structure implements IGroup interface
type Group struct {
	core   *Core  // Point to core structure
	parent *Group // Point to parent group
	prefix string // The prefix of group
}

// NewGroup creates a new group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}

// Implement Get method
func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Get(uri, handler)
}

// Implement Post method
func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Post(uri, handler)
}

// Implement Put method
func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Put(uri, handler)
}

// Implement Delete method
func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	g.core.Delete(uri, handler)
}

// Get the current group's absolute path
func (g *Group) getAbsolutePrefix() string {
	if g.parent != nil {
		return g.parent.getAbsolutePrefix() + g.prefix
	}
	return g.prefix
}

// Implement Group method
// Create a new group with the given uri
// The new group's prefix is the current group's prefix + the given uri
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}
