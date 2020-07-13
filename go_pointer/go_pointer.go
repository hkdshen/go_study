package main
import "fmt"
const max = 3
func main(){
	number := [max]int{5, 6, 7}
	var ptrs [max]*int //指针数组
	//将number数组的值的地址赋给ptrs
	/*error,这个问题是range循环的实现逻辑引起的。跟for循环不一样的地方在于range循环中的x变量是临时变量。range循环只是将值拷贝到x变量中。
	因此内存地址都是一样的。
	for i := 0; i < max; i++ {
	        ptrs[i] = &number[i]
	    }
	*/
	for i:= range number {
		ptrs[i] = &number[i]
	}
	for i, x := range ptrs {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}
	fmt.Printf("go pointer")
}
