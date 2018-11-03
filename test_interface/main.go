package main

import (
	"fmt"
	"time"
)

// 定义一个接口类型
type USB interface {
	Name() string
	Connect()
}

// 上边的方法可以使用下面的一种嵌入式修改
/*
type Connecter interface {
    Connect()
}

type USB interface {
    Name() string
    Connecter
}
*/

// 定义一个结构体
type PhoneConnecter struct {
	name string
}

// 定义一个类型为结构体类型的方法
func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

// usb 引用 PhoneConnecter 的结构体
func Disconnect(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnect:", pc.name)
		return
	}
	fmt.Println("Unknow device")
}

func main() {
	var a USB
	// 实现了所有 usb 接口的  对象 PhoneConnecter
	a = PhoneConnecter{"PhoneConnt"}
	a.Connect()
	Disconnect(a)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) {
		<-ch
		fmt.Println("  ch1  call!")
	}(ch1)
	go func(ch chan int) {
		ch <- 2
		fmt.Println("  ch2  call!")
	}(ch2)

	time.Sleep(time.Second)

	for {
		select {
		case ch1 <- 1:
			fmt.Println("Send operation on ch1 works!")
		case <-ch2:
			fmt.Println("Receive operation on ch2 works!")
		default:
			fmt.Println("Exit now!")
			return
		}
	}
}
