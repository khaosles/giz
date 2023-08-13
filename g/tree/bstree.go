package tree

import (
	"math"

	"github.com/khaosles/giz/constraints"
	"github.com/khaosles/giz/g"
)

/*
   @File: bstree.go
   @Author: khaosles
   @Time: 2023/8/13 16:01
   @Desc:
*/

// BSTree is a binary search tree data structure in which each node has at most two children,
// which are referred to as the left child and the right child.
// In BSTree: leftNode < rootNode < rightNode
// type T should implements Compare function in lancetconstraints.Comparator interface.
type BSTree[T any] struct {
	root       *g.TreeNode[T]
	comparator constraints.Comparator[T]
}

// NewBSTree create a BSTree pointer
// param `comparator` is used to compare values in the tree
func NewBSTree[T any](rootData T, comparator constraints.Comparator[T]) *BSTree[T] {
	root := g.NewTreeNode(rootData)
	return &BSTree[T]{root, comparator}
}

// Insert insert data into BSTree
func (t *BSTree[T]) Insert(data T) {
	root := t.root
	newNode := g.NewTreeNode(data)
	if root == nil {
		t.root = newNode
	} else {
		insertTreeNode(root, newNode, t.comparator)
	}
}

// Delete delete data into BSTree
func (t *BSTree[T]) Delete(data T) {
	deleteTreeNode(t.root, data, t.comparator)
}

// NodeLevel get node level in BSTree
func (t *BSTree[T]) NodeLevel(node *g.TreeNode[T]) int {
	if node == nil {
		return 0
	}
	left := float64(t.NodeLevel(node.Left))
	right := float64(t.NodeLevel(node.Right))

	return int(math.Max(left, right)) + 1
}

// PreOrderTraverse traverse tree node in pre order
func (t *BSTree[T]) PreOrderTraverse() []T {
	return preOrderTraverse(t.root)
}

// PostOrderTraverse traverse tree node in post order
func (t *BSTree[T]) PostOrderTraverse() []T {
	return postOrderTraverse(t.root)
}

// InOrderTraverse traverse tree node in mid order
func (t *BSTree[T]) InOrderTraverse() []T {
	return inOrderTraverse(t.root)
}

// LevelOrderTraverse traverse tree node in level order
func (t *BSTree[T]) LevelOrderTraverse() []T {
	traversal := make([]T, 0)
	levelOrderTraverse(t.root, &traversal)
	return traversal
}

// Depth returns the calculated depth of a binary saerch tree
func (t *BSTree[T]) Depth() int {
	return calculateDepth(t.root, 0)
}

// HasSubTree checks if the tree `t` has `subTree` or not
func (t *BSTree[T]) HasSubTree(subTree *BSTree[T]) bool {
	return hasSubTree(t.root, subTree.root, t.comparator)
}

func hasSubTree[T any](superTreeRoot, subTreeRoot *g.TreeNode[T],
	comparator constraints.Comparator[T]) bool {
	result := false

	if superTreeRoot != nil && subTreeRoot != nil {
		if comparator.Compare(superTreeRoot.Value, subTreeRoot.Value) == 0 {
			result = isSubTree(superTreeRoot, subTreeRoot, comparator)
		}
		if !result {
			result = hasSubTree(superTreeRoot.Left, subTreeRoot, comparator)
		}
		if !result {
			result = hasSubTree(superTreeRoot.Right, subTreeRoot, comparator)
		}
	}
	return result
}

// Print the bstree structure
func (t *BSTree[T]) Print() {
	maxLevel := t.NodeLevel(t.root)
	nodes := []*g.TreeNode[T]{t.root}
	printTreeNodes(nodes, 1, maxLevel)
}
