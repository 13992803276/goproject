package main

import "fmt"

/*
	switch 分支控制
	语法：switch 表达式 {
		case 表达式1..表达式2:
			代码块1
		case 表达式3..表达式4:
			代码块2
		default:
			默认语句块
		}
	注意：1.每一个分支后面不需要加break；
		 2.case 后面可以加多个表达式。
	使用细节：1。switch后面可以不写表达式，相当与if else分支；
			2。case后面的表达式可以是一个范围的判断；
			3。switch后面可以申明和定义一个变量，用；号隔开，但是不推荐使用。
			4.fallthrough 用于switch 穿透，表示执行当前case后继续执行紧接着的一个case代码块。
*/
func main() {
	var key byte
	fmt.Println("请输入一个字符...")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("今天是周一")
		fallthrough
	case 'b':
		fmt.Println("今天是周二")
	case 'c':
		fmt.Println("今天是周三")
	default:
		fmt.Println("输入有误。。。")
	}
}
