package main

import (
	"fmt"
	"github.com/FuradWho/GoMicro/go_kit/endpoint"
	"github.com/FuradWho/GoMicro/go_kit/service"
	"github.com/FuradWho/GoMicro/go_kit/transport"
	"github.com/go-kit/kit/transport/http"
)

func main() {
	fmt.Println("!!!!")
	user := service.UserService{}
	point := endpoint.GenUserEndPoint(user)

	http.NewServer(point, transport.DecodeUserRequest, transport.EncodeUserResponse)

	fmt.Println("ok")
}
