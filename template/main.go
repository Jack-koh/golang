package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {
	user := User{Name: "Jack", Email: "rhgusdud09@naver.com", Age: 22}
	user2 := User{Name: "Yerin", Email: "yerin09@naver.com", Age: 40}
	users := []User{user, user2}
	tmpl, err := template.New("Tmpl1").ParseFiles("tmpl/tmpl1.tmpl", "tmpl/tmpl2.tmpl", "tmpl/tmpl3.tmpl")
	if err != nil {
		panic(err)
	}

	tmpl.ExecuteTemplate(os.Stdout, "tmpl3.tmpl", users)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", user2)
	tmpl.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user)
}
