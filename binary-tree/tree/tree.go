package tree

import "fmt"

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

func (t *Tree) Search(value int) bool {
	if t.Root == nil {
		return false
	}

	target, _ := search_node(t.Root, value)

	if target != nil {
		return true
	} else {
		return false
	}
}

func search_node(n *node, value int) (*node, *node) {
	var father *node

	// busca pelo nó
	for {
		if n == nil {
			return nil, nil
		} else if n.value == value {
			break
		} else {
			father = n
			if value > n.value {
				n = n.right
			} else {
				n = n.left
			}
		}
	}

	return n, father
}

func (t *Tree) PrintInOrder() {
	if t == nil {
		return
	}

	print_in_order(t.Root, 0)
}

func print_in_order(n *node, indent int) {
	if n == nil {
		return
	}

	print_in_order(n.right, indent + 8)

	if indent == 0 {
		fmt.Printf("%*s", 2, "")
	} else {
		fmt.Printf("%*s", indent, "")
		fmt.Print("---| ")
	}

	fmt.Println(n.value)	

	print_in_order(n.left, indent + 8)
	
}

func (t *Tree) Remove(value int) bool {
	if t.Root == nil {
		return false
	}

	return remove_node(t, value)
}

func remove_node(t *Tree, value int) bool {
	n, father := search_node(t.Root, value)
	if n == nil {
		return false
	}

	var aux_node, aux_father *node

	if n.left != nil {
		aux_node, aux_father = get_right_most(n.left)
		if aux_node != n.left {
			aux_father.right = nil
			aux_node.left = n.left
		}
		aux_node.right = n.right
	} else {
		aux_node = n.right
	}

	if father == nil {
		t.Root = aux_node
	} else if father.left == n {
		father.left = aux_node
	} else {
		father.right = aux_node
	}

	return true
}

func get_right_most(n *node) (*node, *node) {
	var aux_father *node
	
	for n.right != nil {
		aux_father = n
		n = n.right
	}

	return n, aux_father
}

func (t *Tree) RotateRight() {
	if t.Root == nil || t.Root.left == nil{
		return
	}

	save_left := t.Root.left
	t.Root.left = t.Root.left.right
	save_left.right = t.Root
	t.Root = save_left
}