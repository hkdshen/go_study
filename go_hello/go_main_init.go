package main
import "fmt"
//if has func named main and init,then will first run init ,next run main
func main(){
	fmt.Printf("hello world main\n")
}
func init(){
	fmt.Printf("hello world init\n")
}