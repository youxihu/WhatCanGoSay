// 数字类型
package main

import (
	"fmt"
	"math"
)

func tureAndFalse() { //bool型
	a := 1
	a1 := 10
	A := a == 1 && a1 == 2
	B := a == 9 || a1 == 10
	C := a != 1
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

func intAndFloat() { //整数和浮点数
	//模拟数字类型转换
	var example1 int32 = 32
	var example2 int64
	example2 = int64(example1)
	fmt.Printf("example1的值是:%v,example1取到的值类型是:%T\n", example1, example1)
	fmt.Printf("example2的值是:%v,example2取到的值类型是:%T\n", example2, example2)
}

func Uint8FromInt(OldValue int) (uint8, error) { // int类型转换为uint8
	fmt.Printf("OldValue type:%T\n", OldValue)
	fmt.Printf("OldValue:%v\n", OldValue)
	if 0 <= OldValue && OldValue <= math.MaxUint8 { // conversion is safe
		fmt.Printf("\x1b[31mthe max uint8 value: \x1b[34m%v\x1b[0m\n", math.MaxUint8)
		return uint8(OldValue), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", OldValue)
}

func Uint8FromIntUsed() { // int类型转换为uint8
	NewValue, err := Uint8FromInt(77)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("NewValue type: %T\n", NewValue) // Print type
		fmt.Printf("NewValue:%v\n", NewValue)       // Print value
	}
}

func FloatFormInt(floatValue float64) (int, error) { //float转换int
	fmt.Printf("floatValue的值:%v\n", floatValue)
	fmt.Printf("floatValue的类型:%T\n", floatValue)
	if 0 <= floatValue && floatValue <= math.MaxFloat64 {
		return int(floatValue), nil //转换为int类型
	}
	return 0, fmt.Errorf("%d is out of the int range", floatValue)
}

func FloatFormIntUsed() { //float转换int
	intValue, err := FloatFormInt(88.5878) //传入floatValue的值
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("intValue的值:%v\n", intValue)
		fmt.Printf("intValue的类型:%T\n", intValue)
	}
}
func Complex() { //复数类型
	// 创建 complex64 类型的复数 cm1，32位实数和虚数
	cmp1 := complex64(7 + 3.4i)
	// 创建 complex128 类型的复数 cm2，64位实数和虚数
	cmp2 := complex128(8 + 3.5i)
	fmt.Printf("cmp1输出:%v\n", cmp1) // 输出: (7+3.4i)
	fmt.Printf("cmp2输出:%v\n", cmp2) // 输出: (8+3.5i)
	cmp3 := complex(7, 9.55)        //complex是内置函数,用于创建复数
	fmt.Printf("cmp3输出:%v\n", cmp3)
	var cmp4 complex64 = 5 + 10i
	fmt.Printf("cmp4输出:%v\n", cmp4)

}

func main() {
	//intAndFloat()
	//tureAndFalse()
	//Uint8FromIntUsed()
	//FloatFormIntUsed()
	Complex()
}
