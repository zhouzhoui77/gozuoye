package main

import (
	"fmt"
	"math"
)

//
//向 intChan放入50000个数
func putNum(intChan chan int) {

	for i := 1; i <= 50000; i++ {
		intChan <- i
	}

	//关闭intChan
	close(intChan)
}

// 从 intChan取出数据，并判断是否为素数,如果是，就
// 	//放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for循环
	var flag bool //
	for {
		//判断是否还能继续取
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true //设置标识，先假设是素数
		//判断num是不是素数
		Num := math.Sqrt(float64(num)) //计算num的开方，提高for循环效率
		for i := 2; i <= int(Num); i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数就放入到primeChan
			primeChan <- num
		}
	}

	//一个协程结束，向 exitChan 写入true
	exitChan <- true

}

func main() {

	intChan := make(chan int, 3000)
	primeChan := make(chan int, 2000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, 8)
	//开启一个协程，向 intChan放入 1-50000
	go putNum(intChan)
	//开启8个协程，从 intChan取出数据，并判断是否为素数,如果是，就放入到primeChan
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	//用一个协程判断是否全部完成
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		//当从exitChan 取出了个结果，关闭 prprimeChan
		close(primeChan)
	}()
	//遍历primeChan ,把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}
