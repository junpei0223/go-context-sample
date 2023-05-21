package main

import (
	"example.com/sample/server"
)

func main() {
	srv := server.DefaultServe
	srv.ListenAndServe()
}
