package tree

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
func (node *TreeNode) TraverseWithChannel() chan *TreeNode{
	out:= make(chan *TreeNode)
	go func(){
		node.Tra(func(node *TreeNode){
			out<-node// 写的操作
		},"haode")
		close(out)//？猜测：注释后 同时读和写
	}()
	return out

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

	node.Left.Tra(f,"")
	node.Right.Tra(f,"")
	//fmt.Print(s)//
	//f(node)

}
