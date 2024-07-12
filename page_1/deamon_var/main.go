package main

import (
	"fmt"
	"os"
	"runtime"
)

//变量
//声明变量的一般形式是使用 var 关键字：var identifier type

// 书写形式
var age, score int //一次性声明变量类型

var CNY int //分别声明变量类型
var TureOrFalse bool
var name string = "YouXiHu"

var (
	a int // 这种因式分解关键字的写法一般用于声明全局变量
	b bool
	c string
)

func Str() {
	str := "You"
	fmt.Println(str)
}

var str2 string = "You"

func main() {
	fmt.Printf("age初始值:%v, score初始值:%v, CNY初始值:%v\n", age, score, CNY) //当一个变量被声明之后，系统自动赋予它该类型的零值
	fmt.Println(TureOrFalse)
	fmt.Println(name)
	Str()
	fmt.Println("直接使用 str:=You 和 var str2 string = You 是一个意思")
	fmt.Println(str2)

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Println("===============")
	a := "G"
	print(a)
	f1()
}

func f1() {
	a := "1"
	print(a)
	f2()
}

func f2() {
	print(a)
}
