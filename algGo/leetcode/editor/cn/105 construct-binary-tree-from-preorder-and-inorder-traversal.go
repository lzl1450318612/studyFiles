package main

//给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并
//返回其根节点。 
//
// 示例 1: 
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//
// 示例 2: 
//输入: preorder = [-1], inorder = [-1]
//输出: [-1]

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStartIdx, preEndIdx int, inorder []int, inStartIdx, inEndIdx int) *TreeNode {
	if preStartIdx > preEndIdx {
		return nil
	}

	rootVal := preorder[preStartIdx]

	inorderRootIdx := 0

	for i := inStartIdx; i <= inEndIdx; i++ {
		if inorder[i] == rootVal {
			inorderRootIdx = i
			break
		}
	}

	root := &TreeNode{
		Val: rootVal,
	}

	root.Left = build(preorder, preStartIdx+1, preStartIdx+inorderRootIdx-inStartIdx, inorder, inStartIdx, inorderRootIdx-1)
	root.Right = build(preorder, preStartIdx+inorderRootIdx-inStartIdx+1, preEndIdx, inorder, inorderRootIdx+1, inEndIdx)

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
