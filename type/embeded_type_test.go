package _type

import (
	"testing"
	"fmt"
)

type notifier interface {
	notify()
}
type user struct {
	name string
	email string
}

func (u *user) notify(){
	fmt.Println("Sending mail to User:",u.name,u.email)
}

type admin struct {
	user            // 嵌入类型，只有类型名字
	level string    // 字段, 有类型名字和字段名
}

func sendNotification(n notifier)  {
	n.notify()
}

func (a *admin) notify() {
	fmt.Println("Sending mail to Admin:",a.name,a.email)
}

func Test_embeded_type(t *testing.T)  {
	ad:= admin{
		user:user{
			name:"tom",
			email:"tom@gmail.com",
		},
		level:"super",
	}

	ad.user.notify()                 // Sending mail to User: tom tom@gmail.com

	ad.notify()                      // Sending mail to Admin: tom tom@gmail.com

	sendNotification(&ad)            // Sending mail to Admin: tom tom@gmail.com

	sendNotification(&ad.user)       // Sending mail to User: tom tom@gmail.com
}
