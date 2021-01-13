package binarySearchTree

import "math"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// FindClosestValue
// Find closest value with target
// Idea: Compare root, left, right to find closest value.
//		If left or right is the closest value, check their
//		descendents too
func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	var previousClosest int
	if target - tree.Value > 0 {
		if tree.Right == nil {
			return tree.Value
		} else {
			previousClosest = tree.Right.FindClosestValue(target)
		}
	} else {
		if tree.Left == nil {
			return tree.Value
		} else {
			previousClosest = tree.Left.FindClosestValue(target)
		}
	}
	if Abs(target - previousClosest) > Abs(target - tree.Value) {
		return tree.Value
	} else {
		return previousClosest
	}
}

func (tree *BST) Insert(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.
	if tree.Value > value {
		if tree.Left != nil {
			tree.Left.Insert(value)
		} else {
			tree.Left = &BST{value, nil, nil}
		}
	} else {
		if tree.Right != nil {
			tree.Right.Insert(value)
		} else {
			tree.Right = &BST{value, nil, nil}
		}
	}
	return tree
}

//func (tree *BST) Insert(value int) *BST {
//	if value < tree.Value {
//		if tree.Left == nil {
//			tree.Left = &BST{Value: value}
//		} else {
//			tree.Left.Insert(value)
//		}
//	} else {
//		if tree.Right == nil {
//			tree.Right = &BST{Value: value}
//		} else {
//			tree.Right.Insert(value)
//		}
//	}
//	return tree
//}
//
//func (tree *BST) Contains(value int) bool{
//	if value < tree.Value {
//		if tree.Left == nil {
//			return false
//		} else {
//			return tree.Left.Contains(value)
//		}
//	} else if value > tree.Value {
//		if tree.Right == nil {
//			return false
//		} else {
//			return tree.Right.Contains(value)
//		}
//	}
//	return true
//}

func (tree *BST) Contains(value int) bool {
	// Write your code here.
	if tree.Value == value {
		return true
	} else {
		if tree.Value > value {
			if tree.Left != nil {
				return tree.Left.Contains(value)
			} else {
				return false
			}
		} else {
			if tree.Right != nil {
				return tree.Right.Contains(value)
			} else {
				return false
			}
		}
	}
}

func (tree *BST) Remove(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.
	// Find the node
	if tree.Value == value {
		// Find the predecessor of that node
		// Delete node and replace with predecessor
	} else {
		if tree.Value < value {
			if tree.Left != nil {
				tree.Left.Remove(value)
			}
		} else {
			if tree.Right != nil {
				tree.Right.Remove(value)
			}
		}
	}

	return tree
}

func (tree *BST) remove(value int, parent *BST) {
	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.remove(value, tree)
		}
	} else if value > tree.Value {
		if tree.Right != nil {
			tree.Right.remove(value, tree)
		}
	} else {
		if tree.Left != nil && tree.Right != nil {
			tree.Value = tree.Right.getMinValue()
			tree.Right.remove(tree.Value, tree)
		} else if parent == nil {
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Right = tree.Left.Right
				tree.Left = tree.Left.Left
			} else if tree.Right != nil {
				tree.Value = tree.Right.Value
				tree.Left = tree.Right.Left
				tree.Right = tree.Right.Right
			} else {

			}
		} else if parent.Left == tree {
			if tree.Left != nil {
				parent.Left = tree.Left
			} else {
				parent.Left = tree.Right
			}
		} else if parent.Right == tree {
			if tree.Right != nil {
				parent.Right = tree.Left
			} else {
				parent.Right = tree.Right
			}
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return  tree.Left.getMinValue()
}

func (tree *BST) Validate() bool {
	return tree.validate(math.MinInt32, math.MaxInt32)
}


func (tree *BST) validate(min, max int) bool {
	// Write your code here.
	if tree.Value < min || tree.Value >= max {
		return false
	}
	if tree.Left != nil && !tree.Left.validate(min, tree.Value) {
		return false
	}
	if tree.Right != nil && !tree.Right.validate(tree.Value, max) {
		return false
	}
	return true
}

func (tree *BST) InOrderTraverse(array []int) []int {
	// Write your code here.
	return nil
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	// Write your code here.
	return nil
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	// Write your code here.
	return nil
}