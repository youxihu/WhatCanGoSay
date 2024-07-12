// 运算符
package main

import "fmt"

//二元运算符
// & "按位与"     | "按位或"      ^ "按位异或"
// 1 & 1 --> 1   1 | 1 --> 1    1 ^ 1 --> 0
// 1 & 0 --> 0   1 | 0 --> 1    1 ^ 0 --> 1
// 0 & 1 --> 0   0 | 1 --> 1    0 ^ 1 --> 1
// 0 & 0 --> 0   0 | 0 --> 0    0 ^ 0 --> 0

//	优先级     运算符
//	7      ^ !
//	6      * / % << >> & &^
//	5      + - | ^
//	4      == != < <= >= >
//	3      <-
//	2      &&
//	1      ||`

type ByteSize float64
type BitFlag int

const (
	Active  BitFlag = 1 << iota //0000000000000001
	Send                        //0000000000000010
	Receive                     //0000000000000100
)

const ( // <<左偏移
	_           = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func CalConst() {
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}

// 算术运算符
// i = i + 1 --> i += 1
// i = i - 1 --> i -= 1

func Cal() {
	for i := 1; i <= 100; i += 10 {
		fmt.Println(i)
	}
}

func main() {
	//CalConst()
	Cal()
}
