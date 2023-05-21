package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"example.com/sample/auth"
	"example.com/sample/db"
)

type MyHandlerFunc func(context.Context, MyRequest)
type MyRequest struct {
	path string
}
type MyResponse struct {
	Code int
	Body string
	Err  error
}

var GetGreeting MyHandlerFunc = func(ctx context.Context, req MyRequest) {
	fmt.Println("GetGreeting")
	var res MyResponse

	userID, err := auth.VerifyAuthToken(ctx)
	if err != nil {
		res = MyResponse{
			Code: 403,
			Err:  err,
		}
		fmt.Println(res)
		return
	}

	dbReqCtx, cancel := context.WithTimeout(ctx, 2*time.Second)

	rcvChan := db.DefaultDB.Search(dbReqCtx, userID)
	data, ok := <-rcvChan
	cancel()

	if !ok {
		res = MyResponse{Code: 408, Err: errors.New("DB request timeout")}
		fmt.Println(res)
		return
	}

	res = MyResponse{
		Code: 200,
		Body: fmt.Sprintf("From path %s, Hello! your ID is %d\ndata -> %s", req.path, userID, data),
	}

	fmt.Println(res)
}

func NotFoundHandler(ctx context.Context, req MyRequest) {

}

func (req *MyRequest) SetPath(path string) {
	req.path = path
}

func (req *MyRequest) GetPath() string {
	return req.path
}
