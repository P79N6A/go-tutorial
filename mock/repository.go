package mock

type User struct {
	Id   int
	Name string
	Age  int
}

type Repository interface {
	Create(user User)
	Get(id int) User
	Update(user User)
	Delete(id int)
}
