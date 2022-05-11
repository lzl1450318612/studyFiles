package main

import (
	"fmt"
)

//给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节
//点为空。 
// 每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
//
// 示例 1: 
//输入:
//           1
//         /   \
//        3     2
//       / \     \  
//      5   3     9 
//
//输出: 4
//解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
//
// 示例 2: 
//输入:
//          1
//         /  
//        3    
//       / \       
//      5   3     
//
//输出: 2
//解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。
//
// 示例 3: 
//输入:
//          1
//         / \
//        3   2 
//       /        
//      5      
//
//输出: 2
//解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。
//
// 示例 4: 
//输入:
//          1
//         / \
//        3   2
//       /     \  
//      5       9 
//     /         \
//    6           7
//输出: 8
//解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
func widthOfBinaryTree(root *TreeNode) int {
	levelArr := make([]int, 0)
	a, _ := getMaxWidth(root, 1, 1, levelArr, 1)
	return a
}

func getMaxWidth(root *TreeNode, curIdx, curLevel int, levelArr []int, maxWidth int) (int, []int) {
	if len(levelArr) < curLevel {
		levelArr = append(levelArr, curIdx)
	} else {
		levelLeftIdx := levelArr[curLevel-1]
		if levelLen := curIdx - levelLeftIdx + 1; levelLen > maxWidth {
			maxWidth = levelLen
		}
	}

	if root.Left != nil {
		maxWidth, levelArr = getMaxWidth(root.Left, 2*curIdx, curLevel+1, levelArr, maxWidth)
	}

	if root.Right != nil {
		maxWidth, levelArr = getMaxWidth(root.Right, 2*curIdx+1, curLevel+1, levelArr, maxWidth)
	}

	return maxWidth, levelArr
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {

	node4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	node7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	node2 := &TreeNode{
		Val:   2,
		Left:  node4,
		Right: node5,
	}

	node3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: node7,
	}

	node1 := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}
	fmt.Println(widthOfBinaryTree(node1))

}
