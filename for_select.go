package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 50)
	ch2 := make(chan int, 50)
	for i := 0; i <= 10; i++ {
		ch1 <- i
		ch2 <- i
	}

	v := <-ch1
	fmt.Println(v)

	//close(ch1)
	//for select 情况下关闭channel 会一直读那个关闭的channel
	//for {// 这里监控好几个管道 for + switch 做的到吗
	//	select {
	//	case v, _ := <-ch1:
	//		fmt.Println("ch1",v)
	//	case v, _ := <-ch2:
	//		fmt.Println("ch2",v)
	//	default:
	//		time.Sleep(time.Second)
	//		fmt.Println("结束了")
	//		break
	//	}
	//}
	//
	//for {
	//	if v,_ := <-ch1 :{
	//
	//	}
	//}
}
