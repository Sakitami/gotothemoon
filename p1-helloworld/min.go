package main

import "fmt"

func main() {
    // 声明一个变量，并初始化(没有默认类型，类型自动匹配)
    var a = "Sakitami"
    // 声明一个变量及其类型（没有初始化，默认为零值）
    var b int
    var c bool
    fmt.Println("Hello,World!")
    /* 输出结果：
       Sakitami 0 false */
    fmt.Println(a,b,c)
}


