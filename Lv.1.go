package main

import (
	"fmt"
)

//
//定义三个全局通道变量
var chA, chB, chC = make(chan string), make(chan string), make(chan string)

func A() {
	chA <- "A" //将"A"传入通道chA
}
func B() {
	chB <- "B" //将"B"传入通道chB
}
func C() {
	chC <- "C" //将"C"传入通道chC
}
func main() {
	var ID string //定义ID变量来接收通道传出的通道名
	for i := 0; i < 10; i++ {
		go A()
		ID = <-chA
		fmt.Printf("%s", ID) //输出通道名
		go B()
		ID = <-chB
		fmt.Printf("%s", ID) //输出通道名
		go C()
		ID = <-chC
		fmt.Printf("%s", ID) //输出通道名
		fmt.Println()        //每输出一次"ABC"换行
	}

}
