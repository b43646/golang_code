package gotest

import "fmt"

type notifier interface {
	notify()
}


type user struct {
	name string
	email string
}

type admin struct {
	name string
	email string
}

func (u *user) notify()  {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func (a *admin) notify()  {
	fmt.Printf("Sending admin user email to %s<%s>\n", a.name, a.email)
}

func sendNotify( n notifier)  {
	n.notify()
}

func main() {
	bill := user{"bill", "bill@email.com"}
	sendNotify(&bill)

	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotify(&lisa)
}
