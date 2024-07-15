// 字符串

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//StrAddStr()
	//StrNum()
	//StrLookup()
	//StrHas()
	//StrPosting()
	StrRepCount()
}

func StrAddStr() { //通过len()函数获取字符串的长度
	str1 := "man what can i say"
	si := str1[len(str1)-2] // 获取 str1 的倒数第二个字符
	fmt.Println(string(si))
	fmt.Printf("str1的长度是:%v\n", len(str1))
	str2 := "mamba out"
	Str := str1 + str2
	fmt.Printf("str1+str2= %v\n", Str)
}

func StrNum() {
	// string 类型的零值为长度为零的字符串，即空字符串 ""
	var str string                  // 声明一个 string 类型的变量 str，它会被初始化为零值 ""（空字符串）
	fmt.Printf("str的空值是:%v\n", str) // 打印 str 变量的值，预期输出是一个空字符串
}

func StrLookup() { //字符串的内容（纯字节）可以通过标准索引法来获取，在中括号 [] 内写入索引，索引从 0 开始计数
	var s1 string = "youxihu very handsome"
	var s2 string = "i agree with you"

	sLookup1 := s1[0:6]
	sLookup2 := s2[4:]

	fmt.Printf("sLookup1: %v\nsLookup2: %v\n", sLookup1, sLookup2)
}

// strings 和 strconv 包
// strings.Prefix / strings.Suffix / strings.contains

func StrHas() {
	var strHas string = "youxihu very handsome"
	fmt.Printf("T/F? Does the string \"%s\" have Prefix \"%s\"?\n", strHas, "youxihu") //在Go语言中，格式化字符串中的 %t 是用来格式化布尔值的占位符
	//使用HasPrefix判断是否以youxihu开头
	fmt.Printf("%t\n", strings.HasPrefix(strHas, "youxihu"))
	fmt.Printf("T/F? Does the string \"%s\" have Suffix \"%s\"?\n", strHas, "youxihu")
	//使用HasSuffix判断是否以youxihu结尾
	fmt.Printf("%t\n", strings.HasSuffix(strHas, "youxihu"))
	fmt.Printf("T/F?,Does the string \"%s\" has Contains \"%s\"?\n", strHas, "youxihu")
	//使用Contains判断是否包含youxihu
	fmt.Printf("%t\n", strings.Contains(strHas, "youxihu"))
}

// strings.Index / strings.LastIndex

func StrPosting() {
	var strPosting string = "Hi, I'm Mc20, but do you like porsche?捷"
	//fmt.Printf("Index of \"Mc20\" in strPosting: %d\n", strings.Index(strPosting, "Mc20")) //使用Index判断出现的字符在字符串中出现的第一个字符索引位置
	//fmt.Printf("Index of the last instance of \"Hi\" in strPosting: %d\n", strings.LastIndex(strPosting, "Hi"))
	//fmt.Printf("Index of the last instance of \"porsche\" in strPosting: %d\n", strings.LastIndex(strPosting, "porsche")) //使用LastIndex判断出现的字符在字符串中出现的第一个字符索引位置
	//fmt.Printf("Index of the last instance of \"porsche\" in strPosting: %d\n", strings.LastIndex(strPosting, "Porsche")) //-1 表示字符不包含在字符串中
	// 打印字符串及其长度（以字节为单位）
	fmt.Printf("String: %s\n", strPosting)
	fmt.Printf("Length in bytes: %d\n", len(strPosting))

	// 备份原始字符串
	originalStr := strPosting

	// 将字符串按照 UTF-8 解码成 Unicode 字符，打印字符及其长度
	for len(strPosting) > 0 {
		r, size := utf8.DecodeRuneInString(strPosting)
		fmt.Printf("%c %d\n", r, size)
		strPosting = strPosting[size:]
	}

	// 使用 IndexRune 函数找到字符 '捷' 的索引位置（在备份上操作）
	index := strings.IndexRune(originalStr, '捷')
	fmt.Printf("Index of '捷': %d\n", index)
}

// strings.Replace / strings.Count /strings.Repeat / strings.ToLower / strings.ToUpper

func StrRepCount() {
	var strRep string = " Mc20,Mc20,hello,Mc20,Mc20,Mc20"
	// 使用 strings.Replace 函数替换字符串，并获取替换后的结果和替换次数
	newStr := strings.Replace(strRep, "Mc20", "Porsche", -1) //n = -1 则替换所有字符串 Mc20 为字符串 porsche
	fmt.Printf("将 hello,Mc20 替换成 hello,Porsche\n%s\n", newStr)
	// 使用 strings.Count 函数统计字符在字符串中出现的次数
	secondNewStr := strings.Count(strRep, "Mc20")
	fmt.Printf("Mc20在strRep中出现的次数:%v\n", secondNewStr)
	// 使用 strings.Repeat重复字符串
	var RepeatStrRep string
	RepeatStrRep = strings.Repeat(strRep, 3)
	fmt.Printf("使用repeat重复后的字符串是:%v\n", RepeatStrRep)
	upperStr := strings.ToUpper(strRep)
	lowerStr := strings.ToLower(strRep)
	fmt.Printf("大写:%v\n小写:%v\n", upperStr, lowerStr)
	// 使用 strings.TrimSpace剔除字符串中的空格
	trimSpaceStr := strings.TrimSpace(strRep)
	fmt.Printf("使用TrimSpace剔除后的字符串是:%v\n", trimSpaceStr)
	// 使用strings.Trim来将开头和结尾的剔除掉
	trimStr := strings.Trim(strRep, "Mc20")
	fmt.Printf("使用Trim剔除后的字符串是:%v\n", trimStr)
	trimLeftStr := strings.TrimLeft(strRep, "Mc20")
	trimRightStr := strings.TrimRight(strRep, "Mc20")
	fmt.Printf("使用Trimleft剔除后的字符串是:%v\n", trimLeftStr)
	fmt.Printf("使用TrimRight剔除后的字符串是:%v\n", trimRightStr)
}
