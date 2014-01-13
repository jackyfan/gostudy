package testchan

import (
	"fmt"
	"time"
)

func Testchan() {
	channel := make(chan string) //注意: buffer为1
	go func() {
		fmt.Println("write \"hello\" start!")
		channel <- "hello"
		fmt.Println("write \"hello\" done!")
		fmt.Println("write \"World\" start!")
		channel <- "World" //Reader在Sleep，这里在阻塞
		fmt.Println("write \"World\" done!")

		fmt.Println("Write go sleep...")
		time.Sleep(3 * time.Second)
		channel <- "channel"
		fmt.Println("write \"channel\" done!")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Reader Wake up...")

	msg := <-channel
	fmt.Println("Reader: ", msg)

	msg = <-channel
	fmt.Println("Reader: ", msg)

	msg = <-channel //Writer在Sleep，这里在阻塞
	fmt.Println("Reader: ", msg)
}
func Testselect() {
	//创建两个channel - c1 c2
	c1 := make(chan string)
	c2 := make(chan string)

	//创建两个goruntine来分别向这两个channel发送数据
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "Hello"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "World"
	}()
	timeout := 0
	for {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1 received", msg1)
		case msg2 := <-c2:
			fmt.Println("msg2 received", msg2)
		case <-time.After(time.Second * 2):
			fmt.Println("Time Out")
			timeout++
		}
		if timeout > 3 {
			break
		}
	}
}
