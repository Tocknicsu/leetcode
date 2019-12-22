import "container/list"

func inorderTraversal(root *TreeNode) []int {
	st := list.New()
	res := []int{}

	st.PushBack(root)

	for st.Len() != 0 {
		now := st.Back().Value.(*TreeNode)
		st.Remove(st.Back())
		if now == nil {
			continue
		}
		if now.Left == nil && now.Right == nil {
			res = append(res, now.Val)
		} else {
			st.PushBack(now.Right)
			left := now.Left
			now.Left = nil
			now.Right = nil
			st.PushBack(now)
			st.PushBack(left)
		}
	}

	return res
}

// lazy to build tree by hand