package server

import (
	"context"
	"fmt"

	"example.com/sample/auth"
	"example.com/sample/handlers"
	"example.com/sample/session"
)

type Myserver struct {
	router map[string]handlers.MyHandlerFunc
}

var DefaultServe Myserver = Myserver{
	router: map[string]handlers.MyHandlerFunc{"/": handlers.GetGreeting},
}

func (srv *Myserver) ListenAndServe() {
	for {
		var path, token string
		fmt.Scan(&path)
		fmt.Scan(&token)
		ctx := session.SetSessionID(context.Background())
		go srv.Request(ctx, path, token)
	}
}

func (srv *Myserver) Request(ctx context.Context, path, token string) {
	var req handlers.MyRequest
	req.SetPath(path)

	ctx = auth.SetAuthToken(ctx, token)

	if handler, ok := srv.router[req.GetPath()]; ok {
		handler(ctx, req)
	} else {
		handlers.NotFoundHandler(ctx, req)
	}
}
