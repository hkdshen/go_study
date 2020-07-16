package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//使用互斥锁
var mutex sync.Mutex
var ticketNum int32
var done = make(chan bool)
func sellTicketMutex(num int) {
	for ticketNum<100 {
		mutex.Lock()
		if ticketNum<100 {
			ticketNum++
			fmt.Printf("售票窗口%d卖出第%d张票\n", num, ticketNum)
			time.Sleep(time.Millisecond * 5)
		}
		mutex.Unlock()
	}
}
//使用原子操作
func sellTicketAtomic(num int){
	for ticketNum<100 {
		mutex.Lock()
		if ticketNum<100 {
			atomic.AddInt32(&ticketNum,1)
			fmt.Printf("售票窗口%d卖出第%d张票\n", num, ticketNum)
			time.Sleep(time.Millisecond * 5)
		}
		mutex.Unlock()
	}
}
//使用通道
func sellTicketChan(num int){
	for ticketNum<100 {
		if <-done {
			if ticketNum<100{
				ticketNum++
				fmt.Printf("售票窗口%d卖出第%d张票\n", num, ticketNum)
				time.Sleep(time.Millisecond * 5)
				done <- true
			}else{
				close(done)
			}
		}
	}
}
func main(){
	go func(){
		done<-true
	}()
	for i:=1;i<5;i++{
		//go sellTicketMutex(i)
		//go sellTicketAtomic(i)
		go sellTicketChan(i)
	}
	//等待线程执行结束
	var input string
	fmt.Scanln(&input)
	fmt.Printf("go routine safe")
}
