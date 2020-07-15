package main
import (
	"fmt"
	"github.com/hkdshen/go_study/go_tree"
)
func main(){
	root := go_tree.Node{3,nil,nil} //普通的二叉树赋值
	root.Left = &go_tree.Node{}
	root.Right = &go_tree.Node{Value:5}
	root.Right.Left = new(go_tree.Node)
	root.Right.Right = &go_tree.Node{6,nil,nil}
	root.Left.Right = go_tree.CreateNode(2)
	root.Left.Left = go_tree.CreateNode(7)
	root.Right.Left.SetValue(4)

	root.IterateTree()
	fmt.Println("go functional programming")
	
}
