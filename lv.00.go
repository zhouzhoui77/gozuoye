package main

import (
	"fmt"
)

//定义全局变量映射myres、整型通道intchan
var myres = make(map[int]int, 20)
var intchan = make(chan int)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	intchan <- res //将n的阶乘结果发送给通道
}

func main() {
	//
	for i := 1; i <= 20; i++ {
		go factorial(i)
		receive := <-intchan //定义变量receive接收通道intchan传出的值
		myres[i] = receive   //将receive的值赋给映射的第n个元素
	}

	for i := 1; i <= 20; i++ {
		fmt.Printf("myres[%d] = %d\n", i, myres[i]) //按顺序遍历输出映射中的元素
	}

}
