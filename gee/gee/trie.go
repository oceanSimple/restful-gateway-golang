package gee

import "strings"

type node struct {
	pattern  string  // complete path
	part     string  // part of path
	children []*node // child node
	isWild   bool    // whether it is a precise match: it is true when part contains ':' or '*'
}

// matchChild: match the child node of the current node
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren: match all child nodes of the current node
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// Insert a path into the trie tree
func (n *node) insert(pattern string, parts []string, height int) {
	// The exit of the recursion
	if len(parts) == height { // if the height is equal to the length of the parts, the recursion ends
		n.pattern = pattern
		return
	}

	// Get the part of the path and match the child node
	part := parts[height]
	child := n.matchChild(part)
	if child == nil { // if the child node does not exist, create a new node
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	// Recursively insert the child node
	child.insert(pattern, parts, height+1)
}

// Search the path in the trie tree
func (n *node) search(parts []string, height int) *node {
	// The exit of the recursion
	// if this is the end of the parts or the part is a wildcard, return the current node
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
