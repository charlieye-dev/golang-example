package main

import (
    "fmt"
)

type A struct {
    Name string
}

type B struct {
    Name string
}

func main() {
    a := A{}
    a.Print()
    b := B{}
    b.Print()
}

//编译器根据接收者的类型，来判断它是属于哪个方法
func (a A) Print() {
    //取一个变量a，a就是接收者，它的接收者的类型就是structA,Print就是方法的名称，参数在Print()的括号中定义
    //receiver就是这个函数的第一个接收者，而且是强制规定的，这个时候就变成了一个方法
    fmt.Println("A")
}
func (b B) Print() {
    fmt.Println("B")
}
