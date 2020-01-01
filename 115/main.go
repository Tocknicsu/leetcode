
func hasPathSum(root *TreeNode, sum int, path []int, ans *[][]int) {
	if root == nil {
			return
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
			path = append(path, root.Val)
			tmp := make([]int, len(path))
			copy(tmp, path)
			*ans = append(*ans, tmp)
			path = path[:len(path) - 1]
			return
	}
	path = append(path, root.Val)
	hasPathSum(root.Left, sum - root.Val, path, ans)
	hasPathSum(root.Right, sum - root.Val, path, ans)
	path = path[:len(path)-1]
}


func pathSum(root *TreeNode, sum int) [][]int {
	path := []int{}
	ans := [][]int{}
	hasPathSum(root, sum, path, &ans)
	return ans
}