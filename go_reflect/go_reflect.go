package main
import "fmt"
import "reflect"
/*
注意：反射只能识别public变量，私有变量不能被识别，表现为
小写字母开头的是私有变量，reflect.TypeOf().NumField()和NumMethod返回都为0
*/
type TestStruct struct {
	TestString string
	TestBool bool
	TestPointer *TestStruct
}
func (test TestStruct) TestFunc1(s string){
	fmt.Printf("struct func test1,%s\n",s)
}
func (test TestStruct) TestFunc2(s string){
	fmt.Printf("struct func test2,%s\n",s)
}

func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
	//调用方法
	for i := 0; i< getType.NumMethod(); i++ {
		m := getType.Method(i).Name //先获取方法的名称
		reflectFunc := getValue.MethodByName(m) //注册方法
		args := []reflect.Value{reflect.ValueOf("test")}//构造参数
		//无参数构造方法
		//args = make([]reflect.Value, 0)
		reflectFunc.Call(args)//调用方法
	}
}
func main(){
	var testArgv int64 = 1234
	var myTestStruct = TestStruct{"teststruct",true,nil}
	//-------------------------------------
	fmt.Println("testArgv type is :",reflect.TypeOf(testArgv))
	fmt.Println("testArgv value is :",reflect.ValueOf(testArgv))

	//------------------------------------------
	fmt.Println("myTestStruct type is :",reflect.TypeOf(myTestStruct))
	fmt.Println("myTestStruct.testString type is :",reflect.TypeOf(myTestStruct.TestString))
	fmt.Println("myTestStruct.testBool type is :",reflect.TypeOf(myTestStruct.TestBool))
	fmt.Println("myTestStruct.testPointer type is :",reflect.TypeOf(myTestStruct.TestPointer))

	//------------------------------------------
    DoFiledAndMethod(myTestStruct)
	fmt.Printf("go reflect")
}
