package main

import (
	"AlgoExpert/arrays"
	"AlgoExpert/binarySearchTree"
	"fmt"
	"math/rand"
	"time"
)

func main () {
	//fmt.Println(arrays.TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
	//fmt.Println(arrays.IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}

	//nodeA := graphs.Node{"A", nil}
	//nodeB := graphs.Node{"B", nil}
	//nodeC := graphs.Node{"C", nil}
	//nodeD := graphs.Node{"D", nil}
	//nodeE := graphs.Node{"E", nil}
	//nodeF := graphs.Node{"F", nil}
	//nodeG := graphs.Node{"G", nil}
	//nodeH := graphs.Node{"H", nil}
	//nodeI := graphs.Node{"I", nil}
	//nodeJ := graphs.Node{"J", nil}
	//nodeK := graphs.Node{"K", nil}
	//nodeA.Children = []*graphs.Node{&nodeB, &nodeC, &nodeD}
	//nodeB.Children = []*graphs.Node{&nodeE, &nodeF}
	//nodeD.Children = []*graphs.Node{&nodeG, &nodeH}
	//nodeF.Children = []*graphs.Node{&nodeI, &nodeJ}
	//nodeG.Children =s []*graphs.Node{&nodeK}
	//
	//fmt.Println(nodeA.DepthFirstSearch([]string{}))
	//fmt.Println(recursion.RecursiveDoubling(2, 10))
	//
	//mat := recursion.FibMatrix(4)
	//fmt.Println(mat[0][0], mat[0][1], mat[1][0], mat[1][1])
	//fmt.Println(searching.FindThreeLargestNumbers([]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}))
	//fmt.Println(sorting.BubbleSort([]int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}))
	//fmt.Println(sorting.InsertionSort([]int{5, 2, 4, 6, 1, 3, 141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}))
	//fmt.Println(strings.CaesarCipherEncryptor("hyamz", 2))
	//fmt.Println(strings.RunLengthEncoding("AAAAAAAAAAAAABBCCCCDD"))
	//fmt.Println(strings.RunLengthEncoding("........______=========AAAA   AAABBBB   BBB"))
	//fmt.Println(arrays.ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
	//fmt.Println(arrays.MoveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2))
	//arr1 := []int{1,	2,	3,	4}
	//arr2 := []int{12,	13,	14,	5}
	//arr3 := []int{11,	16,	15,	6}
	//arr4 := []int{10,	9,	8,	7}
	//arr := [][]int{arr1, arr2, arr3, arr4}
	//fmt.Println(arrays.SpiralTraverse(arr))
	//fmt.Println(arrays.SpiralTraverse([][]int{{1, 2, 3, 4}, {10, 11, 12, 5}, {9, 8, 7, 6}}))
	fmt.Println(arrays.LongestPeak([]int{10, 4, 7, 11, 12, 6, 8, 3, 1, 2, 0, 9, 5}))
	fmt.Println(arrays.FirstDuplicateValue([]int{2, 3, 16, 9, 11, 14, 13, 1, 10, 12, 5, 17, 4, 16, 10, 5, 4}))
	bstRoot := binarySearchTree.BST{Value: 10}
	bstRoot.Insert(5)
	bstRoot.Insert(2)
	bstRoot.Insert(5)
	bstRoot.Insert(1)
	bstRoot.Insert(15)
	bstRoot.Insert(22)

}

func randomArray() []int {
	//Provide seed
	rand.Seed(time.Now().Unix())
	//Generate a random array of length n
	return rand.Perm(rand.Intn(20))
}