package tree

type node struct {
	value int
	left, right *node
}

type Tree struct {
	Root *node
}

func (t *Tree) Add(value int) bool {

	new_node := &node{value, nil, nil}

	if t.Root == nil {
		t.Root = new_node
		return true
	}

	return add_node(t.Root, new_node)
}

func add_node(subtree, new_node *node) bool {
	if new_node.value > subtree.value {
		if subtree.right == nil {
			subtree.right = new_node
			return true
		} else {
			return add_node(subtree.right, new_node)
		}
	} else if new_node.value < subtree.value {
		if subtree.left == nil {
			subtree.left = new_node
			return true
		} else {
			return add_node(subtree.left, new_node)
		}
	}

	return false
}