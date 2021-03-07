package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	//"github.com/go-kit/kit/endpoint"
	"net/http"
)

func HelloDecodeRequest(ctx context.Context,req *http.Request) (interface{},error) {
	name := req.URL.Query().Get("name")
	if name == ""{
		return nil,errors.New("null")
	}
	return HelloRequest{Name: name},nil
}

func HelloEncodeResponse(ctx context.Context, w http.ResponseWriter,res interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(res)
}

func ByeDecodeRequest(ctx context.Context,req *http.Request) (interface{},error) {
	a := req.URL.Query().Get("a")
	b := req.URL.Query().Get("b")

	if a == "" || b == ""{
		return nil,errors.New("null")
	}
	a_, _ := strconv.Atoi(a)
	b_, _ := strconv.Atoi(b)
	return ByeRequest{A: a_, B: b_},nil
}

func ByeEncodeResponse(ctx context.Context, w http.ResponseWriter,res interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(res)
}