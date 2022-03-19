package main

import (
	"fmt"
)

/*
	接口实例
*/

type Usb interface {
	start()
	stop()
}

type phone struct {
	name string
}

type camera struct {
	name string
}

func (p phone) start() {
	fmt.Printf("%v手机开始工作...\n", p.name)
}

func (p phone) stop() {
	fmt.Printf("%v手机停止工作...\n", p.name)
}

func (c camera) start() {
	fmt.Printf("%v照相机开始工作...\n", c.name)
}

func (c camera) stop() {
	fmt.Printf("%v照相机停止工作...\n", c.name)
}

type computer struct {
	name string
}

func (c computer) working(usb Usb) {
	usb.start()
	usb.stop()
}
func main() {
	var c = computer{"华硕"}
	var ca = camera{"佳能"}
	var p = phone{"华为"}
	c.working(ca)
	c.working(p)
}
