package main

import "fmt"

type User struct {
	Name string
	Age int
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetUser(name string, age int) {
	u.Name = name
	u.Age = age
}

func (u *User) GetUser() {
	fmt.Printf("用户名：%s, 年龄：%d\n", u.Name, u.Age)
}


func main() {
	user := NewUser()
	user.GetUser()

	user.SetUser("sisi", 12)

	user.GetUser()
}