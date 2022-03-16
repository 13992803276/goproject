package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	字符串常用函数
*/
func main() {
	// 1.统计字符串的长度。
	str0 := "hello go 世界!"
	fmt.Println(len(str0))
	// 2.字符串遍历（一个中文占3个字节，在遍历的时候会出现乱码，需要进行处理）
	str1 := "hello 西安!"
	str2 := []rune(str1) //将字符串转化成切片
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符= %c \n", str2[i])
	}
	// 3.字符串转int
	str3 := "1234"
	i, err := strconv.Atoi(str3)
	if err == nil {
		fmt.Println("str3 转成 int的结果为：", i)
	} else {
		fmt.Println("转化错误。。。")
	}
	//4.整数转字符串
	n := 12345
	fmt.Println("n转化成的字符串为：", strconv.Itoa(n))
	// 5.字符串转byte
	str4 := "hello golang"
	bytes := []byte(str4)
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("byte第%v个位置的值为 %c \n", i, bytes[i])
	}
	// 6.查找子串，返回一个bool值
	str5 := "hello world!"
	fmt.Println("srt5中是否含有 hello 结果为：", strings.Contains(str5, "hello"))
	// 7.忽略大小写比较相等
	fmt.Println(strings.EqualFold("abc", "ABC"))
	//8.去掉字符串左右两边的空格
	str6 := " hello world ! "
	fmt.Println(strings.TrimSpace(str6))

}
