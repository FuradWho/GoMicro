package main

import (
	"context"
	"fmt"
	"github.com/FuradWho/GoMicro/go_kit_consumer/endpoint"
	"github.com/FuradWho/GoMicro/go_kit_consumer/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"net/url"
	"os"
)

func main() {

	tgt, _ := url.Parse("http://192.168.175.143:8080")
	//创建一个直连client，这里我们必须写两个func,一个是如何请求,一个是响应我们怎么处理
	client := httpTransport.NewClient("GET", tgt, transport.GetUserInfoRequest, transport.GetUserInfoResponse)

	getUserInfo := client.Endpoint()
	ctx := context.Background() //创建一个上下文

	//执行
	res, err := getUserInfo(ctx, endpoint.UserRequest{Uid: 101}) //使用go-kit插件来直接调用服务
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userinfo := res.(endpoint.UserResponse)
	fmt.Println(userinfo.Result)
}
