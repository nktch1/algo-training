package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var (
		res = [][]int{}
		q   = []*TreeNode{root}
	)

	for len(q) > 0 {

		var (
			l = len(q)
			level = []int{}
		)

		for i := 0; i < l; i++ {
			temp := q[0]
			q = q[1:]

			level = append(level, temp.Val)

			if temp.Left != nil {
				q = append(q, temp.Left)
			}

			if temp.Right != nil {
				q = append(q, temp.Right)
			}
		}

		res = append(res, level)
	}

	return res
}