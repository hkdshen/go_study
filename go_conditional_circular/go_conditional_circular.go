package main
import "fmt"
func main(){
	var flag bool
	//while(count<100) {    //go没有while
	for count:=1; count < 100; {
		count++
		flag = true;
		for tmp:=2;tmp<count;tmp++ {
			if count%tmp==0{
				flag = false
			}
		}
		// 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
		if flag == true {
			fmt.Println(count,"素数")
		}else{
			continue
		}
	}
	//----------------------------------------------
	circularRange()
	fmt.Printf("go conditional and circular")
}
func circularRange() {
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}


	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
}
