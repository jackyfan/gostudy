package testtimer

import (
	"fmt"
	"time"
)

func Testtimer() {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C //停止2秒钟，有点像sleep
	fmt.Println("timer expired!")
}
func Testticker() {
	/*每秒钟循环*/
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}
}
func TesttimerTicker() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println(t)
		}
	}()
	//设置一个timer，10钞后停掉ticker
	timer := time.NewTimer(10 * time.Second)
	<-timer.C

	ticker.Stop()
	fmt.Println("timer expired!")
}
