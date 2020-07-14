package main

import (
	"errors"
	"fmt"
)
//使用defer+recover处理error
func test1() int {
	var ret int
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1,num2 := 10,0
	ret = num1/num2
	return ret
}
//使用panic处理error
func test2(a int) error {
	if a > 5 {
		return nil
	}else{
		return errors.New("input error number")
	}
}
func main(){
    fmt.Println(test1())
    err := test2(1)
    if err != nil {
    	panic(err)
	}
	fmt.Printf("go error")
}
