package main

import "fmt"

/*
	封装：
		- JAVA： public, private, protected, 友元
		- GO: 首字母大写外部可访问 | 首字母小写只能在包内使用
	继承：
		- JAVA: extends来实现
		= GO： 结构体组合实现，内嵌一个或者多个struct
	多态：
		- JAVA: 接口和类，重写与重载
		- GO: 不支持重载, 通过接口实现，通过接口定义方法集，编写多套实现
*/

// define notifier interface
type Notifier interface {
	notify()
}

type WechatNotifier struct {
	Name    string
	Message string
}

func (w *WechatNotifier) notify() {
	fmt.Printf("%v notify %v \n", w.Name, w.Message)
}

type QQNotifier struct {
	Name    string
	Message string
}

func (q *QQNotifier) notify() {
	fmt.Printf("%v notifiy %v\n", q.Name, q.Message)
}

type EmailNotifier struct {
	Name    string
	Message string
}

func (e *EmailNotifier) notify() {
	fmt.Printf("%v notifiy %v\n", e.Name, e.Message)
}

func sendNotify(notifier Notifier) {
	notifier.notify()
}

/*
	Go语言中主要通过定义接口，定义不同的结构体实现接口中的方法，可以把入参作为接口类型，根据输入类型的具体实现来做多态。
*/
func main() {
	w := &WechatNotifier{Name: "wechat", Message: "微信通知"}
	q := &QQNotifier{Name: "qq", Message: "QQ通知"}
	e := &EmailNotifier{Name: "email", Message: "邮件通知"}
	// define sender
	sendNotify(w)
	sendNotify(q)
	sendNotify(e)
}
