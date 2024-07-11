package main

import (
	"fmt"
)

// 常量 const
const beef, two, d = "eat", 2, "veg"

const (
	RED    = 31
	ORANGE = 33
	YELLOW = 33
	GREEN  = 32
	BLUE   = 34
	INDIGO = 36
	VIOLET = 35
)

func ColorPrint() {
	fmt.Printf("\033[%dm%s\033[0m ", RED, "RED")
	fmt.Printf("\033[%dm%s\033[0m ", ORANGE, "ORANGE")
	fmt.Printf("\033[%dm%s\033[0m ", YELLOW, "YELLOW")
	fmt.Printf("\033[%dm%s\033[0m ", GREEN, "GREEN")
	fmt.Printf("\033[%dm%s\033[0m ", BLUE, "BLUE")
	fmt.Printf("\033[%dm%s\033[0m ", INDIGO, "INDIGO")
	fmt.Printf("\033[%dm%s\033[0m ", VIOLET, "VIOLET")
}

const (
	//	a = iota
	//	b = iota
	//	c = iota
	a = iota //第一个 iota 相等于 index 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；简写为如下
	b
	c
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(beef, two, d)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	fmt.Println(a, b, c)
	ColorPrint()
}
