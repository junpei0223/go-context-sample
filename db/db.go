package db

import (
	"context"
	"fmt"
	"time"

	"example.com/sample/session"
)

type Mydb struct {
}

var DefaultDB Mydb

func (db *Mydb) Search(ctx context.Context, userID int) <-chan string {
	out := make(chan string)
	go func() {
		inner := make(chan string)

		go func() {
			time.Sleep(10 * time.Second)
			data := "db search finished."

			select {
			case <-ctx.Done():
				fmt.Println("canceled-child, SessionID: ", session.GetSessionID(ctx))
			case inner <- data:
				fmt.Println("normal end")
			}
			close(inner)

		}()

		select {
		case <-ctx.Done():
			fmt.Println("canceled-parent, SessionID: ", session.GetSessionID(ctx))
		case data := <-inner:
			out <- data
			fmt.Println("normal end")
		}
		close(out)
	}()
	return out
}
