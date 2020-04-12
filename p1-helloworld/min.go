package main

import "fmt"

// 这种因式分解关键字的写法一般用于声明全局变量
var (
    gobal1 int
    gobal2 bool
)

func main() {
    // 声明一个变量，根据值自行判定变量类型
    var a = "Sakitami"
    // 声明一个变量及其类型（没有初始化，默认为零值）
    var b int
    var c bool
    var f float64
    var s string
    /* 省略var 声明变量
       若左侧没有声明新的变量，就产生编译错误
       这种不带声明格式(var)的只能在函数体中出现*/
    novar := true
    
    /* 同时声明多个变量*/
    var1,var2 := 1,2 
    var vname1, vname2, vname3 = v1, v2, v3


    fmt.Println("Hello,World!")
    /* 输出结果：
       Sakitami 0 false */
    fmt.Println(a,b,c)
    // 输出结果： 0 ""
    fmt.Printf("%v %q\n", f,s)
    // 输出结果：true 1 2
    fmt.Println(novar,var1,var2)
}


