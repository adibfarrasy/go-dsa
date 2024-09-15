package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertNode(root.Left, val)
	} else if val > root.Val {
		root.Right = insertNode(root.Right, val)
	}
	return root
}

func ConstructBST(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[0].(int)}
	for i := 1; i < len(arr); i++ {
		if arr[i] != nil {
			insertNode(root, arr[i].(int))
		}
	}
	return root
}

func (t *TreeNode) GetByVal(v int) *TreeNode {
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val == v {
			return node
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return nil
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestor(root.Right, p, q)
	}

	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestor(root.Left, p, q)
	}

	return root
}
