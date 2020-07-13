package main
import "fmt"
import "unsafe"
//Commonly used to declare global variables
var (
	a int
	b bool
)
//Constants can be built-in functions to evaluate the value of an expression, otherwise it will build failed.
//built-in functions likes len(),cap(), unsafe.Sizeof()....
const (
	constA = "string"
	constB = len(constA)
	constC = unsafe.Sizeof(constA)   //unsafe.Sizeof return the length of argv type
)
func main(){
	//test default value of follow argv types
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	//-------------------------------------------------
	//var myString string = "string"
	myString := "string"
	fmt.Printf("%q\n",myString)
	//-------------------------------------------------
	//"_" means ignore the value
	_,myInt,myString1 := numbers()
	fmt.Printf("%v,%q\n",myInt,myString1)
	fmt.Println(myInt,myString1)
	//-------------------------------------------------
	fmt.Println(constA,constB,constC)
	fmt.Printf("go argv")
}
func numbers()(int,int,string){
	a , b , c := 1 , 2 , "str"
	return a,b,c
}
