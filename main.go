package main

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"go-kit/pkg"
	"net/http"
)

func main()  {
	// 1.定义的Server/server.go
	s := pkg.Server{}

	// 2.在用EndPoint/endpoint.go 创建业务服务
	hello := pkg.MakeServerEndpointHello(s)
	bye := pkg.MakeServerEndpointBye(s)

	// 3.使用 kit 创建 handler
	// 固定格式
	// 传入 业务服务 以及 定义的 加密解密方法
	helloServer := httpTransport.NewServer(hello,pkg.HelloDecodeRequest,pkg.HelloEncodeResponse)
	byeServer := httpTransport.NewServer(bye,pkg.ByeDecodeRequest,pkg.ByeEncodeResponse)

	// 使用http包启动服务
	go http.ListenAndServe("0.0.0.0:8000", helloServer)
	go http.ListenAndServe("0.0.0.0:8001", byeServer)
	select {}

}
