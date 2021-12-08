package main

import (
	"context"
	"fmt"
	clientEndpoint "github.com/FuradWho/GoMicro/go_kit_consumer/endpoint"
	"github.com/FuradWho/GoMicro/go_kit_consumer/transport"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	httpTransport "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	"io"
	"net/url"
	"os"
)

func main() {
	{
		config := consulapi.DefaultConfig()
		config.Address = "192.168.175.143:8500"
		api_client, _ := consulapi.NewClient(config)
		client := consul.NewClient(api_client)

		var logger log.Logger
		{
			logger = log.NewLogfmtLogger(os.Stdout)
			var Tag = []string{"primary"}
			//第二部创建一个consul的实例
			instancer := consul.NewInstancer(client, logger, "userService", Tag, true) //最后的true表示只有通过健康检查的服务才能被得到
			{
				factory := func(service_url string) (endpoint.Endpoint, io.Closer, error) { //factory定义了如何获得服务端的endpoint,这里的service_url是从consul中读取到的service的address我这里是192.168.3.14:8000
					tart, _ := url.Parse("http://" + service_url)                                                                                 //server ip +8080真实服务的地址
					return httpTransport.NewClient("GET", tart, transport.GetUserInfoRequest, transport.GetUserInfoResponse).Endpoint(), nil, nil //我再GetUserInfo_Request里面定义了访问哪一个api把url拼接成了http://192.168.3.14:8000/v1/user/{uid}的形式
				}
				endpointer := sd.NewEndpointer(instancer, factory, logger)
				endpoints, _ := endpointer.Endpoints()
				fmt.Println("服务有", len(endpoints), "条")
				getUserInfo := endpoints[0] //写死获取第一个
				ctx := context.Background() //第三步：创建一个context上下文对象

				//第四步：执行
				res, err := getUserInfo(ctx, clientEndpoint.UserRequest{Uid: 101})
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				//第五步：断言，得到响应值
				userinfo := res.(clientEndpoint.UserResponse)
				fmt.Println(userinfo.Result)
			}
		}
	}
}

func connectionTest() {
	tgt, _ := url.Parse("http://192.168.175.143:8080")
	//创建一个直连client，这里我们必须写两个func,一个是如何请求,一个是响应我们怎么处理
	client := httpTransport.NewClient("GET", tgt, transport.GetUserInfoRequest, transport.GetUserInfoResponse)

	getUserInfo := client.Endpoint()
	ctx := context.Background() //创建一个上下文

	//执行
	res, err := getUserInfo(ctx, clientEndpoint.UserRequest{Uid: 101}) //使用go-kit插件来直接调用服务
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	userinfo := res.(clientEndpoint.UserResponse)
	fmt.Println(userinfo.Result)
}
