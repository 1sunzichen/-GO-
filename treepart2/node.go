package tree

import (
	"fmt"
)

type TreeNode struct{
	Value int
	Left,Right *TreeNode
}
func CreateNode(value int) *TreeNode{
	return &TreeNode{Value:value}
}
func (node TreeNode) Print(){
	fmt.Print(node.Value,"  ")

}

// func main(){
// 	var root TreeNode
// 	root=TreeNode{value:3}
// 	root.Left=&TreeNode{}
// 	root.Right=&TreeNode{5,nil,nil}
// 	root.Right.Left=new(TreeNode)
// 	fmt.Println(root)
// 	// nodes:=[]treeNode{
// 	// 	{value:3},
// 	// 	{},
// 	// 	{6,nil,&root},
// 	// }
// 	// fmt.Println(nodes)
// 	root.Print()
// 	fmt.Println()
// 	root.traverse()
// }
//值接受者go  指针接受者