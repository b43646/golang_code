package gotest

import "fmt"

type user struct {
	name string
	email string
}

func (u *user) notify()  {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{user{"john", "john@email.com"}, "super"}
	ad.user.notify()
	ad.notify()
}
