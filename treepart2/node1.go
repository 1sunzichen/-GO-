package tree


func (node *TreeNode) traverse(){
	if node == nil{
		return 
	}

	if node.Left !=nil{

		node.Left.traverse()
	}
	node.Print()
	node.Right.traverse()

}