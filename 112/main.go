func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
			return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
			return true
	}
	if hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val) {
			return true
	}
	return false
}