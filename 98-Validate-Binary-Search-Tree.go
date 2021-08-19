package main

func isValidBST(root *TreeNode) bool {
	var ints []int

	traverse(root, &ints)

	for i := 1; i < len(ints); i++ {
		if ints[i-1] >= ints[i] {
			return false
		}
	}

	return true
}

func traverse(n *TreeNode, arr *[]int) {
	if n == nil {
		return
	}

	traverse(n.Left, arr)
	*arr = append(*arr, n.Val)
	traverse(n.Right, arr)
}
