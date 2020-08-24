package tree

import "fmt"

func (node *TreeNode) Traverse(){
	//if node == nil{
	//	return
	//}
	//
	//if node.Left !=nil{
	//
	//	node.Left.traverse()
	//}
	//node.Print()
	//node.Right.traverse()
	node.Tra(func(b *TreeNode){
		b.Print()

	},"begin")

}

func (node *TreeNode) Tra(f func(*TreeNode),s string){
	if node == nil{
		return
	}

	//if node.Left !=nil{
	//
	//	node.Left.Tra(f)
	//}
	//node.Print()

	//node.Right.Tra(f)
	f(node)

	node.Left.Tra(f,"left")
	node.Right.Tra(f,"right")
	fmt.Println(s)
	//f(node)

}