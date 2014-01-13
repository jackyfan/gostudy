package testpanic

import (
"fmt"
)

func g(i int) {
    if i>1 {
        fmt.Println("Panic!")
        /*暂停当前程序，往上逐级传递错误信息，运所有defer。
        	如果遇到defer的recover，就可以使用panic的参数，然后就从调用defer的函数外层继续往下走
        	如果没有遇到defer的recover，程序就死亡了
        */
        panic(fmt.Sprintf("panic %v", i))
    }
 
}
func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
 
    for i := 0; i < 4; i++ {
        fmt.Println("Calling g with ", i)
        g(i)
        fmt.Println("Returned normally from g.")
     }
}
 
func testpanic() {
    f()
    fmt.Println("Returned normally from f.")
}