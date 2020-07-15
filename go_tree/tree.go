package go_tree
import "fmt"

//创建二叉树结构体
type Node struct {
	Value int
	Left,Right *Node
}
//创建一个节点
func CreateNode (v int) *Node {
	return &Node{v,nil,nil}
}
//给节点赋值
func (n *Node) SetValue(v int) {
	n.Value = v
}

func (n *Node) printfValue(){
	fmt.Println(n.Value)
}

func (n *Node) IterateTreeFunc (f func(*Node)) {
	if n == nil {
		return
	}
	//n.printfValue()
	f(n)
	n.Left.IterateTreeFunc(f)
	n.Right.IterateTreeFunc(f)
}

func (n *Node) IterateTree (){
	n.IterateTreeFunc(func (n *Node){
		n.printfValue()
	})
}
