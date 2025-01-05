package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func cancelTest() {
	// Background returns a non-nil, empty [Context] for initialization
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("The task is cancelled")
				return
			default:
				fmt.Println("The task is running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func terminateTest() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	go func() {
		for {
			select {
			case <-ctx.Done():
				if errors.Is(ctx.Err(), context.DeadlineExceeded) {
					fmt.Println("DeadlineExceeded: operation timeout")
				}
				return
			default:
				fmt.Println("The task is running")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(6 * time.Second)
}

type User struct {
	Name string
}

func infoTest() {
	user := User{Name: "Alice"}
	ctx := context.WithValue(context.Background(), "user", user)
	go func(ctx context.Context) { // func parameters definition
		userInfo := ctx.Value("user").(User)
		fmt.Printf("User Info: %s\n", userInfo.Name)
	}(ctx) // pass parameter

	time.Sleep(3 * time.Second)
}

// implements a simple producer to send message.
func main() {
	//cancelTest()
	//terminateTest()
	infoTest()
}
