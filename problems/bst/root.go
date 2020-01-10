// bst stores data in a binary search tree
package bst

// Node of a binary search tree
type Node struct {
	Priority int
	Left     *Node
	Right    *Node
}

// Insert will insert into tree given new and root
func Insert(root, newNode *Node) {
	if root == nil {
		return
	}

	evaluator := root
	for {
		if newNode.Priority <= evaluator.Priority {
			if evaluator.Left == nil {
				evaluator.Left = newNode
				return
			}
			evaluator = evaluator.Left
			continue
		}

		if evaluator.Right == nil {
			evaluator.Right = newNode
			return
		}
		evaluator = evaluator.Right
	}
}

// Search will search for the node with the given priority
func Search(root *Node, target int) *Node {

	if root == nil {
		return nil
	}

	evaluator := root
	for {
		if target == evaluator.Priority {
			return evaluator
		}

		if target < evaluator.Priority {
			if evaluator.Left == nil {
				return nil
			}

			evaluator = evaluator.Left
			continue
		}

		if evaluator.Right == nil {
			return nil
		}

		evaluator = evaluator.Right
		continue
	}
}

func BFS(root *Node, target int) *Node {
	queue := []*Node{root}

	for len(queue) != 0 {
		evaluator := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]
		if evaluator.Priority == target {
			return evaluator
		}

		if evaluator.Left != nil {
			queue = append(queue, evaluator.Left)
		}
		if evaluator.Right != nil {
			queue = append(queue, evaluator.Right)
		}
	}
	return nil
}

/*
  //Pop a vertex from stack to visit next
            	v  =  S.top( )
	    	S.pop( )
      		//Push all the neighbours of v in stack that are not visited
		//for all neighbours w of v in Graph G:
					  if w is not visited :
							       S.push( w )
										   mark w as visited
*/

// Search a tree of nodes in
func DFSInorder(root *Node, target int) *Node {
	stack := []*Node{root}
	visited := make(map[*Node]bool)

	for len(stack) != 0 {
		evaluator := stack[len(stack)-1]
		if evaluator.Priority == target {
			return evaluator
		}

		if evaluator.Left != nil && !visited[evaluator.Left] {
			stack = append(stack, evaluator)
			continue
		}

		if evaluator.Right != nil && !visited[evaluator.Right] {
			stack = append(stack, evaluator)
			continue
		}
		visited[evaluator] = true
		stack = stack[0 : len(stack)-1]
	}
	return nil
}
