package gotest

import "fmt"

type notifier interface {
	notify()
}


type user struct {
	name string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	u := user{"bill", "bill@email.com"}
	u.notify()
}
