package main

import (
	tree "123go/treepart2"
	"fmt"
)

type myTreeNode struct {
	node *tree.TreeNode
}

// func testSparse() {
// 	s := intsets.Sparse{}
// 	s.Insert(1000)
// 	s.Insert(1000000)
// 	fmt.Println(s.Has(10000))
// 	fmt.Println(s.Has(1000))
// }
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	//原因: 指针不能作为接收者, 需要需要定义变量来接送地址
	// myTreeNode{mynode.node.Left}.postOrader()
	// right := myTreeNode{mynode.node.Right}.postOrader()
	// mynode.node.Print()
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}
func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	root.Left.Right.Left = tree.CreateNode(2)
	//root.Left.Right.Left.Right = tree.CreateNode(20)
	//root.Left.Right.Left.Right.Left = tree.CreateNode(200)
	//root.Left.Right.Left.Right.Right = tree.CreateNode(299)
	//root.Right.Left.SetValue(4)   // 左 右 左中 右中 上
	root.Traverse()

	//aa := myTreeNode{&root}
	//aa.postOrder()
	// testSparse()
	// 计数
	nodeCount:=0
	root.Tra(func(node *tree.TreeNode) {
		nodeCount++
	},"++")
	fmt.Println("Node count",nodeCount)
	//  max count
	c:=root.TraverseWithChannel()

	maxNode:=0
	//读
	for node:=range c{
		if(node!=nil){

			if node.Value>maxNode{
				maxNode=node.Value
			}
		}
	}
	fmt.Println("Max node value:",maxNode)
}
