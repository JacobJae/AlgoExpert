package binaryTrees

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value		int
	Left, Right	*BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	// Write your code here.
	return BranchSumsHelper(root, 0, []int{})
}

func BranchSumsHelper(root *BinaryTree, branchSum int, branchSumList []int) []int {
	// Write your code here.
	if root.Left == nil && root.Right == nil {
		return append(branchSumList, branchSum + root.Value)
	}
	var (
		left []int
		right []int
	)
	if root.Left != nil {
		left = BranchSumsHelper(root.Left, branchSum + root.Value, branchSumList)
	} else {
		left = []int{}
	}
	if root.Right != nil {
		right = BranchSumsHelper(root.Right, branchSum + root.Value, branchSumList)
	} else {
		right = []int{}
	}
	return append(left, right...)
}

func NodeDepths(root *BinaryTree) int {
	// Write your code here.
	NodeDepthsHelper(root, 0)
	return NodeDepthsHelper(root, 0)
}

func NodeDepthsHelper(root *BinaryTree, depth int) int {
	// Write your code here.
	depthSum := depth
	if root.Left != nil {
		depthSum += NodeDepthsHelper(root.Left, depth + 1)
	}
	if root.Right != nil {
		depthSum += NodeDepthsHelper(root.Right, depth + 1)
	}
	return depthSum
}