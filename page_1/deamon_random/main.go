// 随机数
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 使用当前时间的纳秒部分作为种子，初始化一个新的随机数生成器
	timeNs := int64(time.Now().Nanosecond())
	r := rand.New(rand.NewSource(timeNs))

	// 生成和打印随机整数
	for i := 0; i < 5; i++ {
		fmt.Printf("%d / ", r.Intn(100))
	}
	fmt.Println()

	// 生成和打印随机浮点数
	for i := 0; i < 5; i++ {
		fmt.Printf("%2.2f / ", 100*r.Float64())
	}
	fmt.Println()
}
