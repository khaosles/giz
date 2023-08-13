package tree

import (
	"fmt"
	"math"

	"github.com/khaosles/giz/constraints"
	"github.com/khaosles/giz/g"
)

/*
   @File: tree_inner.go
   @Author: khaosles
   @Time: 2023/8/13 16:01
   @Desc:
*/

func preOrderTraverse[T any](node *g.TreeNode[T]) []T {
	data := []T{}
	if node != nil {
		data = append(data, node.Value)
		data = append(data, preOrderTraverse(node.Left)...)
		data = append(data, preOrderTraverse(node.Right)...)
	}
	return data
}

func postOrderTraverse[T any](node *g.TreeNode[T]) []T {
	data := []T{}
	if node != nil {
		data = append(data, preOrderTraverse(node.Left)...)
		data = append(data, preOrderTraverse(node.Right)...)
		data = append(data, node.Value)
	}
	return data
}

func inOrderTraverse[T any](node *g.TreeNode[T]) []T {
	data := []T{}
	if node != nil {
		data = append(data, inOrderTraverse(node.Left)...)
		data = append(data, node.Value)
		data = append(data, inOrderTraverse(node.Right)...)
	}
	return data
}

// func preOrderPrint[T any](node *g.TreeNode[T]) {
// 	if node == nil {
// 		return
// 	}

// 	fmt.Printf("%v, ", node.Value)
// 	preOrderPrint(node.Left)
// 	preOrderPrint(node.Right)
// }

// func postOrderPrint[T any](node *g.TreeNode[T]) {
// 	if node == nil {
// 		return
// 	}

// 	postOrderPrint(node.Left)
// 	postOrderPrint(node.Right)
// 	fmt.Printf("%v, ", node.Value)
// }

// func inOrderPrint[T any](node *g.TreeNode[T]) {
// 	if node == nil {
// 		return
// 	}

// 	inOrderPrint(node.Left)
// 	fmt.Printf("%v, ", node.Value)
// 	inOrderPrint(node.Right)
// }

func levelOrderTraverse[T any](root *g.TreeNode[T], traversal *[]T) {
	var q []*g.TreeNode[T] // queue
	var n *g.TreeNode[T]   // temp node

	q = append(q, root)

	for len(q) != 0 {
		n, q = q[0], q[1:]
		*traversal = append(*traversal, n.Value)
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
}

func insertTreeNode[T any](rootNode, newNode *g.TreeNode[T], comparator constraints.Comparator[T]) {
	if comparator.Compare(newNode.Value, rootNode.Value) == -1 {
		if rootNode.Left == nil {
			rootNode.Left = newNode
		} else {
			insertTreeNode(rootNode.Left, newNode, comparator)
		}
	} else {
		if rootNode.Right == nil {
			rootNode.Right = newNode
		} else {
			insertTreeNode(rootNode.Right, newNode, comparator)
		}
	}
}

func deleteTreeNode[T any](node *g.TreeNode[T], data T, comparator constraints.Comparator[T]) *g.TreeNode[T] {
	if node == nil {
		return nil
	}
	if comparator.Compare(data, node.Value) == -1 {
		node.Left = deleteTreeNode(node.Left, data, comparator)
	} else if comparator.Compare(data, node.Value) == 1 {
		node.Right = deleteTreeNode(node.Right, data, comparator)
	} else {
		if node.Left == nil {
			node = node.Right
		} else if node.Right == nil {
			node = node.Left
		} else {
			l := node.Right
			d := inOrderSuccessor(l)
			d.Left = node.Left
			return node.Right
		}
	}

	return node
}

func inOrderSuccessor[T any](root *g.TreeNode[T]) *g.TreeNode[T] {
	cur := root
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

func printTreeNodes[T any](nodes []*g.TreeNode[T], level, maxLevel int) {
	if len(nodes) == 0 || isAllNil(nodes) {
		return
	}

	floor := maxLevel - level
	endgeLines := int(math.Pow(float64(2), (math.Max(float64(floor)-1, 0))))
	firstSpaces := int(math.Pow(float64(2), float64(floor))) - 1
	betweenSpaces := int(math.Pow(float64(2), float64(floor)+1)) - 1

	printSpaces(firstSpaces)

	newNodes := []*g.TreeNode[T]{}
	for _, node := range nodes {
		if node != nil {
			fmt.Printf("%v", node.Value)
			newNodes = append(newNodes, node.Left)
			newNodes = append(newNodes, node.Right)
		} else {
			newNodes = append(newNodes, nil)
			newNodes = append(newNodes, nil)
			printSpaces(1)
		}

		printSpaces(betweenSpaces)
	}

	fmt.Println("")

	for i := 1; i <= endgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			printSpaces(firstSpaces - i)
			if nodes[j] == nil {
				printSpaces(endgeLines + endgeLines + i + 1)
				continue
			}

			if nodes[j].Left != nil {
				fmt.Print("/")
			} else {
				printSpaces(1)
			}

			printSpaces(i + i - 1)

			if nodes[j].Right != nil {
				fmt.Print("\\")
			} else {
				printSpaces(1)
			}
			printSpaces(endgeLines + endgeLines - 1)
		}
		fmt.Println("")
	}

	printTreeNodes(newNodes, level+1, maxLevel)
}

// printSpaces
func printSpaces(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}

func isAllNil[T any](nodes []*g.TreeNode[T]) bool {
	for _, v := range nodes {
		if v != nil {
			return false
		}
	}
	return true
}

func calculateDepth[T any](node *g.TreeNode[T], depth int) int {
	if node == nil {
		return depth
	}
	return max(calculateDepth(node.Left, depth+1), calculateDepth(node.Right, depth+1))
}

func isSubTree[T any](superTreeRoot, subTreeRoot *g.TreeNode[T], comparator constraints.Comparator[T]) bool {
	if subTreeRoot == nil {
		return true
	}
	if superTreeRoot == nil {
		return false
	}
	if comparator.Compare(superTreeRoot.Value, subTreeRoot.Value) != 0 {
		return false
	}
	result := isSubTree(superTreeRoot.Left, subTreeRoot.Left, comparator) && isSubTree(superTreeRoot.Right, subTreeRoot.Right, comparator)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
