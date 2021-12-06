package tree

import (
	"errors"
	"sort"
)

// Record has and ID and the ID of a parent
type Record struct {
	ID     int
	Parent int
}

// Record collection for Sorting Interface
type Records []Record

// Len for Sortint Records
func (rs Records) Len() int {
	return len(rs)
}

// Less for Sorting Records
func (r Records) Less(i, j int) bool {
	return r[i].ID < r[j].ID
}

// Swap for Swapping Records in a collection
func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Define the Node type with an ID and collection of children Nodes
type Node struct {
	ID       int
	Children []*Node
}

// BFS for a node with ID
func (n Node) getNode(id int) *Node {
	nodes := []*Node{&n}
	//nodes = append(nodes, n)

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = append(nodes[1:], node.Children...)
		if id == (*node).ID {
			return node
		}
	}

	return nil
}

func isValidRoot(root Record) bool {
	if root.Parent != 0 || root.ID != 0 {
		return false
	}

	return true
}

func Build(records []Record) (*Node, error) {
	idSet := make(map[int]bool, 0)
	if len(records) == 0 {
		return nil, nil
	}

	// Sort records by ID
	sort.Sort(Records(records))

	// Set root node and paren
	root := Node{
		ID: records[0].ID,
	}
	parent := &root
	if !isValidRoot(records[0]) {
		return nil, errors.New("Root has a parent")
	}
	idSet[records[0].ID] = true

	for i := 1; i < len(records); i++ {
		if records[i].Parent == records[i].ID || records[i].Parent > records[i].ID {
			return nil, errors.New("Cycle detected")
		}

		// Contigous Checks
		if records[i].ID != (records[i-1].ID + 1) {
			return nil, errors.New("Records not contigous")
		}

		// Get the Parent of the node
		currentParentID := records[i].Parent
		if currentParentID != parent.ID {
			parent = root.getNode(currentParentID)
		}

		// Init the children for a new parent
		if parent.Children == nil {
			parent.Children = make([]*Node, 0)
		}

		if isDuplicate := idSet[records[i].ID]; isDuplicate {
			return nil, errors.New("Duplicate ID")
		}

		idSet[records[i].ID] = true

		// add the the record the the parent
		parent.Children = append(parent.Children, &(Node{ID: records[i].ID}))
	}

	return &root, nil
}
