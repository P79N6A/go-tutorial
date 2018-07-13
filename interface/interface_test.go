package _interface

import (
	"testing"
	"fmt"
)

type repository interface {
	save(s string)
}




type mongoRepo struct {
}

type mysqlRepo struct {
}




func saveData(repo repository){
	repo.save("hello")
}





func (mongo mongoRepo) save(s string) {
	fmt.Println("Mongo save",s)
}

func (mysql *mysqlRepo) save(s string) {
	fmt.Println("Mysql save",s)
}



func Test_repository(t *testing.T)  {
	mongo := mongoRepo{}
	saveData(mongo)
	saveData(&mongo)

	mysql := mysqlRepo{}
	//saveData(mysql)
	saveData(&mysql)
}


// Go中实现接口的两种方式:
// 1. 使用值类型接受者
// 2. 使用指针类型接受者


//|      方法接受者       |         值和指针哪些实现了该接口            |
//|      (t  T)          |              T 和 *T                    |
//|      (t *T)          |              *T                         |

