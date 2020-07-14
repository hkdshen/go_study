package main
import "fmt"
type cb func(int) int

/*
struct rectangle
{
    float getArea()
    {
       return length*high
    }
    float length
    float high
}test;
*/
type rectangle struct {
	length float64
	high   float64
}
func getArea(c rectangle) float64 {
	return c.length*c.high
}
func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调2，x：%d\n", x)
		return x
	})
	//---------------------------
	//闭包测试
	funcA := add(1,2)
	for i:=1;i<5; i++{
		fmt.Println(funcA(i,i+1))
	}
	funcB := makeFibGen()
	for i:=1;i<=10;i++ {
		fmt.Println(funcB())
	}
	//-------------------------------
	var rect rectangle
    rect.length = 5.0
    rect.high   = 4.0
    fmt.Println(getArea(rect))
}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调1，x：%d\n", x)
	return x
}
//函数闭包
func add(x1, x2 int) func(x3 int,x4 int)(int,int,int)  {
	i := 0
	return func(x3 int,x4 int) (int,int,int){
		i++
		return i,x1+x2,x3+x4
	}
}
func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, f1+f2
		return f1
	}
}
