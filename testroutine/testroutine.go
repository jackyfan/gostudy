package testroutine

import ()

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func say(s string) {
	runtime.GOMAXPROCS(2) //设置多核
	for i := 0; i < 5; i++ {
		//runtime.Gosched()//默认单线程
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func fibonacciSelect(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

var total_tickets int32 = 10

var mutex sync.Mutex

func sell_tickets(i int) {
	for total_tickets > 0 {
		mutex.Lock()
		if total_tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets--
			fmt.Println(i, total_tickets)
		} else {
			fmt.Println(i, "no tickets")
		}
		mutex.Unlock()
	}
}
func Testroutine() {
	/*a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x :=<-c
	y := <-c
	 // receive from c
	fmt.Println(x, y, x+y)*/
	/*c := make(chan int, 3)//修改2为1就报错，修改2为3可以正常运行
	  c <- 1
	  c <- 2
	  fmt.Println(<-c)
	  fmt.Println(<-c)*/
	/* c := make(chan int, 10)
	   go fibonacci(cap(c), c)
	   for i := range c {
	       fmt.Println(i)
	   }*/
	/*c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)*/
	runtime.GOMAXPROCS(2)        //我的电脑是4核处理器，所以我设置了4
	rand.Seed(time.Now().Unix()) //生成随机种子
	for i := 0; i < 5; i++ {     //并发5个goroutine来卖票
		go sell_tickets(i)
	}
	//等待线程执行完
	var input string
	fmt.Scanln(&input)
	fmt.Println(total_tickets, "done") //退出时打印还有多少票

}
