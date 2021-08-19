package main

// in   [3,9,20,15,7] - показывает рут, в первом элементе
// pre  [9,3,15,20,7] - показывает левое и правое поддерево относительно рута
// post [7,15,20,9,3]

// 1) взять рут
// 2) найти рут в preorder
// 3) то что слева - левое поддерево, справа - правое (в preorder)
// 4) left  = buildTree( [9], [9])
// 5) right = buildTree([20], [20, 15, 7])

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	mid := indexOrMinusOne(inorder, preorder[0])
	if mid == -1 {
		return nil
	}

	root.Left  = buildTree(
		preorder[1:1+mid],
		inorder[:mid],
	)

	root.Right = buildTree(
		preorder[mid+1:],
		inorder[mid+1:],
	)

	return root
}

func indexOrMinusOne(slice []int, value int) int {
	idx := -1

	for i := range slice {
		if slice[i] == value {
			idx = i
			break
		}
	}

	return idx
}
