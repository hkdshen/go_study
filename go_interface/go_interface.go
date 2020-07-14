package main
import "fmt"
//类似c++的基类与派生类
type test interface {
	play()
}

type war3 struct{
	method *struct{
		play func()
	}
}
type dota struct{
	method *struct{
		play func()
	}
}
func (w war3) play() {
	fmt.Printf("now start to play war3 game\n")
}
func (d dota) play(){
	fmt.Printf("now start to play dota game\n")
}
func main(){
	var test test
	test = new(war3)
	test.play()

	test = new(dota)
	test.play()

	var test2 = new(war3)
	test2.play()
	fmt.Printf("go interface")
}
