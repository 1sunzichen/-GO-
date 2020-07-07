package main

import "123go/treepart2"
func main(){
	var root tree.TreeNode
	root=tree.TreeNode{Value:3}
	root.Left=&tree.TreeNode{}
	root.Right=&tree.TreeNode{5,nil,nil}
	root.Right.Left=new(tree.TreeNode)
}