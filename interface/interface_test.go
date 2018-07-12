package _interface

import (
	"testing"
	"fmt"
)

type notifer interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify(){
	fmt.Println("Sending user email to",u.name,u.email)
}

func sendNotification(n notifer)  {
	n.notify()
}

func Test_user(t *testing.T)  {
	u := user{"tom","tom@gmail.com"}
	//sendNotification(u)          // 如果使用user的值，出现编译错误
	sendNotification(&u)           // 使用user的指针，编译通过
}








type repository interface {
	save(s string)
}

type mongoRepo struct {
}

type mysqlRepo struct {
}

func (mondo mongoRepo) save(s string) {
	fmt.Println("Mongo save",s)
}

func (mysql *mysqlRepo) save(s string) {
	fmt.Println("Mysql save",s)
}

func saveData(repo repository){
	repo.save("hello")
}

func Test_repository(t *testing.T)  {
	mongo := mongoRepo{}
	saveData(mongo)
	saveData(&mongo)

	mysql := mysqlRepo{}
	//saveData(mysql)
	saveData(&mysql)
}
